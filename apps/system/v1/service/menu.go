package service

import (
	"cmp"
	"context"
	"ghostbb.io/gb/contrib/dbcache"
	gberror "ghostbb.io/gb/errors/gb_error"
	gbstr "ghostbb.io/gb/text/gb_str"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"github.com/jinzhu/copier"
	"gorm.io/gen"
	"hrapi/apps/system/v1/model"
	"hrapi/internal/query"
	"hrapi/internal/types"
	"hrapi/internal/types/enum"
	"hrapi/internal/utils/deepcopy"
	"slices"
)

var (
	ErrTopRoutesWrong   = gberror.New("top-level routes must start with '/'")
	ErrOtherRoutesWrong = gberror.New("routes above level two cannot have '/' at the beginning.")
)

func Menu(ctx context.Context) IMenu {
	return &menu{ctx: ctx}
}

type (
	IMenu interface {
		// GetByKeyword 根據keyword, status獲取software或web的menu
		GetByKeyword(in model.GetByKeywordMenuReq) (out []*model.GetByKeywordMenuRes, err error)
		// GetByID 根據ID獲取menu
		GetByID(in model.GetByIDMenuReq) (out model.GetByIDMenuRes, err error)
		// Insert 插入menu
		Insert(in model.PostMenuReq) (err error)
		// Update 更新menu
		Update(in model.PutMenuReq) (err error)
		// Delete 刪除menu
		Delete(in model.DeleteMenuReq) error
		// SetStatus 設置menu狀態
		SetStatus(in model.SetStatusMenuReq) (err error)
		// 組裝
		assemble(menus []*types.Menu) (result []*types.Menu)
	}

	menu struct {
		ctx context.Context
	}
)

// GetByKeyword 根據keyword, status獲取software或web的menu
func (m *menu) GetByKeyword(in model.GetByKeywordMenuReq) (out []*model.GetByKeywordMenuRes, err error) {
	var (
		qMenu = query.Menu
		menus []*types.Menu
		conds = []gen.Condition{qMenu.Title.Like("%" + in.Keyword + "%")}
	)
	if in.Status != "" {
		conds = append(conds, qMenu.Status.Is(gbconv.Bool(in.Status)))
	}
	if in.Show != "" {
		switch in.Show {
		case "software":
			conds = append(conds, qMenu.Show.Eq(enum.Software))
		case "web":
			conds = append(conds, qMenu.Show.Eq(enum.Web))
		}
	}

	menus, err = qMenu.WithContext(dbcache.WithCtx(m.ctx)).Where(conds...).Find()
	if err != nil {
		return nil, err
	}

	// 組裝
	newMenus := m.assemble(menus)

	if err = copier.Copy(&out, newMenus); err != nil {
		return
	}

	return
}

// GetByID 根據ID獲取menu
func (m *menu) GetByID(in model.GetByIDMenuReq) (out model.GetByIDMenuRes, err error) {
	var (
		qMenu = query.Menu
	)
	out.Menu, err = qMenu.WithContext(dbcache.WithCtx(m.ctx)).Where(qMenu.ID.Eq(in.ID)).First()
	return
}

// Insert 插入menu
func (m *menu) Insert(in model.PostMenuReq) (err error) {
	var (
		qMenu = query.Menu
	)
	// 路由正確性判斷 & 組件格式化
	if in.ParentID == 0 {
		if !gbstr.HasPrefix(in.Path, "/") {
			return ErrTopRoutesWrong
		}
		if in.Type == enum.Directory {
			in.Component = "LAYOUT"
		}
	} else {
		if gbstr.HasPrefix(in.Path, "/") {
			return ErrOtherRoutesWrong
		}
	}

	return qMenu.WithContext(dbcache.WithCtx(m.ctx)).Create(in.Menu)
}

// Update 更新menu
func (m *menu) Update(in model.PutMenuReq) (err error) {
	var (
		qMenu = query.Menu
	)
	// 路由正確性判斷 & 組件格式化
	if in.ParentID == 0 {
		if !gbstr.HasPrefix(in.Path, "/") {
			return ErrTopRoutesWrong
		}
		if in.Type == enum.Directory {
			in.Component = "LAYOUT"
		}
	} else {
		if gbstr.HasPrefix(in.Path, "/") {
			return ErrOtherRoutesWrong
		}
	}

	// 寫入ID
	in.Menu.ID = in.ID

	return qMenu.WithContext(dbcache.WithCtx(m.ctx)).Save(in.Menu)
}

// Delete 刪除menu
func (m *menu) Delete(in model.DeleteMenuReq) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qMenu = tx.Menu
			M2M   = tx.RoleMenu
		)
		_, err := qMenu.WithContext(dbcache.WithCtx(dbcache.WithCtx(m.ctx))).Unscoped().Where(qMenu.ID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		_, err = M2M.WithContext(dbcache.WithCtx(dbcache.WithCtx(m.ctx))).Unscoped().Where(M2M.MenuID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// commit
		return nil
	})
}

// 組裝menu
func (m *menu) assemble(menus []*types.Menu) (result []*types.Menu) {
	var (
		check   = make(map[uint]struct{})
		item    = make([]*types.Menu, 0)
		menuMap = make(map[uint]*types.Menu)
	)

	// database查詢出來的結果都帶有指針
	// 深層拷貝避免重複建構
	menus = deepcopy.Copy(menus).([]*types.Menu)

	result = make([]*types.Menu, 0)

	// 利用指針組裝
	for _, menu := range menus {
		if _, ok := check[menu.ID]; !ok {
			if menu.ParentID == 0 {
				// 如果沒父節點
				// 直接append回傳陣列
				result = append(result, menu)
			} else {
				// 否則append組裝陣列
				item = append(item, menu)
			}
			if menu.Type == enum.Directory {
				// 如果類別是資料夾
				// 加入map
				// 等待item加入
				menuMap[menu.ID] = menu
			}
			// 標記
			check[menu.ID] = struct{}{}
		}
	}

	// sort
	slices.SortStableFunc(result, func(a, b *types.Menu) int {
		return cmp.Compare(a.Sort, b.Sort)
	})
	slices.SortStableFunc(item, func(a, b *types.Menu) int {
		return cmp.Compare(a.Sort, b.Sort)
	})

	for _, menu := range item {
		if _, ok := menuMap[menu.ParentID]; ok {
			menuMap[menu.ParentID].Children = append(menuMap[menu.ParentID].Children, menu)
		} else {
			result = append(result, menu)
		}
	}

	return
}

// SetStatus 設置menu狀態
func (m *menu) SetStatus(in model.SetStatusMenuReq) (err error) {
	_, err = query.Menu.WithContext(dbcache.WithCtx(m.ctx)).
		Where(query.Menu.ID.Eq(in.ID)).Update(query.Menu.Status, in.Status)
	return err
}
