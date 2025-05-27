package painter

import (
	"image"
	"image/color"
	"image/draw"
	"testing"

	"golang.org/x/exp/shiny/screen"
)

type testReceiver struct {
	lastTexture screen.Texture
}

func (tr *testReceiver) Update(t screen.Texture) {
	tr.lastTexture = t
}

func TestLoop_Post(t *testing.T) {
	var (
		l  Loop
		tr testReceiver
	)
	l.Receiver = &tr

	l.Start(mockScreen{})
	l.Post(WhiteFill{})
	l.Post(GreenFill{})
	l.Post(Update{})
	l.StopAndWait()

	// Перевірка результатів...
}

type mockScreen struct{}

func (m mockScreen) NewTexture(size image.Point) (screen.Texture, error) {
	return &mockTexture{}, nil
}

func (m mockScreen) NewWindow(*screen.NewWindowOptions) (screen.Window, error) { return nil, nil }
func (m mockScreen) NewBuffer(image.Point) (screen.Buffer, error)              { return nil, nil }

type mockTexture struct {
	Colors []color.Color
}

func (m *mockTexture) Release()                                           {}
func (m *mockTexture) Size() image.Point                                  { return image.Pt(800, 800) }
func (m *mockTexture) Bounds() image.Rectangle                            { return image.Rect(0, 0, 800, 800) }
func (m *mockTexture) Upload(image.Point, screen.Buffer, image.Rectangle) {}
func (m *mockTexture) Fill(rect image.Rectangle, src color.Color, op draw.Op) {
	m.Colors = append(m.Colors, src)
}