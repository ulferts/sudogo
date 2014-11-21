package fields

import "github.com/ulferts/sudogo/structure"

type Board struct {
	BoardData structure.Board
}

func (board Board) Fields(field *structure.Field) []*structure.Field {
	fields := []*structure.Field{}
	length := len(board.BoardData[0])

	for x := 0; x < length; x++ {
		for y := 0; y < length; y++ {
			fields = append(fields, &board.BoardData[x][y])
		}
	}

	return fields
}
