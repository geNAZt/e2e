package keyboard

import (
	"github.com/go-vgo/robotgo"
	lua "github.com/yuin/gopher-lua"
	"gomint.io/e2e/test/logger"
	"math/rand"
	"time"
)

func Type(L *lua.LState) int {
	start := time.Now()
	toType := L.ToString(1)

	if L.Get(2) != lua.LNil {
		logger.Debug("Typing '%s' slowly", toType)

		for _, v := range toType {
			if string(v) == "_" {
				robotgo.KeyToggle("lshift", "down")
			}

			robotgo.TypeString(string(v))
			time.Sleep(time.Duration(37+rand.Int31n(100)) * time.Millisecond)

			if string(v) == "_" {
				robotgo.KeyToggle("lshift", "up")
			}
		}
	} else {
		logger.Debug("Typing '%s'", toType)
		robotgo.TypeStr(toType)
	}

	logger.Benchmark(start)
	return 0 // No result
}
