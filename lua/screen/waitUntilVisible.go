package screen

import (
	"github.com/kbinani/screenshot"
	lua "github.com/yuin/gopher-lua"
	"gomint.io/e2e/test/cache"
	"gomint.io/e2e/test/logger"
	"gomint.io/e2e/util/img"
	"time"
)

func WaitUntilVisible(cache *cache.Cache) lua.LGFunction {
	return func(L *lua.LState) int {
		start := time.Now()

		logged := false

		for {
			idx := 1

			for {
				resource := L.ToString(idx)
				idx++

				if !logged {
					logger.Debug("Waiting until '%s' is on screen", resource)
				}

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
						logger.Benchmark(start)
						L.Push(lua.LNumber(point.X))
						L.Push(lua.LNumber(point.Y))
						return 2
					} else {
						logger.Debug("Image not found '%s'", resource)
						Debug(L)
					}
				}

				if L.Get(idx) == lua.LNil {
					break
				}
			}

			logged = true
			time.Sleep(500 * time.Millisecond)
		}
	}
}
