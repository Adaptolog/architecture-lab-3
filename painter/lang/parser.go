package lang

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/Adaptolog/architecture-lab-3/painter"
)

type Parser struct{}

func (p *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
	scanner := bufio.NewScanner(in)
	var ops []painter.Operation

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}

		switch fields[0] {
		case "white":
			ops = append(ops, painter.WhiteFill{})
		case "green":
			ops = append(ops, painter.GreenFill{})
		case "bgrect":
			if len(fields) != 5 {
				return nil, fmt.Errorf("bgrect requires 4 arguments")
			}
			x1, _ := strconv.ParseFloat(fields[1], 32)
			y1, _ := strconv.ParseFloat(fields[2], 32)
			x2, _ := strconv.ParseFloat(fields[3], 32)
			y2, _ := strconv.ParseFloat(fields[4], 32)
			ops = append(ops, painter.BgRect{
				X1: float32(x1), Y1: float32(y1),
				X2: float32(x2), Y2: float32(y2),
			})
		case "figure":
			if len(fields) != 3 {
				return nil, fmt.Errorf("figure requires 2 arguments")
			}
			x, _ := strconv.ParseFloat(fields[1], 32)
			y, _ := strconv.ParseFloat(fields[2], 32)
			ops = append(ops, painter.AddFigure{
				X: float32(x), Y: float32(y),
			})
		case "move":
			if len(fields) != 3 {
				return nil, fmt.Errorf("move requires 2 arguments")
			}
			dx, _ := strconv.ParseFloat(fields[1], 32)
			dy, _ := strconv.ParseFloat(fields[2], 32)
			ops = append(ops, painter.Move{
				DX: float32(dx), DY: float32(dy),
			})
		case "reset":
			ops = append(ops, painter.Reset{})
		case "update":
			ops = append(ops, painter.Update{}) // Додає операцію оновлення
		default:
			return nil, fmt.Errorf("unknown command: %s", fields[0])
		}
	}

	return ops, nil
}
