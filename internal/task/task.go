package task

import (
	"ghostbb.io/gb/frame/g"
	gbcron "ghostbb.io/gb/os/gb_cron"
	gbctx "ghostbb.io/gb/os/gb_ctx"
	"hrapi/internal/task/NewCheckIn"
)

func Init() {
	// 定時任務最好都加上recover()捕捉, 避免在定時任務啟動後, bug引起panic導致程序崩潰

	var (
		err error
		ctx = gbctx.New()
	)

	// 使用系統logger
	gbcron.SetLogger(g.Log())

	// 每天將班表加入需打卡表裡
	_, err = gbcron.Add(ctx, "0 0 1 * * *", NewCheckIn.Start, "new check_in")
	if err != nil {
		panic(err)
	}
}
