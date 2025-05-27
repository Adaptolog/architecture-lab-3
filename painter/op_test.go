package painter

import (
	"image"
	"image/color"
	"testing"
)

func TestWhiteFill_Do(t *testing.T) {
	state := &State{}
	op := WhiteFill{}

	op.Do(state)
	if state.BgColor != color.White {
		t.Error("WhiteFill didn't set background to white")
	}
}

func TestGreenFill_Do(t *testing.T) {
	state := &State{}
	op := GreenFill{}

	op.Do(state)
	expected := color.RGBA{G: 0xff, A: 0xff}
	if state.BgColor != expected {
		t.Error("GreenFill didn't set background to green")
	}
}

func TestBgRect_Do(t *testing.T) {
	state := &State{}
	op := BgRect{X1: 0.1, Y1: 0.1, X2: 0.9, Y2: 0.9}

	op.Do(state)
	if state.BgRect == nil {
		t.Error("BgRect didn't set rectangle")
	}
}

func TestAddFigure_Do(t *testing.T) {
	state := &State{}
	op := AddFigure{X: 0.5, Y: 0.5}

	op.Do(state)
	if len(state.Figures) != 1 {
		t.Error("AddFigure didn't add figure")
	}
}

func TestMove_Do(t *testing.T) {
	state := &State{
		Figures: []Figure{{X: 0.1, Y: 0.1}},
	}
	op := Move{DX: 0.1, DY: 0.1}

	op.Do(state)
	if state.Figures[0].X != 0.2 || state.Figures[0].Y != 0.2 {
		t.Error("Move didn't move figures correctly")
	}
}

func TestReset_Do(t *testing.T) {
	state := &State{
		BgColor: color.White,
		BgRect:  &image.Rectangle{},
		Figures: []Figure{{X: 0.5, Y: 0.5}},
	}
	op := Reset{}

	op.Do(state)
	if state.BgColor != color.Black || state.BgRect != nil || len(state.Figures) != 0 {
		t.Error("Reset didn't reset state correctly")
	}
}

func TestUpdate_Do(t *testing.T) {
	op := Update{}
	if !op.Do(nil) {
		t.Error("Update should return true")
	}
}