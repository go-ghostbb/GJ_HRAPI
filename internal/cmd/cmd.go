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
	Gen = &gbcmd.Command{
		Name:  "gen",
		Usage: "gen",
		Brief: "start gen",
		Func:  genFn,
	}
	Test = &gbcmd.Command{
		Name:  "test",
		Usage: "test",
		Brief: "test run",
		Func:  testFn,
	}
)
