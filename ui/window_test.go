package ui

import (
	"image"
	"image/color"
	"image/draw"
	"testing"
	"time"

	"golang.org/x/exp/shiny/screen"
	"golang.org/x/image/math/f64"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
)

type mockScreen struct{}

func (m mockScreen) NewTexture(size image.Point) (screen.Texture, error) {
	return &mockTexture{}, nil
}

func (m mockScreen) NewWindow(*screen.NewWindowOptions) (screen.Window, error) {
	return &mockWindow{}, nil
}

func (m mockScreen) NewBuffer(image.Point) (screen.Buffer, error) {
	return nil, nil
}

type mockTexture struct{}

func (m *mockTexture) Release() {}
func (m *mockTexture) Size() image.Point {
	return image.Pt(800, 800)
}
func (m *mockTexture) Bounds() image.Rectangle {
	return image.Rect(0, 0, 800, 800)
}
func (m *mockTexture) Upload(image.Point, screen.Buffer, image.Rectangle) {}
func (m *mockTexture) Fill(rect image.Rectangle, src color.Color, op draw.Op) {}

type mockWindow struct{}

func (m *mockWindow) Release() {}

func (m *mockWindow) NextEvent() interface{} {
	// Імітуємо подію закриття вікна через короткий час
	time.Sleep(100 * time.Millisecond)
	return lifecycle.Event{To: lifecycle.StageDead}
}

func (m *mockWindow) Publish() screen.PublishResult {
	return screen.PublishResult{}
}

func (m *mockWindow) Send(event interface{}) {}

func (m *mockWindow) SendFirst(event interface{}) {}

func (m *mockWindow) Upload(dst image.Point, src screen.Buffer, sr image.Rectangle) {}

func (m *mockWindow) Scale(dst image.Rectangle, src screen.Texture, sr image.Rectangle, op draw.Op, opts *screen.DrawOptions) {}

func (m *mockWindow) Copy(dst image.Point, src screen.Texture, sr image.Rectangle, op draw.Op, opts *screen.DrawOptions) {}

func (m *mockWindow) Draw(src2d f64.Aff3, src screen.Texture, sr image.Rectangle, op draw.Op, opts *screen.DrawOptions) {}

func (m *mockWindow) DrawUniform(src2d f64.Aff3, src color.Color, sr image.Rectangle, op draw.Op, opts *screen.DrawOptions) {}

func (m *mockWindow) Fill(rect image.Rectangle, src color.Color, op draw.Op) {}

func TestVisualizer(t *testing.T) {
	v := Visualizer{
		Title: "Test Window",
		done:  make(chan struct{}), // Явно ініціалізуємо канал
		OnScreenReady: func(s screen.Screen) {
			t.Log("Screen ready callback called")
		},
	}

	// Додаємо таймаут для тесту
	timeout := time.After(5 * time.Second)
	done := make(chan struct{})

	go func() {
		defer close(done)
		v.run(mockScreen{})
	}()

	select {
	case <-done:
		// Тест завершився успішно
		t.Log("Visualizer completed successfully")
	case <-timeout:
		// Примусово завершуємо тест після таймауту
		t.Fatal("Test timed out: Visualizer did not complete in time")
	case <-v.done:
		// Отримали сигнал завершення
		t.Log("Received done signal")
	}
}

func TestDetectTerminate(t *testing.T) {
	tests := []struct {
		name  string
		event interface{}
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