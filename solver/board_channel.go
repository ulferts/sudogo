package solver

import "github.com/ulferts/sudogo/structure"

type BoardChannel struct {
	input     chan *structure.Field
	solutions chan *structure.Field
	board     structure.Board
}
