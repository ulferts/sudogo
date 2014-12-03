package fields

import "github.com/ulferts/sudogo/structure"

type Column struct {
	Board structure.Board
}

func (column Column) Fields(field *structure.Field) []*structure.Field {
	fields := []*structure.Field{}

	for y, _ := range column.Board[field.X] {
		fields = append(fields, &column.Board[field.X][y])
	}

	return fields
}

func (column Column) Of(x int) []*structure.Field {
	fields := []*structure.Field{}

	for y, _ := range column.Board[x] {
		fields = append(fields, &column.Board[x][y])
	}

	return fields
}
