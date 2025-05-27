package painter

import (
	"image"
	"image/color"

	"golang.org/x/exp/shiny/screen"
)

type Receiver interface {
	Update(t screen.Texture)
}

type Loop struct {
	Receiver Receiver
	state    State
	next     screen.Texture
	prev     screen.Texture
	stop     chan struct{}
}

func (l *Loop) Start(s screen.Screen) {
	l.state = State{BgColor: color.Black}
	l.next, _ = s.NewTexture(image.Pt(800, 800))
	l.prev, _ = s.NewTexture(image.Pt(800, 800))
	l.stop = make(chan struct{})
	go l.eventLoop()
}

func (l *Loop) Post(op Operation) {
	if op.Do(&l.state) { // Викликати render() ТІЛЬКИ якщо операція повертає true
		l.render()
	}
}

func (l *Loop) StopAndWait() {
	close(l.stop)
}

func (l *Loop) render() {
	l.next.Fill(l.next.Bounds(), l.state.BgColor, screen.Src)

	if l.state.BgRect != nil {
		l.next.Fill(*l.state.BgRect, color.Black, screen.Over)
	}

	for _, fig := range l.state.Figures {
		centerX := int(fig.X * 800)
		centerY := int(fig.Y * 800)
		size := 100

		vertical := image.Rect(centerX-25, centerY-size, centerX+25, centerY+size)
		l.next.Fill(vertical, color.RGBA{R: 255, G: 255, A: 255}, screen.Over)

		horizontal := image.Rect(centerX-size, centerY-25, centerX+size, centerY+25)
		l.next.Fill(horizontal, color.RGBA{R: 255, G: 255, A: 255}, screen.Over)
	}

	l.Receiver.Update(l.next)
	l.next, l.prev = l.prev, l.next
}

func (l *Loop) eventLoop() {
	<-l.stop
}
