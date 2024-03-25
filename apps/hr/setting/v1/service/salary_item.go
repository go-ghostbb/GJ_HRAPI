package service

import (
	"context"
	"ghostbb.io/gb/contrib/dbcache"
	"github.com/jinzhu/copier"
	"hrapi/apps/hr/setting/v1/model"
	"hrapi/internal/query"
	"hrapi/internal/types"
	"hrapi/internal/utils/paginator"
)

func SalaryItem(ctx context.Context, page ...*paginator.Pagination) ISalaryItem {
	if len(page) != 0 {
		return &salaryItem{ctx: ctx, page: page[0]}
	}
	return &salaryItem{ctx: ctx}
}

type (
	ISalaryItem interface {
		// GetByKeywordAdd 根據keyword獲取薪資加項設定
		GetByKeywordAdd(in model.GetByKeywordSalaryAddItemReq) (out model.GetByKeywordSalaryAddItemRes, err error)
		// InsertAdd 新增薪資加項項目
		InsertAdd(in model.PostSalaryAddItemReq) (err error)
		// UpdateAdd 更新薪資加項
		UpdateAdd(in model.PutSalaryAddItemReq) (err error)
		// DeleteAdd 刪除薪資加項
		DeleteAdd(in model.DeleteSalaryAddItemReq) error
		// GetByKeywordReduce 根據keyword獲取薪資減項設定
		GetByKeywordReduce(in model.GetByKeywordSalaryReduceItemReq) (out model.GetByKeywordSalaryReduceItemRes, err error)
		// InsertReduce 新增薪資減項項目
		InsertReduce(in model.PostSalaryReduceItemReq) (err error)
		// UpdateReduce 更新薪資減項
		UpdateReduce(in model.PutSalaryReduceItemReq) (err error)
		// DeleteReduce 刪除薪資減項
		DeleteReduce(in model.DeleteSalaryReduceItemReq) error
	}

	salaryItem struct {
		page *paginator.Pagination
		ctx  context.Context
	}
)

// GetByKeywordAdd 根據keyword獲取薪資加項設定
func (s *salaryItem) GetByKeywordAdd(in model.GetByKeywordSalaryAddItemReq) (out model.GetByKeywordSalaryAddItemRes, err error) {
	var (
		qSalaryAdd    = query.SalaryAddItem
		SalaryAddList []*types.SalaryAddItem
	)
	SalaryAddList, err = qSalaryAdd.WithContext(dbcache.WithCtx(s.ctx)).Scopes(s.page.Where(qSalaryAdd.Name.Like("%" + in.Keyword + "%"))).Find()
	if err != nil {
		return
	}

	if err = copier.Copy(&out.Items, SalaryAddList); err != nil {
		return
	}

	// 寫入總數
	out.Total = s.page.Total

	return
}

// InsertAdd 新增薪資加項項目
func (s *salaryItem) InsertAdd(in model.PostSalaryAddItemReq) (err error) {
	var (
		qSalaryAdd = query.SalaryAddItem
	)

	return qSalaryAdd.WithContext(dbcache.WithCtx(s.ctx)).Create(in.SalaryAddItem)
}

// UpdateAdd 更新薪資加項
func (s *salaryItem) UpdateAdd(in model.PutSalaryAddItemReq) (err error) {
	var (
		qSalaryAdd = query.SalaryAddItem
	)
	in.SalaryAddItem.ID = in.ID

	return qSalaryAdd.WithContext(dbcache.WithCtx(s.ctx)).Save(in.SalaryAddItem)
}

// DeleteAdd 刪除薪資加項
func (s *salaryItem) DeleteAdd(in model.DeleteSalaryAddItemReq) error {
	var (
		qSalaryAdd = query.SalaryAddItem
		err        error
	)

	_, err = qSalaryAdd.WithContext(dbcache.WithCtx(s.ctx)).Where(qSalaryAdd.ID.Eq(in.ID)).Unscoped().Delete()
	if err != nil {
		return err
	}

	return nil
}

// GetByKeywordReduce 根據keyword獲取薪資減項設定
func (s *salaryItem) GetByKeywordReduce(in model.GetByKeywordSalaryReduceItemReq) (out model.GetByKeywordSalaryReduceItemRes, err error) {
	var (
		qSalaryReduce    = query.SalaryReduceItem
		SalaryReduceList []*types.SalaryReduceItem
	)
	SalaryReduceList, err = qSalaryReduce.WithContext(dbcache.WithCtx(s.ctx)).Scopes(s.page.Where(qSalaryReduce.Name.Like("%" + in.Keyword + "%"))).Find()
	if err != nil {
		return
	}

	if err = copier.Copy(&out.Items, SalaryReduceList); err != nil {
		return
	}

	// 寫入總數
	out.Total = s.page.Total

	return
}

// InsertReduce 新增薪資減項項目
func (s *salaryItem) InsertReduce(in model.PostSalaryReduceItemReq) (err error) {
	var (
		qSalaryReduce = query.SalaryReduceItem
	)

	return qSalaryReduce.WithContext(dbcache.WithCtx(s.ctx)).Create(in.SalaryReduceItem)
}

// UpdateReduce 更新薪資減項
func (s *salaryItem) UpdateReduce(in model.PutSalaryReduceItemReq) (err error) {
	var (
		qSalaryReduce = query.SalaryReduceItem
	)
	in.SalaryReduceItem.ID = in.ID

	return qSalaryReduce.WithContext(dbcache.WithCtx(s.ctx)).Save(in.SalaryReduceItem)
}

// DeleteReduce 刪除薪資減項
func (s *salaryItem) DeleteReduce(in model.DeleteSalaryReduceItemReq) error {
	var (
		qSalaryReduce = query.SalaryReduceItem
		err           error
	)

	_, err = qSalaryReduce.WithContext(dbcache.WithCtx(s.ctx)).Where(qSalaryReduce.ID.Eq(in.ID)).Unscoped().Delete()
	if err != nil {
		return err
	}

	return nil
}
