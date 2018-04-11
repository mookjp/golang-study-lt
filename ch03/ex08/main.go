package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math"
	"math/big"
	"math/cmplx"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
)

func Run(mode string, width int, height int, out io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)

			switch mode {
			case "complex64":
				img.Set(px, py, byComplex64(complex64(z)))
			case "complex128":
				img.Set(px, py, byComplex128(z))
			case "float":
				img.Set(px, py, byFloat(real(z)))
			case "rat":
				img.Set(px, py, byRat(int64(real(z))))
			}
		}
	}

	png.Encode(out, img)
}

func byComplex64(z complex64) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if math.Hypot(float64(real(v)), float64(imag(v))) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func byComplex128(z complex128) color.Color {
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

func byFloat(z float64) color.Color {
	const iterations = 200
	const contrast = 15

	v := new(big.Float)
	for n := uint8(0); n < iterations; n++ {
		//v = v*v + z
		v = v.Mul(v, v)
		v.Add(v, new(big.Float).SetFloat64(z))
		res := v.Abs(v).Cmp(new(big.Float).SetFloat64(2))
		if res > 0 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func byRat(z int64) color.Color {
	const iterations = 200
	const contrast = 15

	v := new(big.Rat)
	for n := uint8(0); n < iterations; n++ {
		//v = v*v + z
		v = v.Mul(v, v)
		v.Add(v, new(big.Rat).SetInt64(z))
		res := v.Abs(v).Cmp(new(big.Rat).SetInt64(2))
		if res > 0 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
