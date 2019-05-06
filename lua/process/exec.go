package process

import (
	lua "github.com/yuin/gopher-lua"
	"gomint.io/e2e/test/logger"
	"os/exec"
	"strings"
	"time"
)

func Exec(L *lua.LState) int {
	start := time.Now()
	cmdline := L.ToString(1)
	cmdParts := strings.Split(cmdline, " ")

	logger.Debug("Executing command '%s'", cmdline)

	cmd := exec.Command(cmdParts[0], cmdParts[1:]...)
	if err := cmd.Run(); err != nil {
		logger.Failure(err, "Could not run program '%s'", cmdline)
	}

	logger.Benchmark(start)
	return 0
}
