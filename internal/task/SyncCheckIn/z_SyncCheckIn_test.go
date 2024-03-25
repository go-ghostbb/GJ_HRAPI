package SyncCheckIn_test

import (
	_ "ghostbb.io/gb/contrib/drivers/mssql"
	"ghostbb.io/gb/frame/g"
	gbcfg "ghostbb.io/gb/os/gb_cfg"
	gbctx "ghostbb.io/gb/os/gb_ctx"
	gbtest "ghostbb.io/gb/test/gb_test"
	"hrapi/internal/query"
	"hrapi/internal/task/SyncCheckIn"
	"testing"
)

func TestStart(t *testing.T) {
	gbtest.C(t, func(t *gbtest.T) {
		g.Cfg().GetAdapter().(*gbcfg.AdapterFile).SetFileName(gbtest.DataPath("config.test.yaml"))
		query.SetDefault(g.DB().DB)
		SyncCheckIn.Start(gbctx.New())
	})
}
