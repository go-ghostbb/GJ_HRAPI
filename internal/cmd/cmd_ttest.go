package cmd

import (
	"context"
	"ghostbb.io/gb/contrib/dbcache"
	"ghostbb.io/gb/frame/g"
	gbcmd "ghostbb.io/gb/os/gb_cmd"
	gbctx "ghostbb.io/gb/os/gb_ctx"
	"hrapi/internal/query"
	"hrapi/internal/task/UpdateLeaveCount"
)

func testFn(ctx context.Context, parser *gbcmd.Parser) (err error) {
	// 設定資料庫緩存
	databaseCache := dbcache.New(dbcache.NewRedis(g.Redis("dbcache")))
	if err = g.DB().Use(databaseCache); err != nil {
		return err
	}

	// 設定查詢DB
	query.SetDefault(g.DB().DB)

	leaveCheck := &UpdateLeaveCount.LeaveCheck{Q: query.Q, Log: g.Log()}
	leaveCheck.Run(gbctx.New())
	return
}
