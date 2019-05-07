package lua

import (
	"github.com/yuin/gopher-lua"
	"gomint.io/e2e/lua/keyboard"
	"gomint.io/e2e/lua/log"
	"gomint.io/e2e/lua/mouse"
	"gomint.io/e2e/lua/process"
	"gomint.io/e2e/lua/screen"
	"gomint.io/e2e/lua/util"
	"gomint.io/e2e/lua/http"
	"gomint.io/e2e/lua/window"
	"gomint.io/e2e/test/cache"
	"gomint.io/e2e/test/logger"
	"github.com/kohkimakimoto/gluafs"
)

func Execute(cache *cache.Cache, file string) bool {
	// Create new lua VM
	L := lua.NewState()
	defer L.Close()

	L.PreloadModule("fs", gluafs.Loader)

	// Register log helper
	functions := make(map[string]lua.LGFunction)
	functions["info"] = log.Info
	L.RegisterModule("log", functions)

	// Register mouse helper
	functions = make(map[string]lua.LGFunction)
	functions["move"] = mouse.Move
	functions["moveRelative"] = mouse.MoveRelative
	functions["click"] = mouse.Click
	L.RegisterModule("mouse", functions)

	// Register screen helper
	functions = make(map[string]lua.LGFunction)
	functions["isVisible"] = screen.IsVisible(cache)
	functions["clickOn"] = screen.ClickOn(cache)
	functions["waitUntilVisible"] = screen.WaitUntilVisible(cache)
	functions["waitUntilNotVisible"] = screen.WaitUntilNotVisible(cache)
	functions["debug"] = screen.Debug
	functions["readText"] = screen.ReadText
	L.RegisterModule("screen", functions)

	// Register process helper
	functions = make(map[string]lua.LGFunction)
	functions["waitFor"] = process.WaitForProcess
	functions["kill"] = process.Kill
	functions["exec"] = process.Exec
	functions["run"] = process.Run
	L.RegisterModule("process", functions)

	// Register util helper
	functions = make(map[string]lua.LGFunction)
	functions["sleep"] = util.Sleep
	functions["sleep"] = util.Sleep
	L.RegisterModule("util", functions)

	// Register keyboard helper
	functions = make(map[string]lua.LGFunction)
	functions["type"] = keyboard.Type
	functions["press"] = keyboard.Press
	functions["toggle"] = keyboard.Toggle
	L.RegisterModule("keyboard", functions)

	// Register http helper
	functions = make(map[string]lua.LGFunction)
	functions["download"] = http.Download
	L.RegisterModule("http", functions)

	// Register window helper
	functions = make(map[string]lua.LGFunction)
	functions["getWindows"] = window.GetWindows
	functions["getBounds"] = window.GetBounds
	L.RegisterModule("window", functions)

	// Execute the script
	if err := L.DoFile(file); err != nil {
		logger.Error(err)
		return false
	}

	return true
}
