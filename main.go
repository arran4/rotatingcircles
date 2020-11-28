package main

import (
	"github.com/fogleman/gg"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"log"
	"os"
)

var (
	c1 = color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	}
	c2 = color.RGBA{
		R: 128,
		G: 128,
		B: 0,
		A: 255,
	}
	palet = color.Palette{
		image.Black,
		c1,
		c2,
	}
)

func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)
	images := []*image.Paletted{}
	delays := []int{}
	for i := 0.0; i < 360; i += 10 {
		dc := gg.NewContext(650, 400)
		dc.Push()
		dc.Translate(100, 100)
		dc.RotateAbout(gg.Degrees(i), 100, 100)
		drawCircle(dc)
		dc.Pop()
		dc.Push()
		dc.Translate(350, 100)
		dc.RotateAbout(gg.Degrees(340-i), 100, 100)
		drawCircle(dc)
		dc.Pop()

		frame := dc.Image()

		pframe := image.NewPaletted(frame.Bounds(), palet)
		draw.Draw(pframe, pframe.Rect, frame, frame.Bounds().Min, draw.Over)

		images = append(images, pframe)
		delays = append(delays, 5)
	}
	f, err := os.Create("out.gif")
	if err != nil {
		log.Panicf("Error:%v", err)
	}
	defer f.Close()
	if err := gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
	}); err != nil {
		log.Panicf("Error: %v", err)
	}
}

func drawCircle(dc *gg.Context) {
	step := 30.0
	for i := 0.0; i < 360; i += step {
		dc.Push()
		dc.DrawArc(100, 100, 100, gg.Radians(i-step), gg.Radians(i))
		dc.DrawArc(100, 100, 80, gg.Radians(i), gg.Radians(i-step))
		switch int(i/step) % 2 {
		case 0:
			dc.SetRGB255(int(c1.R), int(c1.G), int(c1.B))
		default:
			dc.SetRGB255(int(c2.R), int(c2.G), int(c2.B))
		}
		dc.FillPreserve()
		dc.ClearPath()
		dc.Pop()
	}
}
