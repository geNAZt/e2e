package screen

import (
	"github.com/kbinani/screenshot"
	lua "github.com/yuin/gopher-lua"
	"gomint.io/e2e/test/cache"
	"gomint.io/e2e/test/logger"
	"gomint.io/e2e/util/img"
	"time"
)

func WaitUntilNotVisible(cache *cache.Cache) lua.LGFunction {
	return func(L *lua.LState) int {
		start := time.Now()

		logged := false

	outer:
		for {
			idx := 1

			for {
				resource := L.ToString(idx)
				idx++

				if !logged {
					logger.Debug("Waiting until '%s' is not on screen", resource)
				}

				loadedImage := cache.LoadOrGetImage(resource)

				n := screenshot.NumActiveDisplays()
				for i := 0; i < n; i++ {
					bounds := screenshot.GetDisplayBounds(i)

					screen, err := screenshot.CaptureRect(bounds)
					if err != nil {
						panic(err)
					}

					_, found := img.Contains(screen, loadedImage)
					if found {
						logged = true
						time.Sleep(16 * time.Millisecond)

						continue outer
					}
				}

				if L.Get(idx) == lua.LNil {
					logger.Benchmark(start)
					return 0
				}
			}
		}
	}
}
