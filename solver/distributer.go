package sudogo

import "github.com/ulferts/sudogo/structure"

type Distributer struct {
	board               structure.Board
	row, column, square chan *structure.Field
	solutions           chan *structure.Field
}

func (d *Distributer) start() {
	for field := range d.solutions {
		d.row <- field
		d.column <- field
		d.square <- field

		if d.board.Solved() {
			d.stop()
		}
	}
}

func (d *Distributer) stop() {

	close(d.column)
	close(d.square)
	close(d.row)
	close(d.solutions)
}
