package sudogo

type ColumnSolver struct {
	board               Board
	row, column, square chan *Field
}

func (solver ColumnSolver) start() {
	for field := range solver.column {
		solver.decreaseCandidates(field.dy, field.value)
	}
}

func (solver ColumnSolver) decreaseCandidates(x, value int) {
	for y := 0; y < 9; y++ {
		field := &solver.board[x][y]

		if field.value == 0 {
			field.decreaseCandidates(value)
		}
	}
}
