package sudogo

import "github.com/ulferts/sudogo/structure"
import "github.com/ulferts/sudogo/solver/fields"

func Solve(input_file string) {
	board := structure.NewBoard()

	input := make(chan *structure.Field, 100)
	output := make(chan *structure.Field, 100)
	intermediate := make(chan *structure.Field, 100)

	setupDecreaser(board, input, intermediate)
	setupPossibilityChecker(board, intermediate, output)
	setupFinalizer(board, output, input)

	input_solver := InputSolver{board, input}

	go input_solver.start(input_file)

	for {
		if board.Solved() {
			break
		}
	}

	board.Print()
}

func setupPossibilityChecker(board structure.Board, input chan *structure.Field, output chan *structure.Field) {
	row_channel := make(chan *structure.Field, 1000)
	column_channel := make(chan *structure.Field, 1000)
	square_channel := make(chan *structure.Field, 1000)

	row := PossibilityChecker{BoardChannel{row_channel, output, board},
		fields.Row{board}}
	column := PossibilityChecker{BoardChannel{column_channel, output, board},
		fields.Column{board}}
	square := PossibilityChecker{BoardChannel{square_channel, output, board},
		fields.Square{board}}

	go row.start()
	go column.start()
	go square.start()

	go distribute(board,
		row_channel,
		column_channel,
		square_channel,
		input)

}

func setupDecreaser(board structure.Board, input chan *structure.Field, output chan *structure.Field) {

	column_channel := make(chan *structure.Field, 1000)

	row_decreaser := Decreaser{BoardChannel{input, column_channel, board}, fields.Row{board}}

	square_channel := make(chan *structure.Field, 1000)

	column_decreaser := Decreaser{BoardChannel{column_channel, square_channel, board}, fields.Column{board}}

	square_decreaser := Decreaser{BoardChannel{square_channel, output, board}, fields.Square{board}}

	go row_decreaser.start()
	go column_decreaser.start()
	go square_decreaser.start()
}

func distribute(board structure.Board, row, column, square, input chan *structure.Field) {
	for field := range input {
		square <- field
		row <- field
		column <- field
	}
}

func setupFinalizer(board structure.Board, input chan *structure.Field, output chan *structure.Field) {

	finalizer := Finalizer{BoardChannel{input, output, board}, fields.Board{board}}

	go finalizer.start()
}

type BoardChannel struct {
	input     chan *structure.Field
	solutions chan *structure.Field
	board     structure.Board
}
