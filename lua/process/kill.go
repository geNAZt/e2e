package process

import (
	lua "github.com/yuin/gopher-lua"
	"gomint.io/e2e/test/logger"
	"os"
	"time"
)

func Kill(L *lua.LState) int {
	start := time.Now()

	pid := L.ToInt(1)
	logger.Debug("Killing PID '%d'", pid)

	process, err := os.FindProcess(pid)
	if err != nil {
		logger.Failure(err, "Could not find PID '%d'", pid)
	}

	if err := process.Kill(); err != nil {
		logger.Failure(err, "Could not kill PID '%d'", pid)
	}

	logger.Benchmark(start)
	return 0
}
