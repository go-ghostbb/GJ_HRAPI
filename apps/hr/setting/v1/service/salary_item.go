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
		// SetAddEmployee 設定加項套用員工
		SetAddEmployee(in model.SetSalaryAddItemEmployeeReq) error
		// GetByKeywordReduce 根據keyword獲取薪資減項設定
		GetByKeywordReduce(in model.GetByKeywordSalaryReduceItemReq) (out model.GetByKeywordSalaryReduceItemRes, err error)
		// InsertReduce 新增薪資減項項目
		InsertReduce(in model.PostSalaryReduceItemReq) (err error)
		// UpdateReduce 更新薪資減項
		UpdateReduce(in model.PutSalaryReduceItemReq) (err error)
		// DeleteReduce 刪除薪資減項
		DeleteReduce(in model.DeleteSalaryReduceItemReq) error
		// SetReduceEmployee 設定減項套用員工
		SetReduceEmployee(in model.SetSalaryReduceItemEmployeeReq) error
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
	SalaryAddList, err = qSalaryAdd.WithContext(dbcache.WithCtx(s.ctx)).Preload(qSalaryAdd.Employee).Scopes(s.page.Where(qSalaryAdd.Name.Like("%" + in.Keyword + "%"))).Find()
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

	in.Employee = nil

	return qSalaryAdd.WithContext(dbcache.WithCtx(s.ctx)).Create(in.SalaryAddItem)
}

// UpdateAdd 更新薪資加項
func (s *salaryItem) UpdateAdd(in model.PutSalaryAddItemReq) (err error) {
	var (
		qSalaryAdd = query.SalaryAddItem
	)
	in.SalaryAddItem.ID = in.ID
	in.Employee = nil

	return qSalaryAdd.WithContext(dbcache.WithCtx(s.ctx)).Save(in.SalaryAddItem)
}

// DeleteAdd 刪除薪資加項
func (s *salaryItem) DeleteAdd(in model.DeleteSalaryAddItemReq) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qSalaryAdd = tx.SalaryAddItem
			qM2M       = tx.SalaryAddItemEmployee
			err        error
		)

		_, err = qSalaryAdd.WithContext(dbcache.WithCtx(s.ctx)).Where(qSalaryAdd.ID.Eq(in.ID)).Unscoped().Delete()
		if err != nil {
			return err
		}

		_, err = qM2M.WithContext(dbcache.WithCtx(s.ctx)).Where(qM2M.SalaryAddItemID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		return nil
	})
}

// SetAddEmployee 設定加項套用員工
func (s *salaryItem) SetAddEmployee(in model.SetSalaryAddItemEmployeeReq) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qM2M   = tx.SalaryAddItemEmployee
			newM2M = make([]*types.SalaryAddItemEmployee, 0)
			err    error
		)

		// 刪除原有
		_, err = qM2M.WithContext(dbcache.WithCtx(s.ctx)).Where(qM2M.SalaryAddItemID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// 建立
		for _, empID := range in.EmployeeID {
			newM2M = append(newM2M, &types.SalaryAddItemEmployee{
				SalaryAddItemID: in.ID,
				EmployeeID:      empID,
			})
		}

		return qM2M.WithContext(dbcache.WithCtx(s.ctx)).Create(newM2M...)
	})
}

// GetByKeywordReduce 根據keyword獲取薪資減項設定
func (s *salaryItem) GetByKeywordReduce(in model.GetByKeywordSalaryReduceItemReq) (out model.GetByKeywordSalaryReduceItemRes, err error) {
	var (
		qSalaryReduce    = query.SalaryReduceItem
		SalaryReduceList []*types.SalaryReduceItem
	)
	SalaryReduceList, err = qSalaryReduce.WithContext(dbcache.WithCtx(s.ctx)).Preload(qSalaryReduce.Employee).Scopes(s.page.Where(qSalaryReduce.Name.Like("%" + in.Keyword + "%"))).Find()
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

	in.Employee = nil

	return qSalaryReduce.WithContext(dbcache.WithCtx(s.ctx)).Create(in.SalaryReduceItem)
}

// UpdateReduce 更新薪資減項
func (s *salaryItem) UpdateReduce(in model.PutSalaryReduceItemReq) (err error) {
	var (
		qSalaryReduce = query.SalaryReduceItem
	)
	in.SalaryReduceItem.ID = in.ID
	in.Employee = nil

	return qSalaryReduce.WithContext(dbcache.WithCtx(s.ctx)).Save(in.SalaryReduceItem)
}

// DeleteReduce 刪除薪資減項
func (s *salaryItem) DeleteReduce(in model.DeleteSalaryReduceItemReq) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qSalaryReduce = tx.SalaryReduceItem
			qM2M          = tx.SalaryReduceItemEmployee
			err           error
		)

		_, err = qSalaryReduce.WithContext(dbcache.WithCtx(s.ctx)).Where(qSalaryReduce.ID.Eq(in.ID)).Unscoped().Delete()
		if err != nil {
			return err
		}

		_, err = qM2M.WithContext(dbcache.WithCtx(s.ctx)).Where(qM2M.SalaryReduceItemID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		return nil
	})
}

// SetReduceEmployee 設定減項套用員工
func (s *salaryItem) SetReduceEmployee(in model.SetSalaryReduceItemEmployeeReq) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qM2M   = tx.SalaryReduceItemEmployee
			newM2M = make([]*types.SalaryReduceItemEmployee, 0)
			err    error
		)

		// 刪除原有
		_, err = qM2M.WithContext(dbcache.WithCtx(s.ctx)).Where(qM2M.SalaryReduceItemID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// 建立
		for _, empID := range in.EmployeeID {
			newM2M = append(newM2M, &types.SalaryReduceItemEmployee{
				SalaryReduceItemID: in.ID,
				EmployeeID:         empID,
			})
		}

		return qM2M.WithContext(dbcache.WithCtx(s.ctx)).Create(newM2M...)
	})
}
