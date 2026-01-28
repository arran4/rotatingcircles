package main

import "testing"

func TestMeasureUnidirectionalArrow(t *testing.T) {
	w, h := measureUnidirectionalArrow()
	if w <= 0 || h <= 0 {
		t.Errorf("expected positive dimensions, got w=%f, h=%f", w, h)
	}
}
