package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		origWidth, origHeight  = 1024, 1024
		width, height          = origWidth / 2, origHeight / 2
	)

	// complexの2次元配列を先に作成する
	var colorMatrix [origWidth][origHeight]color.RGBA
	for py := 0; py < origHeight; py++ {
		y := float64(py)/origHeight*(ymax-ymin) + ymin
		for px := 0; px < origWidth; px++ {
			x := float64(px)/origWidth*(xmax-xmin) + xmin
			z := complex(x, y)

			colorMatrix[px][py] = mandelbrot(z)
			//img.Set(px/2, py/2, mandelbrot(z))
		}
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	// TODO loop
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			subX := x * 2
			subY := y * 2

			newR := colorMatrix[subX][subY].R +
				colorMatrix[subX][subY+1].R +
				colorMatrix[subX+1][subY].R +
				colorMatrix[subX+1][subY+1].R/4
			newG := colorMatrix[subX][subY].R +
				colorMatrix[subX][subY+1].R +
				colorMatrix[subX+1][subY].R +
				colorMatrix[subX+1][subY+1].R/4
			newB := colorMatrix[subX][subY].R +
				colorMatrix[subX][subY+1].R +
				colorMatrix[subX+1][subY].R +
				colorMatrix[subX+1][subY+1].R/4
			newA := colorMatrix[subX][subY].R +
				colorMatrix[subX][subY+1].R +
				colorMatrix[subX+1][subY].R +
				colorMatrix[subX+1][subY+1].R/4

			img.Set(x, y, color.RGBA{newR, newG, newB, newA})
		}
	}

	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.RGBA {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			r, g, b, a := color.Gray{255 - contrast*n}.RGBA()
			return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
		}
	}
	r, g, b, a := color.Black.RGBA()
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}
