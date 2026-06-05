package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{
	color.RGBA{0x05, 0x11, 0x08, 0xff}, // Зелено-черный фон
	color.RGBA{0x00, 0x4d, 0x20, 0xff}, // Темно-зеленый
	color.RGBA{0x00, 0x73, 0x33, 0xff}, // Кувшинковый зеленый
	color.RGBA{0x00, 0x99, 0x44, 0xff}, // Изумрудный
	color.RGBA{0x2e, 0xcc, 0x71, 0xff}, // Неоновый зеленый очень сочный
	color.RGBA{0x39, 0xff, 0x14, 0xff}, // Кислотно-зеленый светящийся
	color.RGBA{0x7c, 0xfc, 0x00, 0xff}, // Зеленая лужайка
	color.RGBA{0xa3, 0xff, 0xb4, 0xff}, // Мятно-белый
}

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}

	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			colorShift := int(t*1.5 + phase*2.0)
			colorIndex := uint8(colorShift%(len(palette)-1) + 1)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
