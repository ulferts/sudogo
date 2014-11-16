package sudogo

import (
	"fmt"
)

type Field struct {
	dx, dy, value int
	candidates    []int
}

func newField(x, y, value int) Field {
	if value != 0 {
		return Field{x, y, value, []int{}}
	} else {
		return Field{x, y, value, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}}
	}
}

func (f *Field) setValue(v int) {
	f.candidates = []int{}

	f.value = v
}

func (f *Field) decreaseCandidates(value int) {
	currentCandidates := f.candidates

	posOfValue := -1

	for i, candidate := range currentCandidates {
		if candidate == value {
			posOfValue = i
		}
	}

	if posOfValue >= 0 {
		newCandidates := append(currentCandidates[:posOfValue], currentCandidates[posOfValue+1:]...)
		f.candidates = newCandidates
	}
	fmt.Println(f.candidates)

	if len(f.candidates) == 1 && f.value == 0 {
		fmt.Println("New value found")
		f.value = f.candidates[0]
		f.candidates = []int{}
	}
}
