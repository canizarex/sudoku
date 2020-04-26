package main

import ( "fmt"
         "strings"
		 "time"
)

const size = 9

var sudoku = [][]int{{4, 0, 0, 0, 9, 5, 0, 0, 0},
	{1, 0, 0, 6, 0, 0, 8, 5, 2},
	{2, 0, 0, 0, 0, 0, 0, 0, 7},
	{0, 9, 0, 0, 0, 1, 0, 2, 0},
	{0, 8, 0, 0, 0, 2, 9, 4, 0},
	{0, 0, 0, 0, 5, 3, 0, 0, 0},
	{9, 0, 3, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 4, 0, 0, 1, 7, 9},
	{0, 0, 6, 1, 0, 0, 2, 0, 0}}

func printMatrix(m [][]int) {
	for _, row := range m {
		fmt.Println(row)
	}
	fmt.Println(strings.Repeat("-", 19))
}

func possible(y, x, n int) bool {

	// Check all the numbers in a given row
	for i := 0; i < size; i++ {
		if sudoku[y][i] == n {
			return false
		}
	}

	// Check all the numbers in a given column
	for i := 0; i < size; i++ {
		if sudoku[i][x] == n {
			return false
		}
	}

	// Check a 3:3 subgrid
	x0 := (x / 3) * 3
	y0 := (y / 3) * 3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if n == sudoku[y0+i][x0+j] {
				return false
			}
		}
	}
	return true
}

func solve(sudoku [][]int) {
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			// Go ahead only if the box is empty (equals zero)
			if sudoku[y][x] != 0 {
				continue
			}
			// For every n check if it is allowed in a given box
			// and if it is, call the function recursively to
			// start again.
			for n := 1; n < 10; n++ {
				if possible(y, x, n) {
					sudoku[y][x] = n
					solve(sudoku)
					// At this point the recursive function has returned
					// because there were no more possibilities so
					// it takes a step back and re-write the last written
					// box with a zero.
					sudoku[y][x] = 0
				}
			}
			// Te recursive function returns here when none n is allowed.
			return 
		}
	}
	printMatrix(sudoku)
}

func main() {
	printMatrix(sudoku)
	start := time.Now()
	solve(sudoku)
	elapsed := time.Since(start)
	fmt.Printf("It took %s to solve the sudoku\n", elapsed)
}
