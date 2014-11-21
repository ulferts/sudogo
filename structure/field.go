package structure

type Field struct {
	X, Y, Value int
	Candidates  []int
}

func newField(x, y, value int) Field {
	if value != 0 {
		return Field{x, y, value, []int{}}
	} else {
		return Field{x, y, value, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}}
	}
}

func (f *Field) SetValue(v int) {
	f.Candidates = []int{}

	f.Value = v
}

func (f *Field) DecreaseCandidates(value int) bool {
	currentCandidates := f.Candidates
	ret := false

	posOfValue := -1

	for i, candidate := range currentCandidates {
		if candidate == value {
			posOfValue = i
		}
	}

	if posOfValue >= 0 {
		newCandidates := append(currentCandidates[:posOfValue], currentCandidates[posOfValue+1:]...)
		f.Candidates = newCandidates
		ret = true
	}

	return ret
}
