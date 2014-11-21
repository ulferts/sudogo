package sudogo

//
//import "testing"
//
//func prepareSolver() ColumnSolver {
//	board := newBoard()
//	row := make(chan *Field)
//	column := make(chan *Field)
//	square := make(chan *Field)
//	solutions := make(chan *Field)
//
//	solver := ColumnSolver{board,
//		row,
//		column,
//		square,
//		solutions}
//
//	return solver
//}
//
//func TestColumnSolverColumns(t *testing.T) {
//	solver := prepareSolver()
//	x := 2
//
//	fields := solver.fields(x)
//
//	for y, field := range fields {
//		if field != &solver.board[x][y] {
//			t.Fail()
//		}
//	}
//}
