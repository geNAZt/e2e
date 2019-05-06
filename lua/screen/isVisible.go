package screen

import (
	"github.com/kbinani/screenshot"
	lua "github.com/yuin/gopher-lua"
	"gomint.io/e2e/test/cache"
	"gomint.io/e2e/test/logger"
	"gomint.io/e2e/util/img"
	"time"
)

func IsVisible(cache *cache.Cache) lua.LGFunction {
	return func(L *lua.LState) int {
		start := time.Now()

		resource := L.ToString(1)
		logger.Debug("Trying to find '%s' on screen", resource)

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
				L.Push(lua.LTrue)

				logger.Benchmark(start)
				return 1
			}
		}

		logger.Debug("'%s' not found", resource)
		L.Push(lua.LFalse)

		logger.Benchmark(start)
		return 1
	}
}
