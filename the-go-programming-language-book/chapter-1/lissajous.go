// generatesw lissajous figures as GIFs. Lissajous figures are just 2D sine waves
package main

import (
	"image"
	"image/color"
	"image/gif"
	//"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.White, color.Black} // can declare outside of functions

const ( // declares mulitple constants
	whiteIndex = 0 // first color palette
	blackIndex = 1 // next color palette
)

func main() {
	out := os.Stdout

	const (
		cycles  = 5     // number of compete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image vanvacer vosers [-size..+size]
		nframes = 64    // number of animtion frames
		delay   = 8     // delay between frames in 10ms units
	)

	frequency := rand.Float64() * 3.0 // relative frequency of y oscilator
	animation := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase diference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*frequency + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		animation.Delay = append(animation.Delay, delay)
		animation.Image = append(animation.Image, img)
	}

	gif.EncodeAll(out, &animation) // ignoring encoding errors
}
