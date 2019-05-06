package log

import (
	lua "github.com/yuin/gopher-lua"
	"gomint.io/e2e/test/logger"
)

func Info(L *lua.LState) int {
	logger.Info(L.ToString(1))
	return 0
}
