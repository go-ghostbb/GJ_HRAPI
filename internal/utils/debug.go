package utils

import (
	gbcmd "ghostbb.io/gb/os/gb_cmd"
)

const (
	// Debug key for checking if in debug mode.
	commandEnvKeyForDebugKey = "hr.debug"
)

var (
	// isDebugEnabled marks whether GoFrame debug mode is enabled.
	isDebugEnabled = false
)

func init() {
	// Debugging configured.
	value := gbcmd.GetOptWithEnv(commandEnvKeyForDebugKey).String()
	if value == "" || value == "0" || value == "false" {
		isDebugEnabled = false
		//g.Cfg().GetAdapter().(*gbcfg.AdapterFile).SetFileName("config.prod.yaml")
	} else {
		isDebugEnabled = true
		//g.Cfg().GetAdapter().(*gbcfg.AdapterFile).SetFileName("config.dev.yaml")
	}
}

func IsDebugEnabled() bool {
	return isDebugEnabled
}
