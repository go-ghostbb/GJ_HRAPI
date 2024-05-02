package service

import (
	"context"
	"ghostbb.io/gb/contrib/dbcache"
	"ghostbb.io/gb/frame/g"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"github.com/jinzhu/copier"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"hrapi/apps/hr/setting/v1/model"
	"hrapi/internal/query"
	"hrapi/internal/types"
	"hrapi/internal/utils/driver"
	"hrapi/internal/utils/mathx"
	"hrapi/internal/utils/paginator"
	"math"
)

func WorkShift(ctx context.Context, page ...*paginator.Pagination) IWorkShift {
	if len(page) != 0 {
		return &workShift{ctx: ctx, page: page[0]}
	}
	return &workShift{ctx: ctx}
}

type (
	IWorkShift interface {
		// GetByKeyword 根據keyword, status獲取對應班別
		GetByKeyword(in model.GetByKeywordWorkShiftReq) (out model.GetByKeywordWorkShiftRes, err error)
		// GetByID 根據ID獲取班別
		GetByID(in model.GetByIDWorkShiftReq) (out model.GetByIDWorkShiftRes, err error)
		// Insert 新增班別
		Insert(in model.PostWorkShiftReq) (err error)
		// Update 更新班別
		Update(in model.PutWorkShiftReq) (err error)
		// Delete 刪除班別
		Delete(in model.DeleteWorkShiftReq) error
		// SetStatus 設定狀態
		SetStatus(in model.SetStatusWorkShiftReq) (err error)
		// GetScheduleByDate 獲取schedule
		GetScheduleByDate(in model.GetByDateWorkScheduleReq) (out []*model.GetByDateWorkScheduleRes, err error)
		// UpdateWorkScheduleBatch 更新schedule
		UpdateWorkScheduleBatch(in model.PutBatchWorkScheduleReq) error
		// DeleteWorkSchedule 刪除班表
		DeleteWorkSchedule(in model.DeleteWorkScheduleReq) error
	}

	workShift struct {
		ctx  context.Context
		page *paginator.Pagination
	}
)

// GetByKeyword 根據keyword, status獲取對應班別
func (w *workShift) GetByKeyword(in model.GetByKeywordWorkShiftReq) (out model.GetByKeywordWorkShiftRes, err error) {
	var (
		qWorkShift  = query.WorkShift
		qdWorkShift = qWorkShift.WithContext(dbcache.WithCtx(w.ctx))
		conds       = []gen.Condition{qdWorkShift.Where(qWorkShift.Name.Like("%" + in.Keyword + "%")).Or(qWorkShift.Code.Like("%" + in.Keyword + "%"))}
		leaves      []*types.WorkShift
	)
	if in.Status != "" {
		conds = append(conds, qWorkShift.Status.Is(gbconv.Bool(in.Status)))
	}

	leaves, err = qdWorkShift.Scopes(w.page.Where(conds...)).Find()
	if err != nil {
		return
	}

	if err = copier.Copy(&out.Items, leaves); err != nil {
		return
	}

	out.Total = w.page.Total

	return out, err
}

// GetByID 根據ID獲取班別
func (w *workShift) GetByID(in model.GetByIDWorkShiftReq) (out model.GetByIDWorkShiftRes, err error) {
	var (
		qWorkShift = query.WorkShift
	)

	out.WorkShift, err = qWorkShift.WithContext(dbcache.WithCtx(w.ctx)).Where(qWorkShift.ID.Eq(in.ID)).First()
	return
}

// Insert 新增班別
func (w *workShift) Insert(in model.PostWorkShiftReq) (err error) {
	var (
		qWorkShift = query.WorkShift
	)

	// 計算總時數
	workTotal := in.WorkStart.Unix() - in.WorkEnd.Unix()
	if in.WorkStart.Unix() > in.WorkEnd.Unix() {
		// 如果隔夜, 時間反轉
		workTotal = 24*60*60 - workTotal
	}

	restTotal := in.RestStart.Unix() - in.RestEnd.Unix()
	if in.RestStart.Unix() > in.RestEnd.Unix() {
		// 如果隔夜, 時間反轉
		restTotal = 24*60*60 - restTotal
	}

	in.TotalHours = mathx.Round(float64(workTotal-restTotal)/60/60, 2)
	in.TotalHours = math.Abs(in.TotalHours)

	return qWorkShift.WithContext(dbcache.WithCtx(w.ctx)).Create(in.WorkShift)
}

// Update 更新班別
func (w *workShift) Update(in model.PutWorkShiftReq) (err error) {
	var (
		qWorkShift = query.WorkShift
	)
	in.WorkShift.ID = in.ID

	// 計算總時數
	workTotal := in.WorkStart.Unix() - in.WorkEnd.Unix()
	if in.WorkStart.Unix() > in.WorkEnd.Unix() {
		// 如果隔夜, 時間反轉
		workTotal = 24*60*60 - workTotal
	}

	restTotal := in.RestStart.Unix() - in.RestEnd.Unix()
	if in.RestStart.Unix() > in.RestEnd.Unix() {
		// 如果隔夜, 時間反轉
		restTotal = 24*60*60 - restTotal
	}

	in.TotalHours = mathx.Round(float64(workTotal-restTotal)/60/60, 2)
	in.TotalHours = math.Abs(in.TotalHours)

	return qWorkShift.WithContext(dbcache.WithCtx(w.ctx)).Save(in.WorkShift)
}

// Delete 刪除班別
func (w *workShift) Delete(in model.DeleteWorkShiftReq) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qWorkShift    = tx.WorkShift
			qWorkSchedule = tx.WorkSchedule
			err           error
		)
		// 刪除班別
		_, err = qWorkShift.WithContext(dbcache.WithCtx(w.ctx)).Where(qWorkShift.ID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// 刪除關聯班表
		_, err = qWorkSchedule.WithContext(dbcache.WithCtx(w.ctx)).Where(qWorkSchedule.WorkShiftID.Eq(in.ID)).Delete()
		if err != nil {
			return err
		}

		// commit
		return nil
	})
}

// SetStatus 設定狀態
func (w *workShift) SetStatus(in model.SetStatusWorkShiftReq) (err error) {
	_, err = query.WorkShift.WithContext(dbcache.WithCtx(w.ctx)).Where(query.WorkShift.ID.Eq(in.ID)).Update(query.WorkShift.Status, in.Status)
	return err
}

// GetScheduleByDate 獲取schedule
func (w *workShift) GetScheduleByDate(in model.GetByDateWorkScheduleReq) (out []*model.GetByDateWorkScheduleRes, err error) {
	// 驗證
	err = g.Validator().Data(in).Run(w.ctx)
	if err != nil {
		return nil, err
	}

	var (
		qSchedule = query.WorkSchedule
		schedule  []*types.WorkSchedule
	)

	schedule, err = qSchedule.WithContext(dbcache.WithCtx(w.ctx)).Preload(field.Associations).
		Where(
			qSchedule.ScheduleDate.Gte(driver.NewDate(in.Start)),
			qSchedule.ScheduleDate.Lte(driver.NewDate(in.End)),
			qSchedule.EmployeeID.Eq(in.EmployeeID),
		).Find()
	if err != nil {
		return nil, err
	}

	// copier
	if err = copier.Copy(&out, schedule); err != nil {
		return nil, err
	}

	return out, nil
}

// UpdateWorkScheduleBatch 更新schedule
func (w *workShift) UpdateWorkScheduleBatch(in model.PutBatchWorkScheduleReq) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qSchedule = tx.WorkSchedule
			err       error
		)
		// 刪除
		_, err = qSchedule.WithContext(dbcache.WithCtx(w.ctx)).Unscoped().
			Where(qSchedule.EmployeeID.Eq(in.EmployeeID), qSchedule.ScheduleDate.Like(driver.NewString(in.YearMonth+"-%"))).Delete()
		if err != nil {
			return err
		}

		// 將所有ID改為0 && 更改員工ID
		for _, schedule := range in.Schedules {
			schedule.ID = 0
			schedule.EmployeeID = in.EmployeeID
			schedule.Employee = nil
		}

		// 插入新的
		return qSchedule.WithContext(dbcache.WithCtx(w.ctx)).Create(in.Schedules...)
	})
}

// DeleteWorkSchedule 刪除班表
func (w *workShift) DeleteWorkSchedule(in model.DeleteWorkScheduleReq) error {
	_, err := query.WorkSchedule.WithContext(dbcache.WithCtx(w.ctx)).Where(query.WorkSchedule.ID.Eq(in.ID)).Delete()
	return err
}
