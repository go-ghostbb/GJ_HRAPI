package cmd

import (
	"context"
	"fmt"
	gbcmd "ghostbb.io/gb/os/gb_cmd"
	"github.com/goccy/go-json"
	"time"
)

func testFn(ctx context.Context, parser *gbcmd.Parser) (err error) {
	r, _ := json.Marshal(time.Now())
	fmt.Print(string(r))
	return
}
