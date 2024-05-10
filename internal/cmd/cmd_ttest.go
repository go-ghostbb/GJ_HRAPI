package cmd

import (
	"context"
	gbcmd "ghostbb.io/gb/os/gb_cmd"
	"time"
)

func testFn(ctx context.Context, parser *gbcmd.Parser) (err error) {
	// 設定資料庫緩存
	//databaseCache := dbcache.New(dbcache.NewRedis(g.Redis("dbcache")))
	//if err = g.DB().Use(databaseCache); err != nil {
	//	return err
	//}

	// 設定查詢DB
	//query.SetDefault(g.DB().DB)
	tomorrow := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+1, 0, 0, 0, 0, time.Local)
	c := tomorrow.Sub(time.Now()).Hours()

	println(c)
	return
}
