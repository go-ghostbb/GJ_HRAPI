package NewCheckIn

import (
	"context"
	"fmt"
	"ghostbb.io/gb/frame/g"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"hrapi/internal/query"
	"hrapi/internal/types"
	"hrapi/internal/types/enum"
	"math"
	"time"
)

func Start(ctx context.Context) {
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
		schedules      []*types.WorkSchedule
		checkInStatus  = make([]*types.CheckInStatus, 0)
	)

	// 查詢今天班表
	schedules, err = qWorkSchedule.WithContext(ctx).Preload(qWorkSchedule.WorkShift).QueryByDate(todayStr)
	if err != nil {
		g.Log().Error(ctx, "create new check_in status err:", err)
		return
	}

	for _, schedule := range schedules {
		// 計算預計下班日期
		offWork := gbconv.Time(fmt.Sprintf("%s %s", todayStr, schedule.WorkShift.WorkStart.Format(time.TimeOnly)))
		// 換算分鐘並四捨五入, 為了精準
		mins := time.Duration(math.Round(schedule.WorkShift.TotalHours * 60))
		// 加上上班總時數
		offWork = offWork.Add(mins * time.Minute)

		checkInStatus = append(checkInStatus, &types.CheckInStatus{
			EmployeeID:              schedule.EmployeeID,
			WorkShiftID:             schedule.WorkShiftID,
			WorkCheckInDate:         today,                 // 上班時間
			WorkAttendStatus:        enum.WorkNotSwiped,    // 上班狀態, 未刷卡
			WorkAttendProcStatus:    enum.NotProcessed,     // 處理狀態, 未處理
			OffWorkCheckInDate:      offWork,               // 下班時間
			OffWorkAttendStatus:     enum.OffWorkNotSwiped, // 下班狀態, 未刷卡
			OffWorkAttendProcStatus: enum.NotProcessed,     // 處理狀態, 未處理
			AbsenceHours:            float32(schedule.WorkShift.TotalHours),
		})
	}
	// 刪除原有, 避免資料重複
	if _, err = qCheckInStatus.WithContext(ctx).DeleteByDate(todayStr); err != nil {
		g.Log().Error(ctx, "create new check_in status err:", err)
		return
	}

	// 建立
	if err = qCheckInStatus.WithContext(ctx).Create(checkInStatus...); err != nil {
		g.Log().Error(ctx, "create new check_in status err:", err)
		return
	}

	// commit
	if err = tx.Commit(); err != nil {
		g.Log().Error(ctx, "create new check_in status err:", err)
	}
}
