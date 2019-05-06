package screen

import (
	"github.com/go-vgo/robotgo"
	"github.com/kbinani/screenshot"
	lua "github.com/yuin/gopher-lua"
	"gomint.io/e2e/test/cache"
	"gomint.io/e2e/test/logger"
	"gomint.io/e2e/util/img"
	"time"
)

func ClickOn(cache *cache.Cache) lua.LGFunction {
	return func(L *lua.LState) int {
		start := time.Now()

		idx := 1

		for {
			resource := L.ToString(idx)
			idx++

			logger.Debug("Clicking on image '%s'", resource)

			loadedImage := cache.LoadOrGetImage(resource)

			n := screenshot.NumActiveDisplays()
			for i := 0; i < n; i++ {
				bounds := screenshot.GetDisplayBounds(i)

				screen, err := screenshot.CaptureRect(bounds)
				if err != nil {
					panic(err)
				}

				point, found := img.Contains(screen, loadedImage)
				if found {
					logger.Debug("'%s' found at '%d', '%d'", resource, point.X, point.Y)

					robotgo.MoveMouse(point.X, point.Y)
					time.Sleep(50 * time.Millisecond)
					robotgo.MouseClick()

					L.Push(lua.LTrue)

					logger.Benchmark(start)
					return 1
				}
			}

			if L.Get(idx) == lua.LNil {
				break
			}
		}

		logger.Debug("Images not found")
		L.Push(lua.LFalse)

		logger.Benchmark(start)
		return 1
	}
}
