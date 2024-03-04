package service

import (
	"context"
	"ghostbb.io/gb/contrib/dbcache"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"github.com/jinzhu/copier"
	"gorm.io/gen"
	"hrapi/apps/system/v1/model"
	"hrapi/internal/query"
	"hrapi/internal/types"
	"hrapi/internal/utils/paginator"
)

func Role(ctx context.Context, page ...*paginator.Pagination) IRole {
	if len(page) != 0 {
		return &role{ctx: ctx, page: page[0]}
	}
	return &role{ctx: ctx}
}

type (
	IRole interface {
		// GetByKeyword 根據keyword, status查詢角色
		GetByKeyword(in model.GetByKeywordRoleReq) (out model.GetByKeywordRoleRes, err error)
		// GetByID 根據ID獲取角色
		GetByID(in model.GetByIDRoleReq) (out model.GetByIDRoleRes, err error)
		// Insert 新增角色
		Insert(in model.PostRoleReq) (err error)
		// Update 更新角色
		Update(in model.PutRoleReq) (err error)
		// Delete 刪除角色
		Delete(in model.DeleteRoleReq) (err error)
		// SetStatus 設定狀態
		SetStatus(in model.SetStatusRoleReq) (err error)
		// SetMenu 設定可用menu
		SetMenu(in model.SetMenuRoleReq) (err error)
		// SetPerm 設定可用perm
		SetPerm(in model.SetPermRoleReq) (err error)
	}

	role struct {
		ctx  context.Context
		page *paginator.Pagination
	}
)

// GetByKeyword 根據keyword, status查詢角色
func (r *role) GetByKeyword(in model.GetByKeywordRoleReq) (out model.GetByKeywordRoleRes, err error) {
	var (
		qRole  = query.Role
		qdRole = qRole.WithContext(dbcache.WithCtx(r.ctx))
		conds  = []gen.Condition{qdRole.Where(qRole.Name.Like("%" + in.Keyword + "%")).Or(qRole.Code.Like("%" + in.Keyword + "%"))}
		roles  []*types.Role
	)
	if in.Status != "" {
		conds = append(conds, qRole.Status.Is(gbconv.Bool(in.Status)))
	}

	roles, err = qdRole.Scopes(r.page.Where(conds...)).Preload(qRole.Permissions, qRole.Menus).Find()
	if err != nil {
		return
	}

	if err = copier.Copy(&out.Items, roles); err != nil {
		return
	}

	// 組裝menu
	for _, role := range out.Items {
		role.Menus = Menu(r.ctx).assemble(role.Menus)
	}

	// 寫入總數
	out.Total = r.page.Total

	return
}

// GetByID 根據ID獲取角色
func (r *role) GetByID(in model.GetByIDRoleReq) (out model.GetByIDRoleRes, err error) {
	var (
		qRole = query.Role
	)

	out.Role, err = qRole.WithContext(dbcache.WithCtx(r.ctx)).
		Preload(qRole.Permissions, qRole.Menus).Where(qRole.ID.Eq(in.ID)).First()
	if err != nil {
		return
	}

	// 組裝
	out.Menus = Menu(r.ctx).assemble(out.Menus)

	return
}

// Insert 新增角色
func (r *role) Insert(in model.PostRoleReq) (err error) {
	in.Employees = nil
	in.Menus = nil
	in.Permissions = nil

	return query.Role.WithContext(dbcache.WithCtx(r.ctx)).Create(in.Role)
}

// Update 更新角色
func (r *role) Update(in model.PutRoleReq) (err error) {
	var (
		qRole = query.Role
	)
	in.Role.ID = in.ID
	in.Employees = nil
	in.Menus = nil
	in.Permissions = nil

	return qRole.WithContext(dbcache.WithCtx(r.ctx)).Save(in.Role)
}

// Delete 刪除角色
func (r *role) Delete(in model.DeleteRoleReq) (err error) {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qRole    = tx.Role
			qM2MPerm = tx.RolePermission
			qM2MMenu = tx.RoleMenu
		)

		_, err = qRole.WithContext(dbcache.WithCtx(r.ctx)).Unscoped().Where(qRole.ID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		_, err = qM2MPerm.WithContext(dbcache.WithCtx(r.ctx)).Where(qM2MPerm.RoleID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		_, err = qM2MMenu.WithContext(dbcache.WithCtx(r.ctx)).Where(qM2MMenu.RoleID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// commit
		return nil
	})
}

// SetStatus 設定狀態
func (r *role) SetStatus(in model.SetStatusRoleReq) (err error) {
	var (
		qRole = query.Role
	)

	_, err = qRole.WithContext(dbcache.WithCtx(r.ctx)).Where(qRole.ID.Eq(in.ID)).Update(qRole.Status, in.Status)
	return err
}

// SetMenu 設定可用menu
func (r *role) SetMenu(in model.SetMenuRoleReq) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qM2M  = tx.RoleMenu
			qMenu = tx.Menu
			m2m   = make([]*types.RoleMenu, 0)
			err   error
		)
		// 刪除原有連節
		_, err = qM2M.WithContext(dbcache.WithCtx(r.ctx)).Where(qM2M.RoleID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// 查詢所有跟in.MenuId有關係的所有menu
		in.MenuId, err = qMenu.WithContext(r.ctx).QueryAllRelated(in.MenuId)
		if err != nil {
			return err
		}

		for _, menuID := range in.MenuId {
			m2m = append(m2m, &types.RoleMenu{RoleID: in.ID, MenuID: menuID})
		}

		return qM2M.WithContext(dbcache.WithCtx(r.ctx)).Create(m2m...)
	})
}

// SetPerm 設定可用perm
func (r *role) SetPerm(in model.SetPermRoleReq) (err error) {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qM2M = tx.RolePermission
			m2m  = make([]*types.RolePermission, 0)
		)

		_, err = qM2M.WithContext(dbcache.WithCtx(r.ctx)).Where(qM2M.RoleID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		for _, permID := range in.PermissionId {
			m2m = append(m2m, &types.RolePermission{RoleID: in.ID, PermissionID: permID})
		}

		// commit if err is nil
		return qM2M.WithContext(dbcache.WithCtx(r.ctx)).Create(m2m...)
	})
}
