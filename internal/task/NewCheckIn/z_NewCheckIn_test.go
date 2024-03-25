package NewCheckIn_test

import (
	"ghostbb.io/gb/frame/g"
	gbcfg "ghostbb.io/gb/os/gb_cfg"
	gbctx "ghostbb.io/gb/os/gb_ctx"
	gbtest "ghostbb.io/gb/test/gb_test"
	"hrapi/internal/query"
	NewCheckIn "hrapi/internal/task/NewCheckIn"
	"testing"

	_ "ghostbb.io/gb/contrib/drivers/mssql"
)

func TestStart(t *testing.T) {
	gbtest.C(t, func(t *gbtest.T) {
		cfgPath := gbtest.DataPath("config.test.yaml")
		g.Cfg().GetAdapter().(*gbcfg.AdapterFile).SetFileName(cfgPath)

		query.SetDefault(g.DB().DB)

		NewCheckIn.Start(gbctx.New())
	})
}
