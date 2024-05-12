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

//var palette = []color.Color{color.RGBA{0, 255, 0, 255}, color.Black}

var palette_new = []color.Color{
	color.RGBA{0, 0, 0, 255},   // Черный цвет
	color.RGBA{255, 0, 0, 255}, // Красный цвет
	color.RGBA{0, 255, 0, 255}, // Зелёный цвет
	color.RGBA{0, 0, 255, 255}, // Синий цвет
}

const (
	greenIndex_new = 0 // первый цвет палитры
	blackIndex_new = 1 // next цвет палитры
)

func main() {
	lissajou_new(os.Stdout)
}

func lissajou_new(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 120 //можно поиграться с количеством кадров в гифке
		delay   = 8
	)
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {

		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette_new)

		// Устанавливаем черный цвет фона
		for y := 0; y < 2*size+1; y++ {
			for x := 0; x < 2*size+1; x++ {
				img.SetColorIndex(x, y, blackIndex_new)
			}
		}

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// Выбор цвета в зависимости от координаты x
			colorIndex := uint8((math.Sin(x) + 1) * 3 / 2)

			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) //игнори ошибки
}
