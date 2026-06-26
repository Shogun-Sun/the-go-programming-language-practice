package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	cells   = 100         // Количество ячеек сетки
	xyrange = 30.0        // Диапазон осей
	angle   = math.Pi / 6 // Углы осей x, y (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	http.HandleFunc("/", plotHandler)
	log.Println("Сервер успешно запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func plotHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	width := 600
	height := 320
	strokeColor := "grey"

	query := r.URL.Query()

	if wStr := query.Get("width"); wStr != "" {
		if val, err := strconv.Atoi(wStr); err == nil && val > 0 {
			width = val
		}
	}
	if hStr := query.Get("height"); hStr != "" {
		if val, err := strconv.Atoi(hStr); err == nil && val > 0 {
			height = val
		}
	}
	if colStr := query.Get("color"); colStr != "" {
		strokeColor = colStr
	}

	xyscale := float64(width) / 2 / xyrange
	zscale := float64(height) * 0.4

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: %s; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", strokeColor, width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, okA := corner(i+1, j, width, height, xyscale, zscale)
			bx, by, okB := corner(i, j, width, height, xyscale, zscale)
			cx, cy, okC := corner(i, j+1, width, height, xyscale, zscale)
			dx, dy, okD := corner(i+1, j+1, width, height, xyscale, zscale)

			if !okA || !okB || !okC || !okD {
				continue
			}

			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(i, j, width, height int, xyscale, zscale float64) (float64, float64, bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)
	if math.IsNaN(z) || math.IsInf(z, 0) {
		return 0, 0, false
	}

	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	if r == 0 {
		return 1.0
	}
	return math.Sin(r) / r
}
