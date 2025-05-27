package ui

import (
	"image"
	"image/color"
	"testing"

	"golang.org/x/exp/shiny/screen"
	"golang.org/x/image/draw"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
)

type mockScreen struct{}

func (m mockScreen) NewTexture(size image.Point) (screen.Texture, error) {
	return &mockTexture{}, nil
}

func (m mockScreen) NewWindow(*screen.NewWindowOptions) (screen.Window, error) { return nil, nil }
func (m mockScreen) NewBuffer(image.Point) (screen.Buffer, error)              { return nil, nil }

type mockTexture struct{}

func (m *mockTexture) Release()                                           {}
func (m *mockTexture) Size() image.Point                                  { return image.Pt(800, 800) }
func (m *mockTexture) Bounds() image.Rectangle                            { return image.Rect(0, 0, 800, 800) }
func (m *mockTexture) Upload(image.Point, screen.Buffer, image.Rectangle) {}
func (m *mockTexture) Fill(rect image.Rectangle, src color.Color, op draw.Op) {}

func TestVisualizer(t *testing.T) {
	v := Visualizer{
		Title: "Test Window",
		OnScreenReady: func(s screen.Screen) {
			t.Log("Screen ready callback called")
		},
	}

	// Це лише базовий тест, що перевіряє ініціалізацію
	// Реальну роботу вікна важко тестувати без GUI
	go func() {
		v.run(mockScreen{})
	}()
}

func TestDetectTerminate(t *testing.T) {
	tests := []struct {
		name  string
		event any
		want  bool
	}{
		{
			name:  "lifecycle dead",
			event: lifecycle.Event{To: lifecycle.StageDead},
			want:  true,
		},
		{
			name:  "escape key",
			event: key.Event{Code: key.CodeEscape},
			want:  true,
		},
		{
			name:  "other key",
			event: key.Event{Code: key.CodeA},
			want:  false,
		},
		{
			name:  "other event",
			event: mouse.Event{},
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectTerminate(tt.event); got != tt.want {
				t.Errorf("detectTerminate() = %v, want %v", got, tt.want)
			}
		})
	}
}