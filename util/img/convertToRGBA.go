package img

import (
	"fmt"
	"image"
	"reflect"
)

func ConvertToRGBA(pic image.Image) *image.RGBA {
	var data []uint8
	var rect image.Rectangle
	var stride int

	switch t := pic.(type) {
	case *image.NRGBA:
		data = t.Pix
		rect = t.Rect
		stride = t.Stride
	case *image.RGBA:
		return t
	default:
		fmt.Printf("Unknown image format: %v", reflect.TypeOf(pic))
	}

	return &image.RGBA{
		Pix:    data,
		Rect:   rect,
		Stride: stride,
	}
}
