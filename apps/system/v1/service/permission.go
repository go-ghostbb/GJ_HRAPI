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

func Permission(ctx context.Context, page ...*paginator.Pagination) IPermission {
	if len(page) != 0 {
		return &permission{ctx: ctx, page: page[0]}
	}
	return &permission{ctx: ctx}
}

type (
	IPermission interface {
		// GetByKeyword 根據keyword, status查詢權限標示
		GetByKeyword(in model.GetByKeywordPermReq) (out model.GetByKeywordPermRes, err error)
		// GetByID 根據ID獲取權限標示
		GetByID(in model.GetByIDPermReq) (out model.GetByIDPermRes, err error)
		// Insert 新增權限標示
		Insert(in model.PostPermReq) (err error)
		// Update 更新權限標示
		Update(in model.PutPermReq) (err error)
		// Delete 刪除權限標示
		Delete(in model.DeletePermReq) (err error)
		// SetStatus 設定狀態
		SetStatus(in model.SetStatusPermReq) (err error)
	}

	permission struct {
		ctx  context.Context
		page *paginator.Pagination
	}
)

// GetByKeyword 根據keyword, status查詢權限標示
func (p *permission) GetByKeyword(in model.GetByKeywordPermReq) (out model.GetByKeywordPermRes, err error) {
	var (
		qPerm  = query.Permission
		qdPerm = qPerm.WithContext(dbcache.WithCtx(p.ctx))
		conds  = []gen.Condition{qdPerm.Where(qPerm.Name.Like("%" + in.Keyword + "%")).Or(qPerm.Perm.Like("%" + in.Keyword + "%"))}
		perms  []*types.Permission
	)
	if in.Status != "" {
		conds = append(conds, qPerm.Status.Is(gbconv.Bool(in.Status)))
	}

	perms, err = qdPerm.Scopes(p.page.Where(conds...)).Find()
	if err != nil {
		return
	}

	if err = copier.Copy(&out.Items, perms); err != nil {
		return
	}

	// 寫入總數
	out.Total = p.page.Total

	return
}

// GetByID 根據ID獲取權限標示
func (p *permission) GetByID(in model.GetByIDPermReq) (out model.GetByIDPermRes, err error) {
	var (
		qPerm = query.Permission
	)

	out.Permission, err = qPerm.WithContext(dbcache.WithCtx(p.ctx)).Where(qPerm.ID.Eq(in.ID)).First()
	if err != nil {
		return
	}

	return
}

// Insert 新增權限標示
func (p *permission) Insert(in model.PostPermReq) (err error) {
	in.Roles = nil

	return query.Permission.WithContext(dbcache.WithCtx(p.ctx)).Create(in.Permission)
}

// Update 更新權限標示
func (p *permission) Update(in model.PutPermReq) (err error) {
	in.Permission.ID = in.ID
	in.Roles = nil

	return query.Permission.WithContext(dbcache.WithCtx(p.ctx)).Save(in.Permission)
}

// Delete 刪除權限標示
func (p *permission) Delete(in model.DeletePermReq) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qPerm = tx.Permission
			M2M   = tx.RolePermission
			err   error
		)

		if _, err = qPerm.WithContext(dbcache.WithCtx(p.ctx)).Unscoped().Where(qPerm.ID.Eq(in.ID)).Delete(); err != nil {
			return err
		}

		if _, err = M2M.WithContext(dbcache.WithCtx(p.ctx)).Where(M2M.PermissionID.Eq(in.ID)).Delete(); err != nil {
			return err
		}

		// commit
		return nil
	})
}

// SetStatus 設定狀態
func (p *permission) SetStatus(in model.SetStatusPermReq) (err error) {
	var (
		qPerm = query.Permission
	)

	_, err = qPerm.WithContext(dbcache.WithCtx(p.ctx)).Where(qPerm.ID.Eq(in.ID)).Update(qPerm.Status, in.Status)

	return err
}
