package main

import (
	"image/color"
	"image/gif"
	"io"
	"math/rand"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // первый цвет палитры
	blackIndex = 1 // next цвет палитры
)

func main() {

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
}
