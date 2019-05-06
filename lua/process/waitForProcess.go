package process

import (
	"github.com/go-vgo/robotgo"
	lua "github.com/yuin/gopher-lua"
	"gomint.io/e2e/test/logger"
	"time"
)

func WaitForProcess(L *lua.LState) int {
	start := time.Now()

	processName := L.ToString(1)

	logger.Debug("Waiting for process '%s' to show up", processName)

	for {
		ids, err := robotgo.FindIds(processName)
		if err != nil {
			logger.Failure(err, "Error in finding pid by name '%s'", processName)
		}

		if len(ids) > 0 {
			L.Push(lua.LNumber(ids[0]))
			logger.Benchmark(start)
			return 1
		}
	}
}
