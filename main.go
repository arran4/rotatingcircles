package main

import (
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"log"
	"math"
	"os"
)

var (
	c1     = color.RGBA{0x44, 0xff, 0x44, 0xff}
	c2     = color.RGBA{0xff, 0x44, 0x44, 0xff}
	palet  = color.Palette([]color.Color{color.Black, color.White, c1, c2})
	circle = MakeCircle()
)

const RotationAngle = 5

func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)
	images := []*image.Paletted{}
	delays := []int{}

	for i := 0; i < 360/RotationAngle; i++ {
		log.Printf("Frame %v", i)
		frame := drawFrame(i)
		pframe := image.NewPaletted(frame.Rect, palet)
		draw.Draw(pframe, pframe.Rect, frame, frame.Rect.Min, draw.Over)
		images = append(images, pframe)
		delays = append(delays, 25)
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

func drawFrame(frame int) *image.RGBA {
	dest := image.NewRGBA(image.Rect(0, 0, 1000, 800.0))
	gc := draw2dimg.NewGraphicContext(dest)
	gc.Translate(200, 200)
	gc.DrawImage(circle)
	gc.Save()
	gc.Translate(200, 200)
	gc.Rotate(float64(frame) * RotationAngle * math.Pi / 180)
	gc.MoveTo(100, 100)
	//gc.DrawImage(circle)
	//gc.Translate(200, 200)
	gc.DrawImage(circle)
	//gc.Translate(-100, -100)
	//gc.DrawImage(circle)
	return dest
}

func MakeCircle() *image.RGBA {
	dest := image.NewRGBA(image.Rect(0, 0, 200, 200.0))
	gc := draw2dimg.NewGraphicContext(dest)

	drawArc(gc, 100, 100, 0, c1)
	drawArc(gc, 0, 0, 90*(math.Pi/180), c2)
	drawArc(gc, 0, 0, 180*(math.Pi/180), c2)
	drawArc(gc, 0, 0, 270*(math.Pi/180), c1)

	return dest
}

func drawArc(gc *draw2dimg.GraphicContext, x, y, r float64, c color.Color) {
	gc.Translate(x, y)
	gc.Rotate(r)
	gc.SetFillColor(c)
	gc.SetLineWidth(0)
	gc.MoveTo(0, 75)
	gc.LineTo(0, 100)
	gc.QuadCurveTo(100, 100, 100, 0)

	gc.LineTo(75, 0)
	gc.QuadCurveTo(75, 75, 0, 75)
	gc.Close()
	gc.FillStroke()
}
