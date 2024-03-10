package main

import (
	_ "ghostbb.io/gb/contrib/drivers/mssql"
	_ "ghostbb.io/gb/contrib/nosql/redis"
	"hrapi/internal/cmd"

	_ "hrapi/internal/packed"

	gbctx "ghostbb.io/gb/os/gb_ctx"
)

func main() {
	if err := cmd.Main.AddCommand(cmd.Test); err != nil {
		panic(err)
	}
	cmd.Main.Run(gbctx.GetInitCtx())
}
