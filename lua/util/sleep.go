package util

import (
	lua "github.com/yuin/gopher-lua"
	"gomint.io/e2e/test/logger"
	"time"
)

func Sleep(L *lua.LState) int {
	millis := L.ToInt64(1)
	logger.Debug("Sleeping for '%d'", millis)
	time.Sleep(time.Duration(millis) * time.Millisecond)
	return 0
}
