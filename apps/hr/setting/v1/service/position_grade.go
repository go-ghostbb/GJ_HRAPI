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

func Grade(ctx context.Context, page ...*paginator.Pagination) IGrade {
	if len(page) != 0 {
		return &grade{ctx: ctx, page: page[0]}
	}
	return &grade{ctx: ctx}
}

type (
	IGrade interface {
		// GetByKeyword 根據keyword獲取grade
		GetByKeyword(in model.GetByKeywordGradeReq) (out model.GetByKeywordGradeRes, err error)
		// GetByID 根據id獲取grade
		GetByID(in model.GetByIDGradeReq) (out model.GetByIDGradeRes, err error)
		// Insert 新增grade
		Insert(in model.PostGradeReq) (err error)
		// Update 更新grade
		Update(in model.PutGradeReq) (err error)
		// Delete 刪除grade
		Delete(in model.DeleteGradeReq) error
	}

	grade struct {
		ctx  context.Context
		page *paginator.Pagination
	}
)

// GetByKeyword 根據keyword獲取grade
func (g *grade) GetByKeyword(in model.GetByKeywordGradeReq) (out model.GetByKeywordGradeRes, err error) {
	var (
		qGrade  = query.PositionGrade
		qdGrade = qGrade.WithContext(dbcache.WithCtx(g.ctx))
		conds   = []gen.Condition{qdGrade.Where(qGrade.Name.Like("%" + in.Keyword + "%")).Or(qGrade.Code.Like("%" + in.Keyword + "%"))}
		grades  []*types.PositionGrade
	)

	conds = append(conds, qGrade.RankID.Eq(in.RankID))

	grades, err = qdGrade.Scopes(g.page.Where(conds...)).Preload(field.Associations).Find()

	if err = copier.Copy(&out.Items, grades); err != nil {
		return
	}

	out.Total = g.page.Total

	return out, err
}

// GetByID 根據id獲取grade
func (g *grade) GetByID(in model.GetByIDGradeReq) (out model.GetByIDGradeRes, err error) {
	var (
		qGrade = query.PositionGrade
	)
	out.PositionGrade, err = qGrade.WithContext(dbcache.WithCtx(g.ctx)).Preload(field.Associations).Where(qGrade.ID.Eq(in.ID)).First()

	return
}

// Insert 新增grade
func (g *grade) Insert(in model.PostGradeReq) (err error) {
	var (
		qGrade = query.PositionGrade
	)
	in.ID = 0
	in.Rank = nil

	return qGrade.WithContext(dbcache.WithCtx(g.ctx)).Create(in.PositionGrade)
}

// Update 更新grade
func (g *grade) Update(in model.PutGradeReq) (err error) {
	var (
		qGrade = query.PositionGrade
	)

	in.PositionGrade.ID = in.ID
	in.Rank = nil

	return qGrade.WithContext(dbcache.WithCtx(g.ctx)).Save(in.PositionGrade)
}

// Delete 刪除grade
func (g *grade) Delete(in model.DeleteGradeReq) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qGrade = tx.PositionGrade
			err    error
		)

		_, err = qGrade.WithContext(dbcache.WithCtx(g.ctx)).Where(qGrade.ID.Eq(in.ID)).Unscoped().Delete()
		if err != nil {
			return err
		}

		// commit
		return nil
	})
}
