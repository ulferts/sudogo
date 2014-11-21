package sudogo

import (
	"github.com/ulferts/sudogo/solver/fields"
	"github.com/ulferts/sudogo/structure"
)

type PossibilityChecker struct {
	BoardChannel
	fields.Fielded
}

func (checker PossibilityChecker) start() {
	for field := range checker.input {
		fields := checker.Fields(field)
		checker.onlyPossibility(fields, field.Value)

		checker.solutions <- field
	}
}

func (checker PossibilityChecker) onlyPossibility(fields []*structure.Field, value int) {
	counts := make(map[int]int)

	for _, field := range fields {
		for _, val := range field.Candidates {
			counts[val]++
		}
		if field.Value != 0 {
			counts[field.Value]++
		}
	}

	for key, count := range counts {
		if count == 1 {
			for _, field := range fields {

				contained := false

				for _, candidate := range field.Candidates {
					if candidate == key {
						contained = true
					}
				}

				if contained {
					field.SetValue(key)

					checker.solutions <- field
				}
			}
		}
	}
}
