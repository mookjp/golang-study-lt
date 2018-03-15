package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // キャンバスの大きさ
	cells         = 100                 // 格子のマス目の数
	xyrange       = 30.0                // 軸の範囲（-xyrange..+xyrange）
	xyscale       = width / 2 / xyrange // x単位およびy単位あたりの画素数
	zscale        = height * 0.4        // z単位あたりの画素数
	angle         = math.Pi / 6         // x, y 軸の確度（-30度）
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30度), cos(30度)

var minZ float64
var maxZ float64

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az, aerr := corner(i+1, j)
			if aerr != nil {
				continue
			}
			bx, by, bz, berr := corner(i, j)
			if berr != nil {
				continue
			}
			cx, cy, cz, cerr := corner(i, j+1)
			if cerr != nil {
				continue
			}
			dx, dy, dz, derr := corner(i+1, j+1)
			if derr != nil {
				continue
			}
			var color string
			if az > 0.99 || bz > 0.99 || cz > 0.99 || dz > 0.99 {
				color = " style='fill:#ff0000'"
			}
			if az < -0.99 || bz < -0.99 || cz < -0.99 || dz < -0.99 {
				color = " style='fill:#0000ff'"
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'%s/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Println("</svg>")
	//fmt.Printf("\nminZ:%v, maxZ:%v", minZ, maxZ)
}

func corner(i, j int) (float64, float64, float64, error) {
	// マス目(i, j)の角の点(x, y)を見つける
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 面の高さzを計算する
	z := f(x, y)
	if math.IsNaN(z) {
		return 0, 0, 0, fmt.Errorf("invalid value was returned from f(x, y). x: %v, y: %v", x, y)
	}
	if minZ == 0 || z < minZ {
		minZ = z
	}
	if maxZ == 0 || z > maxZ {
		maxZ = z
	}

	// (x, y, z)を 2-D SVGキャンバス(sx, sy)へ等角的に投影
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, nil
}

// (x, y)面の高さを計算します
func f(x, y float64) float64 {
	r := math.Hypot(x, y) // (0, 0)からの距離
	return math.Sin(r)
}
