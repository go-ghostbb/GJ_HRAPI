package NewCheckIn

import (
	"context"
	"ghostbb.io/gb/frame/g"
	"hrapi/apps/hr/filing/v1/model"
	"hrapi/apps/hr/filing/v1/service"
	"time"
)

func Start(ctx context.Context) {
	g.Log().Info(ctx, "執行班表同步至打卡資料程序")

	err := service.CheckInStatus(ctx).Filing(model.FilingCheckInStatusReq{
		DateRange: []string{time.Now().Format(time.DateOnly), time.Now().Format(time.DateOnly)},
	})
	if err != nil {
		g.Log().Errorf(ctx, "NewCheckIn error: %s", err.Error())
		return
	}

	g.Log().Info(ctx, "NewCheckIn done!")
}
