package structure

import (
	"bytes"
	"fmt"
	"strconv"
)

type Board [][]Field

func NewBoard() Board {
	var board Board = make([][]Field, 9)

	for x := 0; x < 9; x++ {
		board[x] = make([]Field, 9)

		for y := 0; y < 9; y++ {
			board[x][y] = newField(x, y, 0)
		}
	}

	return board
}

func (b Board) Solved() bool {

	solved_fields := 0

	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if b[x][y].Value != 0 {
				solved_fields++
			}
		}
	}

	if solved_fields == 81 {
		return true
	} else {
		return false
	}
}

func (b Board) Print() {
	fmt.Println("---------")
	for x := 0; x < 9; x++ {
		var line bytes.Buffer
		line.WriteString("|")

		for y := 0; y < 9; y++ {
			field := b[x][y]

			stringValue := "-"
			if field.Value > 0 {
				stringValue = strconv.Itoa(field.Value)
			}
			line.WriteString(stringValue)
			line.WriteString("|")
		}

		fmt.Println(line.String())
	}
	fmt.Println("---------")
}
