package sudogo

import (
	"github.com/ulferts/sudogo/structure"
	"io/ioutil"
	"strconv"
	"strings"
)

type InputSolver struct {
	board     structure.Board
	solutions chan *structure.Field
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

				field.SetValue(value)

				solver.solutions <- field
			}
		}
	}
}
