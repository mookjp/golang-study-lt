package main

import (
	"net/http"
	"log"
	"io"
	"image/gif"
	"image"
	"math"
	"image/color"
	"math/rand"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0xff, 0x00, 0x00, 0xff}}

const (
	backGroundColorIndex = 0
	firstLineColorIndex  = 1
	secondLineColorIndex  = 2
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// TODO: package
func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	// アニメーションの1フレームを作成
	// 64フレーム分を外側のループで作成
	// 1フレームのimgのどこに色を置くかを、内側のループで作成
	for i := 0; i < nframes; i++ {
		// Rect is shorthand for Rectangle{Pt(x0, y0), Pt(x1, y1)}
		// Pt is shorthand for Point{X, Y}.
		// A Point is an X, Y coordinate pair. The axes increase right and down.
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		// p.17
		// > すべての画素は最初にパレットのゼロ値（パレットの0番目の色）に設定され…（略）
		// TODO: どこのドキュメントに書いてある？
		img := image.NewPaletted(rect, palette)
		// imgの特定座標の色をセットする
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			var colorIndex uint8
			if t > cycles {
				colorIndex = firstLineColorIndex
			} else {
				colorIndex = secondLineColorIndex
			}
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
