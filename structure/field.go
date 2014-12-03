package structure

import "sync"

type Field struct {
	X, Y, Value int
	Candidates  []int
	lock        sync.Mutex
}

func newField(x, y, value int) Field {
	if value != 0 {
		return Field{X: x, Y: y, Value: value, Candidates: []int{}}
	} else {
		return Field{X: x, Y: y, Value: value, Candidates: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}}
	}
}

func (f *Field) SetValue(v int) {
	f.lock.Lock()

	f.Candidates = []int{}

	f.Value = v

	f.lock.Unlock()
}

func (f *Field) DecreaseCandidates(value int) bool {
	f.lock.Lock()

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

	f.lock.Unlock()

	return ret
}
