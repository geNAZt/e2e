package keyboard

import (
	"github.com/go-vgo/robotgo"
	lua "github.com/yuin/gopher-lua"
	"gomint.io/e2e/test/logger"
	"time"
)

func Press(L *lua.LState) int {
	start := time.Now()
	key := L.ToString(1)
	robotgo.KeyTap(key)

	logger.Debug("Pressing key '%s'", key)
	logger.Benchmark(start)
	return 0
}
