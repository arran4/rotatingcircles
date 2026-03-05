package main

import (
	"github.com/fogleman/gg"
	"testing"
)

func BenchmarkDrawUnidirectionalArrow(b *testing.B) {
	dc := gg.NewContext(1000, 1000)
	for i := 0; i < b.N; i++ {
		drawUnidirectionalArrow(dc)
		dc.ClearPath()
	}
}
