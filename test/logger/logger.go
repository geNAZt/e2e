package logger

import (
	"fmt"
	"os"
	"time"
)

func TestFolderHeader(amountOfTests int, suiteName string) {
	fmt.Printf("⚡ Executing '%d' tests in suite '%s'\n\n", amountOfTests, suiteName)
}

func TestHeader(testName string) {
	fmt.Printf("    ⏰ Starting test '%s' ...\n", testName)
}

func Info(format string) {
	fmt.Printf("        🚩 " + format + "\n")
}

func Error(err error) {
	fmt.Printf("Caught error: %v", err)
}

func Failure(err error, format string, args ...interface{}) {
	fmt.Printf("        ✂ "+format+"\n", args...)

	panic(err) // let the lua VM catch this
}

func Benchmark(start time.Time) {
	fmt.Printf("            🚀 Execution took '%s'\n", time.Since(start))
}

func TestFooter(success bool, testName string, duration time.Duration) {
	symbol := "✅"
	status := "succeeded"

	if !success {
		symbol = "❌"
		status = "failed"
	}

	fmt.Printf("    %s test '%s' %s after %s\n\n", symbol, testName, status, duration)
}

func Fatal(format string, args ...interface{}) {
	fmt.Printf("❌❗ "+format+"\n", args...)
	os.Exit(1)
}

func Debug(format string, args ...interface{}) {
	fmt.Printf("        🚨 "+format+"\n", args...)
}
