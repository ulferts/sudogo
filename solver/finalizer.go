package sudogo

import "github.com/ulferts/sudogo/solver/fields"
import "github.com/ulferts/sudogo/structure"

type Finalizer struct {
	BoardChannel
	fields.Fielded
}

func (finalizer Finalizer) start() {
	for field := range finalizer.input {
		fields := finalizer.Fields(field)
		finalizer.setOnlyCandidate(fields)

		finalizer.solutions <- field
	}
}

func (finalizer Finalizer) setOnlyCandidate(fields []*structure.Field) {
	for _, field := range fields {
		if field.Value == 0 && len(field.Candidates) == 1 {
			field.SetValue(field.Candidates[0])

			finalizer.solutions <- field
		}
	}
}
