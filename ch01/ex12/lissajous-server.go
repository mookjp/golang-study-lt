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
	"fmt"
	"os"
	"strconv"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0xff, 0x00, 0x00, 0xff}}

var parameterErrorMessage = "Parameter Error: %v / %v=%v"

const (
	backGroundColorIndex = 0
	firstLineColorIndex  = 1
	secondLineColorIndex  = 2
)

type lissajousParams struct {
	cycles  int // number of complete x oscillator revolutions
	res     float64 // angular resolution
	size    int // image canvas covers [-size..+size]
	nframes int // number of animation frames
	delay   int // delay between frames in 10ms units
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// default
		params := lissajousParams{
			cycles: 5,
			res: 0.001,
			size: 100,
			nframes: 64,
			delay: 8,
		}
		for queryKey, queryVals := range r.URL.Query() {
			fmt.Fprintf(os.Stdout, "query: %s=%s\n", queryKey, queryVals)

			// TODO: commonize
			switch queryKey {
			case "cycles":
				converted, err := strconv.Atoi(queryVals[0])
				if err != nil {
					w.WriteHeader(400)
					fmt.Fprintf(w, parameterErrorMessage, err, queryKey, queryVals)
					return
				} else {
					params.cycles = converted
				}
			case "res":
				converted, err := strconv.ParseFloat(queryVals[0], 64)
				if err != nil {
					w.WriteHeader(400)
					fmt.Fprintf(w, parameterErrorMessage, err, queryKey, queryVals)
					return
				} else {
					params.res = converted
				}
			case "size":
				converted, err := strconv.Atoi(queryVals[0])
				if err != nil {
					w.WriteHeader(400)
					fmt.Fprintf(w, parameterErrorMessage, err, queryKey, queryVals)
					return
				} else {
					params.size = converted
				}
			case "nframes":
				converted, err := strconv.Atoi(queryVals[0])
				if err != nil {
					w.WriteHeader(400)
					fmt.Fprintf(w, parameterErrorMessage, err, queryKey, queryVals)
					return
				} else {
					params.nframes = converted
				}
			case "delay":
				converted, err := strconv.Atoi(queryVals[0])
				if err != nil {
					w.WriteHeader(400)
					fmt.Fprintf(w, parameterErrorMessage, err, queryKey, queryVals)
					return
				} else {
					params.delay = converted
				}
			}
		}
		lissajous(w, params)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// TODO: package
func lissajous(out io.Writer, params lissajousParams) {
	fmt.Fprintf(os.Stdout, "lissajous params: %v\n", params)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: params.nframes}
	phase := 0.0 // phase difference

	// アニメーションの1フレームを作成
	// 64フレーム分を外側のループで作成
	// 1フレームのimgのどこに色を置くかを、内側のループで作成
	for i := 0; i < params.nframes; i++ {
		// Rect is shorthand for Rectangle{Pt(x0, y0), Pt(x1, y1)}
		// Pt is shorthand for Point{X, Y}.
		// A Point is an X, Y coordinate pair. The axes increase right and down.
		rect := image.Rect(0, 0, 2*params.size+1, 2*params.size+1)
		// p.17
		// > すべての画素は最初にパレットのゼロ値（パレットの0番目の色）に設定され…（略）
		// TODO: どこのドキュメントに書いてある？
		img := image.NewPaletted(rect, palette)
		// imgの特定座標の色をセットする
		for t := 0.0; t < float64(params.cycles)*2.0*math.Pi; t += params.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			var colorIndex uint8
			if t > float64(params.cycles) {
				colorIndex = firstLineColorIndex
			} else {
				colorIndex = secondLineColorIndex
			}
			img.SetColorIndex(params.size + int(x*float64(params.size)+0.5), params.size + int(y*float64(params.size)+0.5),
				colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, params.delay)
		anim.Image = append(anim.Image, img)
	}
	encodeErr := gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
	if encodeErr != nil {
		fmt.Fprintf(os.Stdout, "encode error: %v", encodeErr)
	}
}
