// +build !linux

package window

import (
	lua "github.com/yuin/gopher-lua"
	"gomint.io/e2e/test/logger"
)

func GetWindows(L *lua.LState) int {
	logger.Error("GetWindows not supported on this platform")
	table := lua.LTable{}
	L.Push(&table)
	return 1
}