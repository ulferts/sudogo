package fields

import "github.com/ulferts/sudogo/structure"

type Row struct {
	Board structure.Board
}

func (row Row) Fields(field *structure.Field) []*structure.Field {
	fields := []*structure.Field{}
	length := len(row.Board[0])

	for x := 0; x < length; x++ {
		fields = append(fields, &row.Board[x][field.Y])
	}

	return fields
}
