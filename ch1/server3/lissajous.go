package main

import (
	"image"
	"image/color"
	"image/gif"
	"math"
	"math/rand"
	"net/http"
)

var palette = []color.Color{color.White, color.Black, color.RGBA{R: 0xff, A: 0xff}, color.RGBA{G: 0xff, A: 0xff}}

const (
	whiteIndex = 0
	blackIndex = 1
)

func lissajous(w http.ResponseWriter, cycles float64) {
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	//fmt.Println(anim.Config)
	gif.EncodeAll(w, &anim)
}
