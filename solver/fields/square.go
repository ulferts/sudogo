package fields

import (
	"github.com/ulferts/sudogo/structure"
	"math"
)

type Square struct {
	Board structure.Board
}

func (square Square) Fields(field *structure.Field) []*structure.Field {
	x0 := field.X - int(math.Mod(float64(field.X), 3))
	y0 := field.Y - int(math.Mod(float64(field.Y), 3))
	fields := []*structure.Field{}

	for xc := 0; xc < 3; xc++ {
		for yc := 0; yc < 3; yc++ {
			fields = append(fields, &square.Board[x0+xc][y0+yc])
		}
	}

	return fields
}
