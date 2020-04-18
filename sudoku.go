package main

import "fmt"

var sudoku = [][]int{{4, 0, 0, 0, 9, 5, 0, 0, 0},
	{1, 0, 0, 6, 0, 0, 8, 5, 2},
	{2, 0, 0, 0, 0, 0, 0, 0, 7},
	{0, 9, 0, 0, 0, 1, 0, 2, 0},
	{0, 8, 0, 0, 0, 2, 9, 4, 0},
	{0, 0, 0, 0, 5, 3, 0, 0, 0},
	{9, 0, 3, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 4, 0, 0, 1, 7, 9},
	{0, 0, 6, 1, 0, 0, 2, 0, 0}}

func printMatrix() {
	for _, row := range sudoku {
		fmt.Println(row)
	}
	fmt.Println("--------------------")
}

func possible(y, x, n int) bool {

	// Check all the numbers in a given row
	for i := 0; i < 9; i++ {
		if sudoku[y][i] == n {
			return false
		}
	}

	// Check all the numbers in a given column
	for i := 0; i < 9; i++ {
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

func solve() {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
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
					solve()
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
	printMatrix()
}

func main() {
	printMatrix()
	solve()
}
