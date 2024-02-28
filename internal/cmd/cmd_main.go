package cmd

import (
	"context"
	"ghostbb.io/gb/contrib/dbcache"
	"ghostbb.io/gb/frame/g"
	gbcmd "ghostbb.io/gb/os/gb_cmd"
	"hrapi/apps/system"
	"hrapi/internal/middleware"
	"hrapi/internal/query"
	"hrapi/internal/utils"
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
	g.Server().Bind(system.New()...)

	// Register static resource
	utils.MkdirIfNotExist("assets")
	s.Static("/assets", "./assets")

	// 執行
	s.Run()
	return nil
}
