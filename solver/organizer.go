package solver

import "github.com/ulferts/sudogo/structure"
import "github.com/ulferts/sudogo/solver/fields"

type Organizer struct {
	board structure.Board
	//	solvers *Solver
}

func (organizer Organizer) start(input_file string) {
	solutions := make(fieldChannel, 100)
	inputs := []fieldChannel{}

	row_inputs := organizer.setupRowSolvers(solutions)

	for _, channel := range row_inputs {
		inputs = append(inputs, channel)
	}

	column_inputs := organizer.setupColumnSolvers(solutions)

	for _, channel := range column_inputs {
		inputs = append(inputs, channel)
	}

	square_inputs := organizer.setupSquareSolvers(solutions)

	for _, channel := range square_inputs {
		inputs = append(inputs, channel)
	}

	go organizer.distribute(inputs, solutions)

	input_solver := InputSolver{organizer.board, solutions}

	go input_solver.start(input_file)
}

func (organizer Organizer) distribute(inputs []fieldChannel, solutions fieldChannel) {
	for field := range solutions {
		for _, input := range inputs {
			input <- field
		}
	}
}

func (organizer Organizer) setupRowSolvers(solutions fieldChannel) []fieldChannel {
	fieldsets := []fields.Fieldset{}

	factory := fields.Row{organizer.board}

	for x := 0; x < len(organizer.board); x++ {
		fieldsets = append(fieldsets, factory.Of(x))
	}

	inputs := organizer.setupSolversForFields(fieldsets, solutions)

	return inputs
}

func (organizer Organizer) setupColumnSolvers(solutions fieldChannel) []fieldChannel {
	fieldsets := []fields.Fieldset{}

	factory := fields.Column{organizer.board}

	for x := 0; x < len(organizer.board); x++ {
		fieldsets = append(fieldsets, factory.Of(x))
	}

	inputs := organizer.setupSolversForFields(fieldsets, solutions)

	return inputs
}

func (organizer Organizer) setupSquareSolvers(solutions fieldChannel) []fieldChannel {
	fieldsets := []fields.Fieldset{}

	factory := fields.Square{organizer.board}

	for x := 0; x < (len(organizer.board) / 3); x++ {
		for y := 0; y < (len(organizer.board) / 3); y++ {
			fieldsets = append(fieldsets, factory.Of(x, y))
		}
	}

	inputs := organizer.setupSolversForFields(fieldsets, solutions)

	return inputs
}

func (organizer Organizer) setupSolversForFields(fieldsets []fields.Fieldset, solutions fieldChannel) []fieldChannel {
	inputs := []fieldChannel{}

	for _, fieldset := range fieldsets {
		input := make(fieldChannel, 100)

		solver := Solver{fieldset, solutions, input}

		go solver.start()

		inputs = append(inputs, input)
	}

	return inputs

}
