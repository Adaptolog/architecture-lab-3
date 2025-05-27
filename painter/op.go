package painter

import (
	"image"
	"image/color"
)

type Operation interface {
	Do(state *State) (ready bool)
}

type OperationList []Operation

func (ol OperationList) Do(state *State) (ready bool) {
	for _, op := range ol {
		ready = op.Do(state) || ready
	}
	return
}

type State struct {
	BgColor color.Color
	BgRect  *image.Rectangle
	Figures []Figure
}

type Figure struct {
	X, Y float32
}

type WhiteFill struct{}

func (op WhiteFill) Do(state *State) bool {
	state.BgColor = color.White
	return false // Не викликає негайне оновлення
}

type GreenFill struct{}

func (op GreenFill) Do(state *State) bool {
	state.BgColor = color.RGBA{G: 0xff, A: 0xff}
	return false
}

type BgRect struct {
	X1, Y1, X2, Y2 float32
}

func (op BgRect) Do(state *State) bool {
	if op.X1 >= op.X2 || op.Y1 >= op.Y2 {
		return false
	}
	rect := image.Rect(
		int(op.X1*800), int(op.Y1*800),
		int(op.X2*800), int(op.Y2*800),
	)
	state.BgRect = &rect
	return false
}

type AddFigure struct {
	X, Y float32
}

func (op AddFigure) Do(state *State) bool {
	state.Figures = append(state.Figures, Figure{X: op.X, Y: op.Y})
	return false
}

type Move struct {
	DX, DY float32
}

func (op Move) Do(state *State) bool {
	for i := range state.Figures {
		state.Figures[i].X += op.DX
		state.Figures[i].Y += op.DY
	}
	return false
}

type Reset struct{}

func (op Reset) Do(state *State) bool {
	state.BgColor = color.Black
	state.BgRect = nil
	state.Figures = nil
	return false
}

type Update struct{}

func (op Update) Do(state *State) bool {
	return true // Сигналізує про необхідність оновлення
}
