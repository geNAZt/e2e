package screen

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/kbinani/screenshot"
	lua "github.com/yuin/gopher-lua"
	"gomint.io/e2e/test/logger"
	"image/color"
	"image/png"
	"os"
	"time"
)

func Debug(L *lua.LState) int {
	logger.Debug("Taking debug screenshots")

	n := screenshot.NumActiveDisplays()
	for i := 0; i < n; i++ {
		logger.Debug("Found display for debug screenshot '%d'", i)
		bounds := screenshot.GetDisplayBounds(i)

		screen, err := screenshot.CaptureRect(bounds)
		if err != nil {
			logger.Failure(err, "Could not capture screen for bounds '%s'", bounds)
		}

		// show mouse
		x, y := robotgo.GetMousePos()
		screen.SetRGBA(x-1, y, color.RGBA{
			R: 255,
			G: 0,
			B: 0,
			A: 255,
		})
		screen.SetRGBA(x+1, y, color.RGBA{
			R: 255,
			G: 0,
			B: 0,
			A: 255,
		})
		screen.SetRGBA(x, y-1, color.RGBA{
			R: 255,
			G: 0,
			B: 0,
			A: 255,
		})
		screen.SetRGBA(x, y+1, color.RGBA{
			R: 255,
			G: 0,
			B: 0,
			A: 255,
		})
		screen.SetRGBA(x, y, color.RGBA{
			R: 255,
			G: 0,
			B: 0,
			A: 255,
		})

		logger.Debug("Mouse position: '%d', '%d'", x, y)

		debugFile, err := os.Create(fmt.Sprintf("./debug/%d_%d.png", time.Now().Unix(), i))
		if err != nil {
			logger.Failure(err, "Could not store debug screen")
		}

		enc := &png.Encoder{CompressionLevel: png.NoCompression}
		enc.Encode(debugFile, screen)
		debugFile.Close()
	}

	return 0
}
