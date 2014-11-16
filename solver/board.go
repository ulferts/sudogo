package sudogo

type Board [][]Field

func newBoard() Board {
	var board Board = make([][]Field, 9)

	for x := 0; x < 9; x++ {
		board[x] = make([]Field, 9)

		for y := 0; y < 9; y++ {
			board[x][y] = newField(x, y, 0)
		}
	}

	return board
}
