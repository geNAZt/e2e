package mouse

import (
	"github.com/go-vgo/robotgo"
	lua "github.com/yuin/gopher-lua"
	"gomint.io/e2e/test/logger"
	"time"
)

func Click(L *lua.LState) int {
	start := time.Now()
	robotgo.MouseClick()
	logger.Debug("Clicking with the mouse")
	logger.Benchmark(start)
	return 0
}
