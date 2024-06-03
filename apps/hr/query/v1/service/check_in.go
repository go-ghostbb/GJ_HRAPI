package service

import (
	"context"
	"ghostbb.io/gb/contrib/dbcache"
	"github.com/jinzhu/copier"
	"gorm.io/gen/field"
	"hrapi/apps/hr/query/v1/model"
	"hrapi/internal/query"
	"hrapi/internal/types"
	"hrapi/internal/types/enum"
	"hrapi/internal/utils/driver"
	"hrapi/internal/utils/paginator"
)

func CheckIn(ctx context.Context, page ...*paginator.Pagination) ICheckIn {
	if len(page) > 0 {
		return &checkIn{ctx, page[0]}
	}
	return &checkIn{ctx: ctx}
}

type (
	ICheckIn interface {
		// QueryStatus 查詢出勤狀態
		QueryStatus(in model.QueryStatusReq) (out []*model.QueryStatusRes, err error)
		// QueryStatusByKeyword 根據關鍵字查詢
		QueryStatusByKeyword(in model.QueryStatusByKeywordReq) (out model.QueryStatusByKeywordRes, err error)
		// UpdateCheckInStatus 更新狀態表
		UpdateCheckInStatus(in model.PutCheckInStatusReq) error
	}

	checkIn struct {
		ctx  context.Context
		page *paginator.Pagination
	}
)

// QueryStatus 查詢出勤狀態
func (c *checkIn) QueryStatus(in model.QueryStatusReq) (out []*model.QueryStatusRes, err error) {
	var (
		qCheckInStatus = query.CheckInStatus
		result         []*types.CheckInStatus
	)

	result, err = qCheckInStatus.WithContext(c.ctx).
		Preload(field.Associations, qCheckInStatus.Employee.Department).
		QueryByDateRangeAndEmpID(in.EmployeeID, in.DateRangeStart, in.DateRangeEnd, in.Abnormal)
	if err != nil {
		return nil, err
	}

	if err = copier.Copy(&out, result); err != nil {
		return nil, err
	}

	return out, nil
}

// QueryStatusByKeyword 根據關鍵字查詢
func (c *checkIn) QueryStatusByKeyword(in model.QueryStatusByKeywordReq) (out model.QueryStatusByKeywordRes, err error) {
	var (
		qCheckInStatus = query.CheckInStatus
		result         []*types.CheckInStatus
	)

	result, err = qCheckInStatus.WithContext(c.ctx).
		Preload(field.Associations, qCheckInStatus.Employee.Department).
		QueryByDateRangeAndKeyword("%"+in.Keyword+"%", in.DateRangeStart, in.DateRangeEnd, in.Abnormal, c.page.Offset, c.page.Limit)
	if err != nil {
		return
	}

	if err = copier.Copy(&out.Items, result); err != nil {
		return
	}

	// count
	out.Total, err = qCheckInStatus.WithContext(c.ctx).CountByDateRangeAndKeyword("%"+in.Keyword+"%", in.DateRangeStart, in.DateRangeEnd, in.Abnormal)
	if err != nil {
		return
	}

	return out, nil
}

// UpdateCheckInStatus 更新狀態表
func (c *checkIn) UpdateCheckInStatus(in model.PutCheckInStatusReq) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qCheckInStatus = tx.CheckInStatus
			qWorkShift     = tx.WorkShift
			qWorkSchedule  = tx.WorkSchedule
			status         *types.CheckInStatus
			workShift      *types.WorkShift
			err            error
		)

		status, err = qCheckInStatus.WithContext(dbcache.WithCtx(c.ctx)).Where(qCheckInStatus.ID.Eq(in.ID)).First()
		if err != nil {
			return err
		}

		workShift, err = qWorkShift.WithContext(dbcache.WithCtx(c.ctx)).Where(qWorkShift.ID.Eq(in.WorkShiftID)).First()
		if err != nil {
			return err
		}

		status.WorkCheckIn, status.WorkAttendStatus, status.WorkAttendProcStatus = in.Work, enum.WorkNotSwiped, enum.NotProcessed
		status.OffWorkCheckIn, status.OffWorkAttendStatus, status.OffWorkAttendProcStatus = in.OffWork, enum.OffWorkNotSwiped, enum.NotProcessed
		status.WorkShiftID = in.WorkShiftID

		// 正常需上班時段
		var (
			basicWork    = status.Date.DateTime(workShift.WorkStart)
			basicOffWork = status.Date.DateTime(workShift.WorkEnd)
			layout       = "2006-01-02 15:04"
		)

		if workShift.WorkStart.After(workShift.WorkEnd) {
			basicOffWork = basicOffWork.AddDate(0, 0, 1)
		}

		// 遲到
		if !status.WorkCheckIn.IsZero() {
			if driver.NewDateTime(basicWork.Format(layout), layout).Before(driver.NewDateTime(status.WorkCheckIn.Format(layout), layout)) {
				status.WorkAttendStatus = enum.WorkLate
			} else {
				status.WorkAttendStatus = enum.WorkNormal
				status.WorkAttendProcStatus = enum.ProcNormal
			}
		}

		// 早退
		if !status.OffWorkCheckIn.IsZero() {
			if driver.NewDateTime(basicOffWork.Format(layout), layout).After(driver.NewDateTime(status.OffWorkCheckIn.Format(layout), layout)) {
				status.OffWorkAttendStatus = enum.OffWorkEarly
			} else {
				status.OffWorkAttendStatus = enum.OffWorkNormal
				status.OffWorkAttendProcStatus = enum.ProcNormal
			}
		}

		// 缺勤計算
		var (
			absenceHours float32
		)

		if status.WorkCheckIn.IsZero() || status.OffWorkCheckIn.IsZero() {
			// 如果上下班其中一個沒打卡，缺勤時數=整天
			absenceHours, err = qWorkSchedule.WithContext(c.ctx).WorkHourCount(status.EmployeeID, basicWork.Format(), basicOffWork.Format())
			if err != nil {
				return err
			}
		} else {
			// 如果都有打卡
			// 遲到時間+早退時間
			var temp float32

			if status.WorkAttendStatus == enum.WorkLate {
				// late
				temp, err = qWorkSchedule.WithContext(c.ctx).WorkHourCount(status.EmployeeID, basicWork.Format(), status.WorkCheckIn.Format())
				if err != nil {
					return err
				}
				absenceHours += temp
			}

			if status.OffWorkAttendStatus == enum.OffWorkEarly {
				// early
				temp, err = qWorkSchedule.WithContext(c.ctx).WorkHourCount(status.EmployeeID, status.OffWorkCheckIn.Format(), basicOffWork.Format())
				if err != nil {
					return err
				}
				absenceHours += temp
			}
		}
		status.AbsenceHours = absenceHours

		// 減目前請假時間
		status.AbsenceHours -= status.LeaveHours

		// 更新狀態
		if status.AbsenceHours <= 0 {
			if status.WorkAttendProcStatus != enum.ProcNormal {
				status.WorkAttendProcStatus = enum.Processed
			}
			if status.OffWorkAttendProcStatus != enum.ProcNormal {
				status.OffWorkAttendProcStatus = enum.Processed
			}
		}

		// save
		return qCheckInStatus.WithContext(dbcache.WithCtx(c.ctx)).Save(status)
	})
}
