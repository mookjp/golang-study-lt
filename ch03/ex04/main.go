package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
)

var parameterErrorMessage = "Parameter Error: %v / %v=%v"

const (
	defaultWidth, defaultHeight = 600, 320    // キャンバスの大きさ
	cells                       = 100         // 格子のマス目の数
	xyrange                     = 30.0        // 軸の範囲（-xyrange..+xyrange）
	angle                       = math.Pi / 6 // x, y 軸の確度（-30度）
	defaultColor                = "fefefe"
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30度), cos(30度)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")

		width := defaultWidth
		height := defaultHeight
		color := defaultColor
		for queryKey, queryVals := range r.URL.Query() {
			fmt.Fprintf(os.Stdout, "query: %s=%s\n", queryKey, queryVals)

			switch queryKey {
			case "width":
				converted, err := strconv.Atoi(queryVals[0])
				if err != nil {
					handleError(w, err, queryKey, queryVals)
					return
				} else {
					width = converted
				}
			case "height":
				converted, err := strconv.Atoi(queryVals[0])
				if err != nil {
					handleError(w, err, queryKey, queryVals)
					return
				} else {
					height = converted
				}
			case "color":
				color = queryVals[0]
			}
		}
		surface(w, float64(width), float64(height), fmt.Sprintf("#%s", color))
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func surface(w http.ResponseWriter, width, height float64, color string) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: %s; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", color, width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, aerr := corner(width, height, i+1, j)
			if aerr != nil {
				continue
			}
			bx, by, berr := corner(width, height, i, j)
			if berr != nil {
				continue
			}
			cx, cy, cerr := corner(width, height, i, j+1)
			if cerr != nil {
				continue
			}
			dx, dy, derr := corner(width, height, i+1, j+1)
			if derr != nil {
				continue
			}
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprint(w, "</svg>")
}

func corner(width, height float64, i, j int) (float64, float64, error) {
	xyscale := width / 2 / xyrange // x単位およびy単位あたりの画素数
	zscale := height * 0.4         // z単位あたりの画素数

	// マス目(i, j)の角の点(x, y)を見つける
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 面の高さzを計算する
	z := f(x, y)
	if math.IsNaN(z) {
		return 0, 0, fmt.Errorf("invalid value was returned from f(x, y). x: %v, y: %v", x, y)
	}

	// (x, y, z)を 2-D SVGキャンバス(sx, sy)へ等角的に投影
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, nil
}

// (x, y)面の高さを計算します
func f(x, y float64) float64 {
	r := math.Hypot(x, y) // (0, 0)からの距離
	return math.Sin(r)
}

func handleError(w http.ResponseWriter, err error, queryKey string, queryVals []string) {
	w.WriteHeader(400)
	fmt.Fprintf(w, parameterErrorMessage, err, queryKey, queryVals)
	return
}
