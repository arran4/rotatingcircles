package main

import (
	"github.com/fogleman/gg"
	"math"
	"testing"
)

func BenchmarkCurrent(b *testing.B) {
	dc := gg.NewContext(650, 400)
	const distance = 35
	for i := 0; i < b.N; i++ {
		arrowWidth, arrowHeight := measureUnidirectionalArrow()
		dc.Push()
		dc.Translate(100-distance+100-arrowWidth/2, 100+100-arrowHeight/2-10*(math.Round((90)/90)))
		drawUnidirectionalArrow(dc)
		dc.Pop()
		dc.Push()
		dc.Translate(350+distance+100-arrowWidth/2, 100+100-arrowHeight/2+10*(math.Round((90)/90)))
		dc.RotateAbout(gg.Radians(180), arrowWidth/2, arrowHeight/2)
		drawUnidirectionalArrow(dc)
		dc.Pop()
	}
}

func BenchmarkOptimized(b *testing.B) {
	dc := gg.NewContext(650, 400)
	const distance = 35
	for i := 0; i < b.N; i++ {
		arrowWidth, arrowHeight := measureUnidirectionalArrow()
		dc.Push()
		dc.Translate(100-distance+100-arrowWidth/2, 100+100-arrowHeight/2-10*1.0)
		drawUnidirectionalArrow(dc)
		dc.Pop()
		dc.Push()
		dc.Translate(350+distance+100-arrowWidth/2, 100+100-arrowHeight/2+10*1.0)
		dc.RotateAbout(gg.Radians(180), arrowWidth/2, arrowHeight/2)
		drawUnidirectionalArrow(dc)
		dc.Pop()
	}
}
