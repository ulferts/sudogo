package sudogo

type LineSolver struct {
	board               Board
	row, column, square chan *Field
}

func (solver LineSolver) start() {
}
