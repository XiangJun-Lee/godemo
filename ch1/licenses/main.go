package main

import (
	"image/color"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	rand.Seed(time.Now().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out *os.File) {

}
