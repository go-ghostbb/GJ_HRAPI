package service

import (
	"context"
	"ghostbb.io/gb/contrib/dbcache"
	"github.com/jinzhu/copier"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"hrapi/apps/hr/setting/v1/model"
	"hrapi/internal/query"
	"hrapi/internal/types"
	"hrapi/internal/utils/paginator"
)

func Rank(ctx context.Context, page ...*paginator.Pagination) IRank {
	if len(page) != 0 {
		return &rank{ctx: ctx, page: page[0]}
	}
	return &rank{ctx: ctx}
}

type (
	IRank interface {
		// GetByKeyword 根據keyword獲取rank
		GetByKeyword(in model.GetByKeywordRankReq) (out model.GetByKeywordRankRes, err error)
		// GetByID 根據id獲取rank
		GetByID(in model.GetByIDRankReq) (out model.GetByIDRankRes, err error)
		// Insert 新增rank
		Insert(in model.PostRankReq) (err error)
		// Update 更新rank
		Update(in model.PutRankReq) (err error)
		// Delete 刪除rank
		Delete(in model.DeleteRankReq) error
	}

	rank struct {
		ctx  context.Context
		page *paginator.Pagination
	}
)

// GetByKeyword 根據keyword獲取rank
func (r *rank) GetByKeyword(in model.GetByKeywordRankReq) (out model.GetByKeywordRankRes, err error) {
	var (
		qRank  = query.PositionRank
		qdRank = qRank.WithContext(dbcache.WithCtx(r.ctx))
		conds  = []gen.Condition{qdRank.Where(qRank.Name.Like("%" + in.Keyword + "%")).Or(qRank.Code.Like("%" + in.Keyword + "%"))}
		ranks  []*types.PositionRank
	)

	ranks, err = qdRank.Scopes(r.page.Where(conds...)).Preload(field.Associations).Find()

	if err = copier.Copy(&out.Items, ranks); err != nil {
		return
	}

	out.Total = r.page.Total

	return out, err
}

// GetByID 根據id獲取rank
func (r *rank) GetByID(in model.GetByIDRankReq) (out model.GetByIDRankRes, err error) {
	var (
		qRank = query.PositionRank
	)
	out.PositionRank, err = qRank.WithContext(dbcache.WithCtx(r.ctx)).Preload(field.Associations).Where(qRank.ID.Eq(in.ID)).First()

	return
}

// Insert 新增rank
func (r *rank) Insert(in model.PostRankReq) (err error) {
	var (
		qRank = query.PositionRank
	)
	in.ID = 0
	in.PositionRank.Grade = nil

	return qRank.WithContext(dbcache.WithCtx(r.ctx)).Create(in.PositionRank)
}

// Update 更新rank
func (r *rank) Update(in model.PutRankReq) (err error) {
	var (
		qRank = query.PositionRank
	)

	in.PositionRank.ID = in.ID

	return qRank.WithContext(dbcache.WithCtx(r.ctx)).Save(in.PositionRank)
}

// Delete 刪除rank
func (r *rank) Delete(in model.DeleteRankReq) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qRank  = tx.PositionRank
			qGrade = tx.PositionGrade
			err    error
		)

		_, err = qRank.WithContext(dbcache.WithCtx(r.ctx)).Where(qRank.ID.Eq(in.ID)).Unscoped().Delete()
		if err != nil {
			return err
		}

		_, err = qGrade.WithContext(dbcache.WithCtx(r.ctx)).Where(qGrade.RankID.Eq(in.ID)).Unscoped().Delete()
		if err != nil {
			return err
		}

		// commit
		return nil
	})
}
