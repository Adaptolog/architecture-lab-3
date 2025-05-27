package lang

import (
	"strings"
	"testing"

	"github.com/Adaptolog/architecture-lab-3/painter"
)

func TestParser_Parse(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr bool
	}{
		{"empty", "", 0, false},
		{"white", "white", 1, false},
		{"green", "green", 1, false},
		{"bgrect valid", "bgrect 0.1 0.1 0.9 0.9", 1, false},
		{"bgrect invalid", "bgrect 0.1 0.1", 0, true},
		{"figure valid", "figure 0.5 0.5", 1, false},
		{"figure invalid", "figure 0.5", 0, true},
		{"move valid", "move 0.1 0.1", 1, false},
		{"move invalid", "move 0.1", 0, true},
		{"reset", "reset", 1, false},
		{"update", "update", 1, false},
		{"unknown", "unknown", 0, true},
		{"multiple commands", "white\ngreen\nupdate", 3, false},
	}

	p := &Parser{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := p.Parse(strings.NewReader(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("Parse() got = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func TestParser_CommandTypes(t *testing.T) {
	p := &Parser{}
	ops, err := p.Parse(strings.NewReader("white\nbgrect 0.1 0.1 0.9 0.9\nfigure 0.5 0.5\nmove 0.1 0.1\nreset\nupdate"))
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if _, ok := ops[0].(painter.WhiteFill); !ok {
		t.Error("First operation should be WhiteFill")
	}
	if _, ok := ops[1].(painter.BgRect); !ok {
		t.Error("Second operation should be BgRect")
	}
	if _, ok := ops[2].(painter.AddFigure); !ok {
		t.Error("Third operation should be AddFigure")
	}
	if _, ok := ops[3].(painter.Move); !ok {
		t.Error("Fourth operation should be Move")
	}
	if _, ok := ops[4].(painter.Reset); !ok {
		t.Error("Fifth operation should be Reset")
	}
	if _, ok := ops[5].(painter.Update); !ok {
		t.Error("Sixth operation should be Update")
	}
}