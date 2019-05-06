package main

import (
	"gomint.io/e2e/test"
	"os"
)

func main() {
	// Clear debug
	os.RemoveAll("./debug/")
	os.Mkdir("./debug/", 0777)

	// Create global test state
	state := test.NewState()
	state.Setup()

	// Teardown
	state.Teardown()
}
