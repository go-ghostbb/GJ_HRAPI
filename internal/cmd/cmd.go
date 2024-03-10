package cmd

import (
	gbcmd "ghostbb.io/gb/os/gb_cmd"
)

var (
	Main = &gbcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func:  mainFn,
	}
	Test = &gbcmd.Command{
		Name:  "test",
		Usage: "test",
		Brief: "test run",
		Func:  testFn,
	}
)
