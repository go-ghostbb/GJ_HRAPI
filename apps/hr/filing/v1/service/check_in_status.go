package service

import (
	"context"
	"fmt"
	"ghostbb.io/gb/frame/g"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"hrapi/apps/hr/filing/v1/model"
	"hrapi/internal/query"
	"hrapi/internal/types"
	"hrapi/internal/types/enum"
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
		schedules, err = qWorkSchedule.WithContext(c.ctx).Preload(qWorkSchedule.WorkShift).QueryByDateRange(in.DateRange[0], in.DateRange[1])
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
			empScheMap[schedule.EmployeeID][schedule.ScheduleDate.Format(time.DateOnly)] = schedule
		}

		// 查詢所有在職員工id
		var empIDs []uint
		err = qEmployee.WithContext(c.ctx).Select(qEmployee.ID).Where(qEmployee.EmploymentStatus.Eq(enum.Active)).Scan(&empIDs)
		if err != nil {
			return err
		}

		// 遍歷時間區間
		startDate := gbconv.Time(in.DateRange[0])
		endDate := gbconv.Time(in.DateRange[1])
		for endDate.After(startDate) {
			tmpDate := startDate.Format(time.DateOnly) // 暫存日期(string)
			// 遍歷所有員工
			for _, empID := range empIDs {
				// 查詢該名員工是否有班表
				if v, ok := empScheMap[empID][tmpDate]; ok {
					// 存在
					// 計算預計下班日期
					offWork := gbconv.Time(fmt.Sprintf("%s %s", v.ScheduleDate.Format(time.DateOnly), v.WorkShift.WorkStart.Format(time.TimeOnly)))
					// 換算分鐘並四捨五入, 為了精準
					mins := time.Duration(math.Round(v.WorkShift.TotalHours * 60))
					// 加上上班總時數
					offWork = offWork.Add(mins * time.Minute)

					newCheckInStatus = append(newCheckInStatus, &types.CheckInStatus{
						EmployeeID:              empID,
						WorkShiftID:             v.WorkShiftID,
						WorkCheckInDate:         v.ScheduleDate,        // 上班時間
						WorkAttendStatus:        enum.WorkNotSwiped,    // 上班狀態, 未刷卡
						WorkAttendProcStatus:    enum.NotProcessed,     // 處理狀態, 未處理
						OffWorkCheckInDate:      offWork,               // 下班時間
						OffWorkAttendStatus:     enum.OffWorkNotSwiped, // 下班狀態, 未刷卡
						OffWorkAttendProcStatus: enum.NotProcessed,     // 處理狀態, 未處理
						AbsenceHours:            float32(v.WorkShift.TotalHours),
					})
				} else {
					// 不存在, 代表休息日
					newCheckInStatus = append(newCheckInStatus, &types.CheckInStatus{
						EmployeeID:              empID,
						WorkCheckInDate:         startDate,
						WorkAttendStatus:        enum.WorkOffDay, // 休息日
						WorkAttendProcStatus:    enum.ProcNormal, // 處理狀態, 正常
						OffWorkCheckInDate:      startDate,
						OffWorkAttendStatus:     enum.OffWorkOffDay, // 休息日
						OffWorkAttendProcStatus: enum.ProcNormal,    // 處理狀態, 正常
					})
				}
			}

			startDate = startDate.AddDate(0, 0, 1)
		}

		// 刪除原有, 避免資料重複
		if _, err = qCheckInStatus.WithContext(c.ctx).DeleteByDateRange(in.DateRange[0], in.DateRange[1]); err != nil {
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
			err             error
			qCheckInStatus  = tx.CheckInStatus
			updateDates     = make([]string, 0)
			checkUpdateDate = make(map[string]struct{})
		)

		// 遍歷+更新
		for _, data := range in {
			var (
				fullDateTime = gbconv.Time(data.DateTime)
				dateOnly     = fullDateTime.Format(time.DateOnly)
				timeOnly     = fullDateTime.Format(time.TimeOnly)
				rowsAffected int64
			)

			switch data.CheckInType {
			case "1":
				// 上班
				rowsAffected, err = qCheckInStatus.WithContext(c.ctx).UpdateTime(timeOnly, dateOnly, data.WorkShiftCode, data.CardNumber, true)
			case "2":
				// 下班
				rowsAffected, err = qCheckInStatus.WithContext(c.ctx).UpdateTime(timeOnly, dateOnly, data.WorkShiftCode, data.CardNumber, false)
			}

			if err != nil {
				return err
			}
			if rowsAffected == 0 {
				// 如果更新數量為0, warning
				g.Log().Warningf(c.ctx, "%s %s not found", dateOnly, data.CardNumber)
			} else {
				if _, ok := checkUpdateDate[dateOnly]; !ok {
					updateDates = append(updateDates, dateOnly)
					checkUpdateDate[dateOnly] = struct{}{}
				}
			}
		}

		// 計算狀態
		if _, err = qCheckInStatus.WithContext(c.ctx).UpdateStatus(updateDates); err != nil {
			return err
		}

		// commit
		return nil
	})
}
