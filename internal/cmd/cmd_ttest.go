package cmd

import (
	"context"
	"ghostbb.io/gb/contrib/dbcache"
	"ghostbb.io/gb/frame/g"
	gbcmd "ghostbb.io/gb/os/gb_cmd"
	"hrapi/internal/query"
)

func testFn(ctx context.Context, parser *gbcmd.Parser) (err error) {
	// 設定資料庫緩存
	databaseCache := dbcache.New(dbcache.NewRedis(g.Redis("dbcache")))
	if err = g.DB().Use(databaseCache); err != nil {
		return err
	}

	// 設定查詢DB
	query.SetDefault(g.DB().DB)
	println("test")
	return
}
