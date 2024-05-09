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
	cardNum int = iota
	dateTime
	checkInType
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
		qCheckInStatus = tx.CheckInStatus
	)
	// 更新表裡時間
	for _, row := range rows {
		var (
			fullDateTime = gbconv.Time(row[dateTime]).Format(time.DateTime)
			isWork       bool
		)

		switch row[checkInType] {
		case "1":
			// 上班
			isWork = true
		case "2":
			// 下班
			isWork = false
		}

		err = qCheckInStatus.WithContext(ctx).UpdateTime(fullDateTime, row[cardNum], isWork)
		if err != nil {
			g.Log().Error(ctx, "sync check_in data err:", err)
			return
		}
	}

	// 計算狀態
	//if err = qCheckInStatus.WithContext(ctx).UpdateStatus(gbstr.Join(updateDates, ",")); err != nil {
	//	g.Log().Error(ctx, "sync check_in data err:", err)
	//	return
	//}

	// 提交(commit)
	if err = tx.Commit(); err != nil {
		g.Log().Error(ctx, "sync check_in data err:", err)
	}
}
