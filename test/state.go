package test

import (
	"gomint.io/e2e/lua"
	"gomint.io/e2e/test/cache"
	"gomint.io/e2e/test/logger"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"time"
)

type State struct {
	cache            *cache.Cache
	currentTestState bool
}

func NewState() *State {
	return &State{
		cache: cache.NewCache(),
	}
}

func (state *State) Teardown() {
	// do we have a specific setup?
	if _, err := os.Stat("./tests/teardown/" + runtime.GOOS); !os.IsNotExist(err) {
		state.runTests("./tests/teardown/" + runtime.GOOS)
	} else {
		logger.Fatal("Could not find specific platform teardown for '%s'", runtime.GOOS)
	}
}

func (state *State) Setup() {
	// do we have a specific setup?
	if _, err := os.Stat("./tests/setup/" + runtime.GOOS); !os.IsNotExist(err) {
		state.runTests("./tests/setup/" + runtime.GOOS)
	} else {
		logger.Fatal("Could not find specific platform setup for '%s'", runtime.GOOS)
	}
}

func (state *State) runTests(folder string) {
	state.cache.Reset()
	state.cache.SetFolder(folder)

	files, _ := ioutil.ReadDir(folder)
	logger.TestFolderHeader(len(files), strings.Split(folder, "/")[2])
	state.currentTestState = true

	for _, v := range files {
		if v.IsDir() {
			continue
		}

		logger.TestHeader(v.Name())
		start := time.Now()
		result := lua.Execute(state.cache, folder+"/"+v.Name())
		logger.TestFooter(result, v.Name(), time.Since(start))
	}
}
