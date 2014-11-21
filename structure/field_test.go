package structure

import "testing"

func TestDecreaseCandidates(t *testing.T) {
	f := newField(0, 0, 0)

	f.DecreaseCandidates(5)
	f.DecreaseCandidates(1)
	f.DecreaseCandidates(7)

	expected := []int{2, 3, 4, 6, 8, 9}
	for k, v := range expected {
		if f.Candidates[k] != v {
			t.Fail()
		}
	}
}

func TestSetValue(t *testing.T) {
	f := newField(0, 0, 0)

	f.SetValue(5)

	if len(f.Candidates) != 0 {
		t.Fail()
	}

	if f.Value != 5 {
		t.Fail()
	}
}
