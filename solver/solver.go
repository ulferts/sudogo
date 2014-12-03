package solver

import "github.com/ulferts/sudogo/structure"

func Solve(input_file string) {
	board := structure.NewBoard()

	organizer := Organizer{board}

	organizer.start(input_file)

	for {
		if board.Solved() {
			break
		}
	}

	board.Print()
}

type Solver struct {
	fields    []*structure.Field
	solutions fieldChannel
	input     fieldChannel
}

func (solver Solver) start() {
	for field := range solver.input {
		if solver.isResponsible(field) {
			solver.decreaseCandidates(field.Value)
			solver.setOnlyCandidate()
			solver.onlyPossibility()
		}
	}
}

func (solver Solver) decreaseCandidates(value int) {
	for _, field := range solver.fields {
		if field.Value == 0 {
			field.DecreaseCandidates(value)
		}
	}
}

func (solver Solver) setOnlyCandidate() {
	for _, field := range solver.fields {
		if field.Value == 0 && len(field.Candidates) == 1 {
			field.SetValue(field.Candidates[0])

			solver.solutions <- field
		}
	}
}

func (solver Solver) onlyPossibility() {
	counts := make(map[int]int)

	for _, field := range solver.fields {
		for _, val := range field.Candidates {
			counts[val]++
		}
		if field.Value != 0 {
			counts[field.Value]++
		}
	}

	for key, count := range counts {
		if count == 1 {
			for _, field := range solver.fields {

				contained := false

				for _, candidate := range field.Candidates {
					if candidate == key {
						contained = true
					}
				}

				if contained {
					field.SetValue(key)

					solver.solutions <- field
				}
			}
		}
	}
}

func (solver Solver) isResponsible(candidate *structure.Field) bool {
	ret := false

	for _, field := range solver.fields {
		if field == candidate {
			ret = true
		}
	}

	return ret
}
