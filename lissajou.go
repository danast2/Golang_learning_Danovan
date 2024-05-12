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

var palette = []color.Color{color.RGBA{0, 255, 0, 255}, color.Black}

const (
	greenIndex = 0 // первый цвет палитры
	blackIndex = 1 // next цвет палитры
)

func main() {
	lissajou(os.Stdout)
}

func lissajou(out io.Writer) {
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

		// Устанавливаем черный цвет фона
		for y := 0; y < 2*size+1; y++ {
			for x := 0; x < 2*size+1; x++ {
				img.SetColorIndex(x, y, blackIndex)
			}
		}

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) //игнори ошибки
}

//чтобы получить красивую гифку, а не бинарный вывод в консоль, достаточно сохранить
//вывод программы в GIF go run lissajou.go > lissajou.gif
// затем гифку можно будет просто открыть
