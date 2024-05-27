package cmd

import (
	"context"
	"ghostbb.io/gb/contrib/dbcache"
	"ghostbb.io/gb/frame/g"
	gbcmd "ghostbb.io/gb/os/gb_cmd"
	"github.com/gin-gonic/gin"
	"hrapi/apps/hr/daily"
	"hrapi/apps/hr/filing"
	queryApp "hrapi/apps/hr/query"
	"hrapi/apps/hr/salary"
	"hrapi/apps/hr/setting"
	signoff "hrapi/apps/hr/sign-off"
	"hrapi/apps/system"
	"hrapi/apps/upload"
	"hrapi/internal/middleware"
	"hrapi/internal/query"
	"hrapi/internal/task"
	"hrapi/internal/utils"
	"hrapi/internal/utils/identity"
	"hrapi/internal/utils/response"
)

func mainFn(ctx context.Context, parser *gbcmd.Parser) (err error) {
	// 設定資料庫緩存
	databaseCache := dbcache.New(dbcache.NewRedis(g.Redis("dbcache")))
	if err = g.DB().Use(databaseCache); err != nil {
		return err
	}

	// 設定查詢DB
	query.SetDefault(g.DB().DB)

	// 註冊全域server中間件
	s := g.Server()
	g.Server().Use(middleware.Cors(), middleware.I18n())
	g.Server().SetRecoveryFn(response.RecoveryFn)

	// TODO: bind
	g.Server().Bind(
		system.New(),
		upload.New(),
		setting.New(),
		filing.New(),
		salary.New(),
		daily.New(),
		signoff.New(),
		queryApp.New(),
	)

	// Register static resource
	utils.MkdirIfNotExist("assets")
	s.Static("/assets", "./assets")

	// 健康檢查
	g.Server().GET("health", func(c *gin.Context) {
		c.JSON(200, g.Map{
			"code": 0,
			"data": g.Map{},
			"msg":  "ok",
		})
	})

	// 定時任務
	task.Init()

	// 併發可用，自動遞增單號或員工編號等......
	if err = identity.Init(); err != nil {
		return err
	}

	// 執行
	s.Run()
	return nil
}
