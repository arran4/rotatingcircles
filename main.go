package main

import (
	"github.com/fogleman/gg"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"log"
	"math"
	"os"
)

var (
	c1 = color.RGBA{
		R: 0,
		G: 0,
		B: 255,
		A: 255,
	}
	c2 = color.RGBA{
		R: 222,
		G: 200,
		B: 55,
		A: 255,
	}
	palet = color.Palette{
		color.Gray{Y: 256 - 64 - 32},
		image.Black,
		image.White,
		c1,
		c2,
	}
)

func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)
	images := []*image.Paletted{}
	delays := []int{}
	const step = 30
	for i := 0.0; i < 360; i += step {
		dc := gg.NewContext(650, 400)
		dc.Push()
		dc.Translate(100-25, 100)
		dc.RotateAbout(gg.Radians(i), 100, 100)
		drawCircle(dc)
		dc.Pop()
		dc.Push()
		dc.Translate(350+25, 100)
		dc.RotateAbout(gg.Radians(340-i), 100, 100)
		drawCircle(dc)
		dc.Pop()
		arrowWidth, arrowHeight := measureUnidirectionalArrow()
		dc.Push()
		dc.Translate(100-25+100-arrowWidth/2, 100+100-arrowHeight/2-10*(math.Round((90)/90)))
		drawUnidirectionalArrow(dc)
		dc.Pop()
		dc.Push()
		dc.Translate(350+25+100-arrowWidth/2, 100+100-arrowHeight/2+10*(math.Round((90)/90)))
		dc.RotateAbout(gg.Radians(180), arrowWidth/2, arrowHeight/2)
		drawUnidirectionalArrow(dc)
		dc.Pop()

		frame := dc.Image()

		pframe := image.NewPaletted(frame.Bounds(), palet)
		draw.Draw(pframe, pframe.Rect, frame, frame.Bounds().Min, draw.Over)

		images = append(images, pframe)
		delays = append(delays, step/3)
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
	log.Printf("Done")
}

func drawUnidirectionalArrow(dc *gg.Context) {
	dc.Push()
	/*
		. . # . .
		. . . . .
		# # . # #
		. . . . .
		. . . . .
		. . . . .
		. # . # .
	*/
	const n = 10
	points := [][]float64{
		{2 * n, 0 * n},
		{4 * n, (2 + 1) * n},
		{2.5 * n, (2 + 1) * n},
		{2.5 * n, (6 + 1) * n},
		{1.5 * n, (6 + 1) * n},
		{1.5 * n, (2 + 1) * n},
		{0 * n, (2 + 1) * n},
	}
	dc.NewSubPath()
	for i := 0; i < len(points); i += 1 {
		dc.LineTo(points[i%len(points)][0], points[i%len(points)][1])
	}
	dc.ClosePath()
	dc.SetColor(image.Black)
	dc.SetLineWidth(2)
	dc.StrokePreserve()
	dc.FillPreserve()
	dc.Pop()
}

func measureUnidirectionalArrow() (w, h float64) {
	const n = 10
	return 4 * n, 7 * n
}

func drawCircle(dc *gg.Context) {
	step := 90.0
	radius := 80.0
	thickness := 35.0
	for i := 0.0; i < 360; i += step {
		dc.Push()
		dc.DrawArc(100, 100, radius+thickness, gg.Radians(i-step), gg.Radians(i))
		dc.DrawArc(100, 100, radius, gg.Radians(i), gg.Radians(i-step))
		switch int(i/step) % 2 {
		case 0:
			dc.SetColor(c1)
		default:
			dc.SetColor(c2)
		}
		dc.FillPreserve()
		dc.ClearPath()
		dc.Pop()
	}
}
