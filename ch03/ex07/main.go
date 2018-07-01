package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"math"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)

			img.Set(px, py, mandelbrot(z))
		}
	}

	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func newton(img image.RGBA, px int, py int, firstVal complex128) image.RGBA {
	x := firstVal
	var x2 complex128

	for x2 = x - (x * x - 2) /(x * 2); cmplx.Abs(x2 - x) > 0.0001; x = x2 {
		c := color.RGBA{}
		img.Set(px, py, c)
	}
	return x2
}
