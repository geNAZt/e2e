package mouse

import (
	"github.com/go-vgo/robotgo"
	lua "github.com/yuin/gopher-lua"
	"gomint.io/e2e/test/logger"
	"time"
)

func Move(L *lua.LState) int {
	start := time.Now()

	x := L.ToInt(1)
	y := L.ToInt(2)

	robotgo.MoveMouse(x, y)

	logger.Debug("Moving mouse to '%d', '%d'", x, y)
	logger.Benchmark(start)
	return 0
}
