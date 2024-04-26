package task

import (
	"context"
	"ghostbb.io/gb/frame/g"
	gbcron "ghostbb.io/gb/os/gb_cron"
	"hrapi/internal/task/NewCheckIn"
	"hrapi/internal/task/SyncCheckIn"
)

func Init() {
	// 定時任務最好都加上recover()捕捉, 避免在定時任務啟動後, bug引起panic導致程序崩潰

	var (
		err error
		ctx context.Context
	)

	// 使用系統logger
	gbcron.SetLogger(g.Log())

	// 每天將班表加入需打卡表裡
	_, err = gbcron.Add(ctx, "@daily", NewCheckIn.Start, "new check_in")
	if err != nil {
		panic(err)
	}

	// 每天固定10點同步打卡機資料
	_, err = gbcron.Add(ctx, "0 0 10 * * *", SyncCheckIn.Start, "sync check_in")
	if err != nil {
		panic(err)
	}
}
