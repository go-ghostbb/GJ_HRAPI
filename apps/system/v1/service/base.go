package service

import (
	"cmp"
	"context"
	"ghostbb.io/gb/contrib/dbcache"
	gberror "ghostbb.io/gb/errors/gb_error"
	"ghostbb.io/gb/frame/g"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"hrapi/apps/system/v1/model"
	"hrapi/internal/query"
	"hrapi/internal/types"
	"hrapi/internal/types/enum"
	"hrapi/internal/utils/jwtx"
	"hrapi/internal/utils/jwtx/claims"
	"hrapi/internal/utils/password"
	"hrapi/internal/utils/token"
	"slices"
)

func Base(ctx context.Context) IBase {
	return &base{ctx: ctx}
}

type (
	IBase interface {
		Migrate() error
		Login(loginType enum.MenuShow, in model.LoginReq) (out model.LoginRes, err error)
		Logout(logoutType enum.MenuShow, employeeID uint, accessToken string) error
		GetEmployeeInfo(employeeID uint) (out model.GetEmployeeInfoRes, err error)
		GetPermission(roleCodes []string) (out []string, err error)
		GetMenu(roleCodes []string, show enum.MenuShow) (out []*model.GetMenuRes, err error)
		ChangePassword(in model.ChangePasswordReq) (err error)
	}

	base struct {
		ctx context.Context
	}
)

var (
	ErrUserFrozen        = gberror.New("this user has been frozen")
	ErrPasswordIncorrect = gberror.New("password incorrect")
	ErrUserNotFound      = gberror.New("user not found")
	ErrPermissionDenied  = gberror.New("permission denied")
)

func (b *base) Migrate() error {
	g.Log().Info(b.ctx, "db auto migrate")
	if err := g.DB().AutoMigrate(types.Basic()...); err != nil {
		g.Log().Error(b.ctx, "db auto migrate error:", err.Error())
		return err
	}
	return nil
}

func (b *base) Login(loginType enum.MenuShow, in model.LoginReq) (out model.LoginRes, err error) {
	var (
		qLoginInfo = query.LoginInformation
		loginInfo  *types.LoginInformation
	)

	loginInfo, err = qLoginInfo.WithContext(dbcache.WithCtx(b.ctx)).
		Where(qLoginInfo.Account.Value(in.Account)).
		Preload(qLoginInfo.Employee, qLoginInfo.Employee.Roles).First()
	if err != nil {
		if gberror.Is(err, gorm.ErrRecordNotFound) {
			return model.LoginRes{}, ErrUserNotFound
		}
		return
	}
	// 檢查帳號是否啟用
	if !loginInfo.Status {
		return model.LoginRes{}, ErrUserFrozen
	}
	// 檢查密碼正確性
	if !password.Check(loginInfo.Password, in.Password) {
		return model.LoginRes{}, ErrPasswordIncorrect
	}
	// 如果是軟體，判斷是否有登入軟體權限
	if loginType == enum.Software && !loginInfo.Employee.Backend {
		return model.LoginRes{}, ErrPermissionDenied
	}

	// Create token
	roles := make([]string, 0)
	for _, role := range loginInfo.Employee.Roles {
		roles = append(roles, role.Code)
	}
	accessToken, refreshToken, err := jwtx.Service.CreateToken(loginType, claims.BaseClaims{
		EmployeeID: loginInfo.Employee.ID,
		RealName:   loginInfo.Employee.RealName,
		Account:    loginInfo.Account,
		Roles:      roles,
	})
	if err != nil {
		return model.LoginRes{}, err
	}

	// Put the refreshToken into redis
	key := string(loginType) + ":" + gbconv.String(loginInfo.EmployeeID)
	if err = token.SetRefreshToken(g.Redis(), key, refreshToken, jwtx.Service.GetRefreshTokenEp()); err != nil {
		return model.LoginRes{}, err
	}

	return model.LoginRes{Token: accessToken}, nil
}

func (b *base) Logout(logoutType enum.MenuShow, employeeID uint, accessToken string) error {
	// Delete the refreshToken from redis
	key := string(logoutType) + ":" + gbconv.String(employeeID)
	if err := token.DelRefreshToken(g.Redis(), key); err != nil {
		return err
	}
	// Blacklist the accessToken to avoid using the same of the accessToken for verification
	if err := token.SetBlack(g.Redis(), accessToken); err != nil {
		return err
	}
	return nil
}

func (b *base) GetEmployeeInfo(employeeID uint) (out model.GetEmployeeInfoRes, err error) {
	var (
		qEmployee = query.Employee
		employees *types.Employee
	)

	employees, err = qEmployee.WithContext(dbcache.WithCtx(b.ctx)).
		Preload(qEmployee.LoginInformation, qEmployee.Roles).
		Where(qEmployee.ID.Eq(employeeID)).First()
	if err != nil {
		return
	}

	// copier
	if err = copier.Copy(&out, employees); err != nil {
		return
	}

	return
}

func (b *base) GetPermission(roleCodes []string) (out []string, err error) {
	var (
		qRole = query.Role
		roles []*types.Role
	)
	out = make([]string, 0)

	roles, err = qRole.WithContext(dbcache.WithCtx(b.ctx)).
		Preload(qRole.Permissions).Where(qRole.Code.In(roleCodes...)).Find()
	if err != nil {
		return nil, err
	}

	checkMap := make(map[string]struct{})
	for _, role := range roles {
		for _, permission := range role.Permissions {
			if _, ok := checkMap[permission.Perm]; ok {
				continue
			}
			checkMap[permission.Perm] = struct{}{}
			out = append(out, permission.Perm)
		}
	}

	return
}

func (b *base) GetMenu(roleCodes []string, show enum.MenuShow) (out []*model.GetMenuRes, err error) {
	var (
		qRole = query.Role
		qMenu = query.Menu
		roles []*types.Role
	)

	// query
	roles, err = qRole.WithContext(dbcache.WithCtx(b.ctx)).
		Preload(qRole.Menus.On(
			qMenu.Status.Is(true),
			qMenu.Show.Eq(show),
		).Order(qMenu.Sort)).
		Where(qRole.Code.In(roleCodes...)).Find()
	if err != nil {
		return
	}

	var (
		check   = make(map[uint]struct{})
		item    = make([]*model.GetMenuRes, 0)
		menuMap = make(map[uint]*model.GetMenuRes)
	)
	// 利用指針組裝
	for _, role := range roles {
		for _, menu := range role.Menus {
			tmp := new(model.GetMenuRes)
			if err = copier.Copy(tmp, menu); err != nil {
				return
			}

			if _, ok := check[menu.ID]; !ok {
				if menu.ParentID == 0 {
					// 如果沒父節點
					// 直接append回傳陣列
					out = append(out, tmp)
				} else {
					// 否則append組裝陣列
					item = append(item, tmp)
				}
				if menu.Type == enum.Directory {
					// 如果類別是資料夾
					// 加入map
					// 等待item加入
					menuMap[menu.ID] = tmp
				}
				// 標記
				check[menu.ID] = struct{}{}
			}
		}
	}

	// sort
	slices.SortStableFunc(out, func(a, b *model.GetMenuRes) int {
		return cmp.Compare(a.Sort, b.Sort)
	})
	slices.SortStableFunc(item, func(a, b *model.GetMenuRes) int {
		return cmp.Compare(a.Sort, b.Sort)
	})

	// 組裝
	for _, menu := range item {
		if _, ok := menuMap[menu.ParentID]; ok {
			menuMap[menu.ParentID].Children = append(menuMap[menu.ParentID].Children, menu)
		}
	}

	return
}

func (b *base) ChangePassword(in model.ChangePasswordReq) (err error) {
	// 輸入驗證
	if err = g.Validator().Data(in).Run(b.ctx); err != nil {
		return err
	}

	// 查詢此使用者原本的密碼
	var (
		originPassword string
		qLoginInfo     = query.LoginInformation
	)
	err = qLoginInfo.WithContext(b.ctx).Select(qLoginInfo.Password).Where(qLoginInfo.EmployeeID.Eq(in.EmployeeID)).Scan(&originPassword)
	if err != nil {
		return err
	}

	// 驗證密碼
	if !password.Check(originPassword, in.OldPassword) {
		return gberror.New("incorrect password")
	}

	// 進行密碼修改
	_, err = qLoginInfo.WithContext(dbcache.WithCtx(b.ctx)).
		Where(qLoginInfo.EmployeeID.Eq(in.EmployeeID)).
		Update(qLoginInfo.Password, password.Hash(in.NewPassword))
	if err != nil {
		return err
	}

	return nil
}
