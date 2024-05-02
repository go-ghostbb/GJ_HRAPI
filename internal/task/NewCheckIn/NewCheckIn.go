package NewCheckIn

import (
	"context"
	"fmt"
	"ghostbb.io/gb/frame/g"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"hrapi/internal/query"
	"hrapi/internal/types"
	"hrapi/internal/types/enum"
	"hrapi/internal/utils/driver"
	"math"
	"time"
)

func Start(ctx context.Context) {
	g.Log().Info(ctx, "執行班表同步至打卡資料程序")

	// 開啟事務
	var err error
	tx := query.Q.Begin()
	defer func() {
		// 只要發生panic或有error, 直接rollback
		if recover() != nil || err != nil {
			_ = tx.Rollback()
		}
	}()

	var (
		todayStr       = time.Now().Format(time.DateOnly)
		today, _       = time.Parse(time.DateOnly, todayStr)
		qWorkSchedule  = tx.WorkSchedule
		qCheckInStatus = tx.CheckInStatus
		qEmployee      = tx.Employee
		schedules      []*types.WorkSchedule
		checkInStatus  = make([]*types.CheckInStatus, 0)
	)

	// 查詢今天班表
	schedules, err = qWorkSchedule.WithContext(ctx).Preload(qWorkSchedule.WorkShift).Where(qWorkSchedule.ScheduleDate.Eq(driver.NewDate(todayStr))).Find()
	if err != nil {
		g.Log().Error(ctx, "create new check_in status err:", err)
		return
	}

	// 建立employee和schedule map表
	// map[empID]*types.WorkSchedule
	var empScheMap = make(map[uint]*types.WorkSchedule)
	for _, schedule := range schedules {
		empScheMap[schedule.EmployeeID] = schedule
	}

	// 查詢所有在職員工id
	var empIDs []uint
	err = qEmployee.WithContext(ctx).Select(qEmployee.ID).Where(qEmployee.EmploymentStatus.Eq(enum.Active)).Scan(&empIDs)
	if err != nil {
		g.Log().Error(ctx, "create new check_in status err:", err)
		return
	}

	// 遍歷所有員工
	for _, empID := range empIDs {
		// 查詢該名員工是否有班表
		if v, ok := empScheMap[empID]; ok {
			// 存在
			// 計算預計下班日期
			offWork := gbconv.Time(fmt.Sprintf("%s %s", v.ScheduleDate.Format(), v.WorkShift.WorkStart.Format()))
			// 換算分鐘並四捨五入, 為了精準
			mins := time.Duration(math.Round(v.WorkShift.TotalHours * 60))
			// 加上上班總時數
			offWork = offWork.Add(mins * time.Minute)

			checkInStatus = append(checkInStatus, &types.CheckInStatus{
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
			checkInStatus = append(checkInStatus, &types.CheckInStatus{
				EmployeeID:              empID,
				Date:                    driver.Date(today.Unix()),
				WorkAttendStatus:        enum.WorkOffDay,    // 休息日
				WorkAttendProcStatus:    enum.ProcNormal,    // 處理狀態, 正常
				OffWorkAttendStatus:     enum.OffWorkOffDay, // 休息日
				OffWorkAttendProcStatus: enum.ProcNormal,    // 處理狀態, 正常
			})
		}
	}

	// 刪除原有, 避免資料重複
	if _, err = qCheckInStatus.WithContext(ctx).Where(qCheckInStatus.Date.Eq(driver.NewDate(today))).Delete(); err != nil {
		g.Log().Error(ctx, "create new check_in status err:", err)
		return
	}

	// 建立
	if err = qCheckInStatus.WithContext(ctx).CreateInBatches(checkInStatus, 100); err != nil {
		g.Log().Error(ctx, "create new check_in status err:", err)
		return
	}

	// commit
	if err = tx.Commit(); err != nil {
		g.Log().Error(ctx, "create new check_in status err:", err)
	}
}
