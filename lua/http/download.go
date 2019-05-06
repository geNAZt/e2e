package http

import (
	lua "github.com/yuin/gopher-lua"
	"gomint.io/e2e/test/logger"
	"net/http"
	"time"
	"os"
	"io"
)

func Download(L *lua.LState) int {
	start := time.Now()
	url := L.ToString(1)
	destPath := L.ToString(2)

	logger.Debug("Downloading '%s' to '%s'", url, destPath)

	resp, err := http.Get(url)
	if err != nil {
		logger.Failure(err, "Could not resolve url '%s'", url)
	}
	defer resp.Body.Close()

	out, err := os.Create(destPath)
	if err != nil {
		logger.Failure(err, "Could not create file '%s'", destPath)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		logger.Failure(err, "Could not save file '%s'", destPath)
	}

	logger.Benchmark(start)
	return 0 // No result
}
