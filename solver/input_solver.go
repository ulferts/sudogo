package sudogo

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type InputSolver struct {
	board               Board
	row, column, square chan *Field
}

func (solver InputSolver) start(path string) {
	dat, _ := ioutil.ReadFile(path)

	input := string(dat)
	lines := strings.Split(input, "\n")

	for x := 0; x < 9; x++ {
		fields := strings.Split(lines[x], "|")

		for y := 0; y < 9; y++ {
			value, err := strconv.Atoi(fields[y])

			if err == nil {
				field := &solver.board[x][y]

				field.setValue(value)

				solver.row <- field
				solver.column <- field
				//solver.square <- field
			}
		}
	}

	close(solver.column)
	close(solver.square)
	close(solver.row)
}
