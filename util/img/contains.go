package img

import (
	"bytes"
	"image"
)

/*func Contains(image, subimage *img.RGBA) (img.Point, bool) {
	debugImage := img.NewRGBA(image.Rect)

	toCheckWidth := image.Rect.Dx() - subimage.Rect.Dx()
	toCheckHeight := image.Rect.Dy() - subimage.Rect.Dy()

	// brute force N^2 check all places in the image
	for i := 0; i <= toCheckWidth; i++ {
	check_subimage:
		for j := 0; j <= toCheckHeight; j++ {
			for ii := 0; ii < subimage.Rect.Dx(); ii++ {
				for jj := 0; jj < subimage.Rect.Dx(); jj++ {
					if !isEqual(subimage.At(ii, jj).(color.RGBA), image.At(i + ii, j + jj).(color.RGBA), debugImage, i + ii, j + jj) {
						continue check_subimage
					}
				}
			}

			// if here, all pixels matched
			return img.Point{X:i, Y:j}, true
		}
	}

	file, err := os.Create(fmt.Sprintf("./debug/contains_%d.png", time.Now().UnixNano()))
	if err == nil {
		enc := &png.Encoder{CompressionLevel: png.NoCompression}
		enc.Encode(file, debugImage)
		file.Close()
	}

	return img.Point{}, false
}

func isEqual(a,b color.RGBA, debugImage *img.RGBA, x,y int) bool {
	if a.R == b.R && a.G == b.G && a.B == b.B {
		debugImage.SetRGBA(x,y, color.RGBA{0,0,0,255})
		return true
	}

	debugImage.SetRGBA(x,y, color.RGBA{uint8(math.Abs(float64(a.R - b.R))), uint8(math.Abs(float64(a.G - b.G))),uint8(math.Abs(float64(a.B - b.B))),255})
	return math.Abs(float64(a.R - b.R)) < 2 && math.Abs(float64(a.G - b.G)) < 2 && math.Abs(float64(a.B - b.B)) < 2
}*/

func Contains(m0, m1 *image.RGBA) (image.Point, bool) {
	s1 := m1.Rect.Size()
	s0 := m0.Rect.Size().Sub(s1)

	for y0 := 0; y0 < s0.Y; y0++ {
	loopx0:
		for x0 := 0; x0 < s0.X; x0++ {
			i0 := y0*m0.Stride + x0*4
			j0 := i0 + s1.X*4
			i1 := 0
			j1 := s1.X * 4

			for y1 := 0; y1 < s1.Y; y1++ {
				if !bytes.Equal(m0.Pix[i0:j0], m1.Pix[i1:j1]) {
					continue loopx0
				}

				if i1 == 0 {
					addM1 := (s1.Y - 1) * m1.Stride
					addM0 := (s1.Y - 1) * m0.Stride

					if j0+addM0 < len(m0.Pix) {
						if !bytes.Equal(m0.Pix[i0+addM0:j0+addM0], m1.Pix[i1+addM1:j1+addM1]) {
							continue loopx0
						}
					}
				}

				i0 += m0.Stride
				j0 += m0.Stride
				i1 += m1.Stride
				j1 += m1.Stride
			}

			return m0.Rect.Min.Add(image.Pt(x0, y0)), true
		}
	}

	return image.Point{}, false
}
