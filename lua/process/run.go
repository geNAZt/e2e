package process

import (
	lua "github.com/yuin/gopher-lua"
	"gomint.io/e2e/test/logger"
	"os/exec"
	"time"
)

func Run(L *lua.LState) int {
	start := time.Now()
	cmdParts, cmdline := ResolveCmdParts(L, 1)

	logger.Debug("Running command '%s'", cmdline)

	cmd := exec.Command(cmdParts[0], cmdParts[1:]...)
	if err := cmd.Start(); err != nil {
		logger.Failure(err, "Could not run program '%s'", cmdline)
	}

	L.Push(lua.LNumber(cmd.Process.Pid))

	logger.Benchmark(start)
	return 1
}
