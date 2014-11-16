package sudogo

type RowSolver struct {
	board               Board
	row, column, square chan *Field
}

func (solver RowSolver) start() {
	for field := range solver.row {
		solver.decreaseCandidates(field.dy, field.value)
	}
}

func (solver RowSolver) decreaseCandidates(y, value int) {
	for x := 0; x < 9; x++ {
		field := &solver.board[x][y]

		if field.value == 0 {
			field.decreaseCandidates(value)
		}
	}
}
