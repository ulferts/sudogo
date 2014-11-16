package sudogo

import (
	//	"fmt"
	"sync"
)

func Solve(input_file string) {
	board := newBoard()

	row := make(chan *Field)
	column := make(chan *Field)
	square := make(chan *Field)

	input_solver := InputSolver{board, row, column, square}
	row_solver := RowSolver{board, row, column, square}
	column_solver := ColumnSolver{board, row, column, square}

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		// Decrement the counter when the goroutine completes.
		defer wg.Done()
		input_solver.start(input_file)
	}()

	go func() {
		// Decrement the counter when the goroutine completes.
		defer wg.Done()
		row_solver.start()
	}()

	go func() {
		// Decrement the counter when the goroutine completes.
		defer wg.Done()
		column_solver.start()
	}()

	wg.Wait()

	//close(row)
	//close(column)
	//close(square)
}

//func rowSolver(channel chan *Field) {
//	for f := range channel {
//		fmt.Println(f)
//	}
//}
