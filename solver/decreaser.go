package sudogo

import "github.com/ulferts/sudogo/solver/fields"
import "github.com/ulferts/sudogo/structure"

type Decreaser struct {
	BoardChannel
	fields.Fielded
}

func (decreaser Decreaser) start() {
	for field := range decreaser.input {
		fields := decreaser.Fields(field)
		decreaser.decreaseCandidates(fields, field.Value)

		decreaser.solutions <- field
	}
}

func (decreaser Decreaser) decreaseCandidates(fields []*structure.Field, value int) {
	for _, field := range fields {
		if field.Value == 0 {
			field.DecreaseCandidates(value)
		}
	}
}
