package SyncCheckIn

import (
	"context"
	"ghostbb.io/gb/frame/g"
	gbenv "ghostbb.io/gb/os/gb_env"
	gbconv "ghostbb.io/gb/util/gb_conv"
	"github.com/xuri/excelize/v2"
	"hrapi/internal/query"
	"time"
)

func init() {
	result := gbenv.GetWithCmd("hr.check.in.data.path")
	if result.String() != "" {
		dataPath = result.String()
	}
}

var dataPath = ""

const (
	cardNum       = 0
	dateTime      = 1
	workShiftCode = 2
	checkInType   = 3
)

func Start(ctx context.Context) {
	g.Log().Info(ctx, "執行同步打卡程序")

	if dataPath == "" {
		g.Log().Error(ctx, "path is empty")
		return
	}

	// 讀取excel
	excelFile, err := excelize.OpenFile(dataPath)
	if err != nil {
		g.Log().Error(ctx, "sync check_in data err:", err)
		return
	}
	defer func() {
		excelFile.Close()
	}()

	// 從這邊開啟db事務
	tx := query.Q.Begin()
	defer func() {
		// 只要發生panic或有error, 直接rollback
		if recover() != nil || err != nil {
			_ = tx.Rollback()
		}
	}()

	var (
		rows [][]string
	)

	// 獲取data上所有儲存格
	rows, err = excelFile.GetRows("data")
	if err != nil {
		g.Log().Error(ctx, "sync check_in data err:", err)
		return
	}
	rows = rows[1:] // 第一列為header, 直接丟棄

	// 更新
	var (
		qCheckInStatus  = tx.CheckInStatus
		updateDates     = make([]string, 0)
		checkUpdateDate = make(map[string]struct{})
	)
	// 更新表裡時間
	for _, row := range rows {
		var (
			fullDateTime = gbconv.Time(row[dateTime])
			dateOnly     = fullDateTime.Format(time.DateOnly)
			timeOnly     = fullDateTime.Format(time.TimeOnly)
			rowsAffected int64
		)

		switch row[checkInType] {
		case "1":
			// 上班
			rowsAffected, err = qCheckInStatus.WithContext(ctx).UpdateTime(timeOnly, dateOnly, row[workShiftCode], row[cardNum], true)
		case "2":
			// 下班
			rowsAffected, err = qCheckInStatus.WithContext(ctx).UpdateTime(timeOnly, dateOnly, row[workShiftCode], row[cardNum], false)
		}

		if err != nil {
			g.Log().Error(ctx, "sync check_in data err:", err)
			return
		}
		if rowsAffected == 0 {
			// 如果更新數量為0, warning
			g.Log().Warningf(ctx, "%s %s not found", dateOnly, row[0])
		} else {
			if _, ok := checkUpdateDate[dateOnly]; !ok {
				updateDates = append(updateDates, dateOnly)
				checkUpdateDate[dateOnly] = struct{}{}
			}
		}
	}

	// 計算狀態
	if _, err = qCheckInStatus.WithContext(ctx).UpdateStatus(updateDates); err != nil {
		g.Log().Error(ctx, "sync check_in data err:", err)
		return
	}

	// 提交(commit)
	if err = tx.Commit(); err != nil {
		g.Log().Error(ctx, "sync check_in data err:", err)
	}
}
