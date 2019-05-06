package screen

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/kbinani/screenshot"
	lua "github.com/yuin/gopher-lua"
	"gomint.io/e2e/test/logger"
	"image/png"
	"os"
	"os/exec"
	"time"
)

func ReadText(L *lua.LState) int {
	x := L.ToInt(1)
	y := L.ToInt(2)

	mouseX, mouseY := robotgo.GetMousePos()

	// Capture given area
	screen, err := screenshot.Capture(mouseX, mouseY, x, y)
	if err != nil {
		logger.Failure(err, "Could not get screen for OCR")
	}

	// We need a tmp picture for tesseract
	file, err := os.Create(fmt.Sprintf("./debug/mcpee2e_ocr_%d_%d_%d.png", time.Now().Unix(), mouseX, mouseY))
	if err != nil {
		logger.Failure(err, "Could not store picture for OCR")
	}

	// Store the picture
	png.Encode(file, screen)
	file.Close()
	// defer os.Remove(file.Name())

	// OCR with tesseract
	body, err := exec.Command("tesseract", file.Name(), "stdout", "-l", "eng").Output()
	if err != nil {
		logger.Failure(err, "Tesseract failed")
	}

	logger.Debug("Got text from OCR '%s'", string(body))
	L.Push(lua.LString(string(body)))
	return 1
}
