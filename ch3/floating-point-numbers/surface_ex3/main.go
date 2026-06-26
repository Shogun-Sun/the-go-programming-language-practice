package main

import (
	"fmt"
	"math"
)

const (
	width, height = 1920, 1080          // Размер канвы в пикселях
	cells         = 100                 // Количество ячеек сетки
	xyrange       = 30.0                // Диапазон осей
	xyscale       = width / 2 / xyrange // Пикселей в единице x или y
	zscale        = height * 0.4        // Пикселей в единице z
	angle         = math.Pi / 6         // Углы осей x, y (=30°)
)

const (
	minZ = -0.2172
	maxZ = 1.0
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: black; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)

			zAvg := (az + bz + cz + dz) / 4.0

			colorStr := getColor(zAvg)

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, colorStr)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	// Поиск угловой точки (x,y) ячейки (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Вычисление высоты поверхности z
	z := f(x, y)

	// Изометрически проецируем (x,y,z) на двумерную канву SVG (sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // Расстояние от (0,0)
	if r == 0 {
		return 1.0
	}
	return math.Sin(r) / r
}

func getColor(z float64) string {
	if z < minZ {
		z = minZ
	}
	if z > maxZ {
		z = maxZ
	}

	t := (z - minZ) / (maxZ - minZ)

	r := int(t * 255)
	b := int((1 - t) * 255)
	g := 0

	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}
