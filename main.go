package main

import (
	"gomint.io/e2e/test"
	"gomint.io/e2e/util"
	"os"
)

func main() {
	// Clear debug
	os.RemoveAll("./debug/")
	os.Mkdir("./debug/", 0777)

	util.PlatformInit()

	// Create global test state
	state := test.NewState()
	state.Setup()

	// Teardown
	state.Teardown()
}
