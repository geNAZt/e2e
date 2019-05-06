package process

import (
	lua "github.com/yuin/gopher-lua"
	"gomint.io/e2e/test/logger"
	"os/exec"
	"strings"
	"time"
)

func ResolveCmdParts(L *lua.LState, index int) (cmdParts []string, cmdline string) {
	cmdPartsT := L.ToTable(index)
	if cmdPartsT != nil {
		cmdParts = make([]string, cmdPartsT.Len())
		i := 0
		cmdPartsT.ForEach(func (k lua.LValue, v lua.LValue) {
			cmdParts[i] = v.String()
			i++
		})
		cmdline = strings.Join(cmdParts, " ")
	} else {
		cmdline = L.ToString(index)
		cmdParts = strings.Split(cmdline, " ")
	}
	return
}

func Exec(L *lua.LState) int {
	start := time.Now()
	cmdParts, cmdline := ResolveCmdParts(L, 1)

	logger.Debug("Executing command '%s'", cmdline)
	cmd := exec.Command(cmdParts[0], cmdParts[1:]...)
	if err := cmd.Run(); err != nil {
		logger.Failure(err, "Could not run program '%s'", cmdline)
	}

	logger.Benchmark(start)
	return 0
}
