package cache

import (
	"gomint.io/e2e/util/img"
	"image"
	"image/png"
	"os"
)

type Cache struct {
	lastSeenDisplay int
	pictureCache    map[string]*image.RGBA
	folder          string
}

func NewCache() *Cache {
	return &Cache{
		lastSeenDisplay: 0,
		pictureCache:    make(map[string]*image.RGBA),
		folder:          "",
	}
}

func (state *Cache) SetFolder(folder string) {
	state.folder = folder
}

func (state *Cache) SetLastSeenDisplay(display int) {
	state.lastSeenDisplay = display
}

func (state *Cache) GetLastSeenDisplay() int {
	return state.lastSeenDisplay
}

func (state *Cache) Reset() {
	state.pictureCache = make(map[string]*image.RGBA)
	state.lastSeenDisplay = 0
}

func (state *Cache) LoadOrGetImage(resource string) *image.RGBA {
	// Check if cache has this image
	if pic, ok := state.pictureCache[resource]; ok {
		return pic
	}

	// Load image and store if possible
	pngFile, err := os.Open(state.folder + "/resources/" + resource)
	if err != nil {
		// TODO: Add test fail handling
		panic(err)
	}

	defer pngFile.Close()

	loadedImage, err := png.Decode(pngFile)
	if err != nil {
		// TODO: Add test fail handling
		panic(err)
	}

	state.pictureCache[resource] = img.ConvertToRGBA(loadedImage)
	return state.pictureCache[resource]
}
