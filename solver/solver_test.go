package sudogo

import "testing"
import "github.com/ulferts/sudogo/structure"
import "fmt"

func setupSolver() Solver {
	input := make(chan *structure.Field, 100)
	output := make(chan *structure.Field, 100)
	board := structure.NewBoard()

	solver := Solver{input, output, board}

	return solver
}

func TestDecreaseCandidates(t *testing.T) {
	solver := setupSolver()

	field1 := &structure.Field{0, 0, 1, []int{}}
	field2 := &structure.Field{0, 1, 0, []int{2, 3, 4, 5, 6, 7, 8, 9}}
	field3 := &structure.Field{0, 2, 0, []int{2, 3}}

	fields := []*structure.Field{field1, field2, field3}

	solver.decreaseCandidates(fields, 3)

	//result := <-input

	if len(field3.Candidates) != 0 {
		t.Fail()
	}

	if field3.Value != 2 {
		t.Fail()
	}

	//	if field3 != result {
	//		t.Fail()
	//	}

	// 2 is missing because of field3 having had only 2, 3
	// as candidates now has the value 2
	expectedField2 := []int{4, 5, 6, 7, 8, 9}

	for k, v := range expectedField2 {
		if field2.Candidates[k] != v {
			t.Fail()
		}
	}

	if field1.Value != 1 {
		t.Fail()
	}

	if len(field1.Candidates) != 0 {
		t.Fail()
	}

	//	close(input)
	//	close(output)
}

func TestOnlyPossibility(t *testing.T) {
	solver := setupSolver()

	field1 := &structure.Field{0, 0, 1, []int{}}
	field2 := &structure.Field{0, 1, 0, []int{3, 4, 5, 6}}
	field3 := &structure.Field{0, 2, 0, []int{2, 3}}
	field4 := &structure.Field{0, 3, 0, []int{4, 5, 6}}
	field5 := &structure.Field{0, 3, 0, []int{4, 5, 6}}

	fields := []*structure.Field{field1, field2, field3, field4, field5}

	solver.onlyPossibility(fields, 1)

	if field3.Value != 2 {
		t.Fail()
	}

	expectedField4 := []int{4, 5, 6}
	for k, v := range expectedField4 {
		fmt.Println(field4.Candidates)
		if field4.Candidates[k] != v {
			t.Fail()
		}
	}
}
