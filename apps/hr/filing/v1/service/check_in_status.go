package service

import (
	"context"
	"fmt"
	"ghostbb.io/gb/contrib/dbcache"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"gorm.io/gen"
	"hrapi/apps/hr/filing/v1/model"
	"hrapi/internal/query"
	"hrapi/internal/types"
	"hrapi/internal/types/enum"
	"hrapi/internal/utils/driver"
	"math"
	"time"
)

func CheckInStatus(ctx context.Context) ICheckInStatus {
	return &checkInStatus{ctx: ctx}
}

type (
	ICheckInStatus interface {
		// Filing 輸入的日期區間, 將班表寫進打卡狀態裡
		Filing(in model.FilingCheckInStatusReq) error
		// UploadData 上傳資料
		UploadData(in []*model.UploadDataReq) error
	}

	checkInStatus struct {
		ctx context.Context
	}
)

// Filing 輸入的日期區間, 將班表寫進打卡狀態裡
func (c *checkInStatus) Filing(in model.FilingCheckInStatusReq) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			qWorkSchedule    = tx.WorkSchedule
			qCheckInStatus   = tx.CheckInStatus
			qEmployee        = tx.Employee
			schedules        []*types.WorkSchedule
			newCheckInStatus = make([]*types.CheckInStatus, 0)
			err              error
		)

		// 查詢區間內所有班表
		schedules, err = qWorkSchedule.WithContext(dbcache.WithCtx(c.ctx)).
			Preload(qWorkSchedule.WorkShift).
			Where(qWorkSchedule.ScheduleDate.Gte(driver.NewDate(in.DateRange[0])), qWorkSchedule.ScheduleDate.Lte(driver.NewDate(in.DateRange[1]))).
			Find()
		if err != nil {
			return err
		}

		// 建立employee和schedule map表
		// map[empID]map[date]*types.WorkSchedule
		var empScheMap = make(map[uint]map[string]*types.WorkSchedule)
		for _, schedule := range schedules {
			if _, ok := empScheMap[schedule.EmployeeID]; !ok {
				// 不存在empID map
				empScheMap[schedule.EmployeeID] = make(map[string]*types.WorkSchedule)
			}
			empScheMap[schedule.EmployeeID][schedule.ScheduleDate.Format()] = schedule
		}

		// 查詢所有在職員工id
		if in.EmployeeID == nil || len(in.EmployeeID) == 0 {
			err = qEmployee.WithContext(c.ctx).Select(qEmployee.ID).Where(qEmployee.EmploymentStatus.Eq(enum.Active)).Scan(&in.EmployeeID)
			if err != nil {
				return err
			}
		}

		// 遍歷時間區間
		startDate := driver.NewDate(in.DateRange[0])
		endDate := driver.NewDate(in.DateRange[1]).AddDate(0, 0, 1)
		for endDate.After(startDate) {
			tmpDate := startDate.Format() // 暫存日期(string)
			// 遍歷所有員工
			for _, empID := range in.EmployeeID {
				// 查詢該名員工是否有班表
				if v, ok := empScheMap[empID][tmpDate]; ok {
					// 存在
					// 計算預計下班日期
					offWork := gbconv.Time(fmt.Sprintf("%s %s", v.ScheduleDate.Format(), v.WorkShift.WorkStart.Format()))
					// 換算分鐘並四捨五入, 為了精準
					mins := time.Duration(math.Round(v.WorkShift.TotalHours * 60))
					// 加上上班總時數
					offWork = offWork.Add(mins * time.Minute)

					newCheckInStatus = append(newCheckInStatus, &types.CheckInStatus{
						EmployeeID:              empID,
						WorkShiftID:             v.WorkShiftID,
						Date:                    driver.Date(v.ScheduleDate.Unix()), // 上班時間
						WorkAttendStatus:        enum.WorkNotSwiped,                 // 上班狀態, 未刷卡
						WorkAttendProcStatus:    enum.NotProcessed,                  // 處理狀態, 未處理
						OffWorkAttendStatus:     enum.OffWorkNotSwiped,              // 下班狀態, 未刷卡
						OffWorkAttendProcStatus: enum.NotProcessed,                  // 處理狀態, 未處理
						AbsenceHours:            float32(v.WorkShift.TotalHours),
					})
				} else {
					// 不存在, 代表休息日
					newCheckInStatus = append(newCheckInStatus, &types.CheckInStatus{
						EmployeeID:              empID,
						Date:                    startDate,
						WorkAttendStatus:        enum.WorkOffDay,    // 休息日
						WorkAttendProcStatus:    enum.ProcNormal,    // 處理狀態, 正常
						OffWorkAttendStatus:     enum.OffWorkOffDay, // 休息日
						OffWorkAttendProcStatus: enum.ProcNormal,    // 處理狀態, 正常
					})
				}
			}

			startDate = startDate.AddDate(0, 0, 1)
		}

		// 刪除原有, 避免資料重複
		// in.DateRange[0] <= date <= in.DateRange[1]
		var whereSQL = []gen.Condition{
			qCheckInStatus.Date.Gte(driver.NewDate(in.DateRange[0])),
			qCheckInStatus.Date.Lte(driver.NewDate(in.DateRange[1])),
		}
		if in.EmployeeID != nil && len(in.EmployeeID) > 0 {
			whereSQL = append(whereSQL, qCheckInStatus.EmployeeID.In(in.EmployeeID...))
		}
		if _, err = qCheckInStatus.WithContext(c.ctx).Unscoped().Where(whereSQL...).Delete(); err != nil {
			return err
		}

		// 建立
		if err = qCheckInStatus.WithContext(c.ctx).CreateInBatches(newCheckInStatus, 100); err != nil {
			return err
		}

		// commit
		return nil
	})
}

// UploadData 上傳資料
func (c *checkInStatus) UploadData(in []*model.UploadDataReq) error {
	return query.Q.Transaction(func(tx *query.Query) error {
		var (
			err            error
			qCheckInStatus = tx.CheckInStatus
		)

		// 遍歷+更新
		for _, data := range in {
			var (
				fullDateTime = gbconv.Time(data.DateTime).Format(time.DateTime)
				ttype        enum.CheckInType
			)

			switch data.CheckInType {
			case "0":
				// 上班
				ttype = enum.Work
			case "1":
				// 下班
				ttype = enum.OffWork
			}

			err = qCheckInStatus.WithContext(c.ctx).UpdateTime(fullDateTime, data.CardNumber, string(ttype))
			if err != nil {
				return err
			}
		}

		// commit
		return nil
	})
}
