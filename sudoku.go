package main

import (
	"fmt"
// 	"os"
//	"os/exec" 
	"strings"
	"time"
)

const size = 9

var sudoku = [][]int{{8, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 3, 6, 0, 0, 0, 0, 0},
	{0, 7, 0, 0, 9, 0, 2, 0, 0},
	{0, 5, 0, 0, 0, 7, 0, 0, 0},
	{0, 0, 0, 0, 4, 5, 7, 0, 0},
	{0, 0, 0, 1, 0, 0, 0, 3, 0},
	{0, 0, 1, 0, 0, 0, 0, 6, 8},
	{0, 0, 8, 5, 0, 0, 0, 1, 0},
	{0, 9, 0, 0, 0, 0, 4, 0, 0}}

func printMatrix(m [][]int) {
	vsep, hsep, xsep := "|", "-", "+"
	hline := fmt.Sprintf("%s", strings.Repeat(hsep, 25))
	hlinex := fmt.Sprintf("%[1]s%[2]s%[3]s%[2]s%[3]s%[2]s%[1]s", vsep, strings.Repeat(hsep,7), xsep)

	for i, row := range m {
		if i == 0 {
			fmt.Printf("%s\n", hline)
		}
		for i, n := range row {
			switch {
			case i == 0:
				fmt.Printf("%s%2d", vsep, n)
			case (i+1) % 3 == 0:
				fmt.Printf("%2d%2s", n, vsep)
			default:
				fmt.Printf("%2d", n)
			}
		} 
        switch {
		case i == 8:
			fmt.Printf("\n%s\n", hline)
		case (i+1) % 3 == 0:
			fmt.Printf("\n%s\n", hlinex)
		default:
			fmt.Println()
		}
	} 
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

var count int 
var solved bool = false

func solve(sudoku [][]int) {

	count++
/*  
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	printMatrix(sudoku)
	fmt.Println(count)
	time.Sleep(100 * time.Millisecond)
 */
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
					// box with a zero. To avoid undoing all the changes
					// once solved, we have to add a check first.
					if solved {
						return
					}
					sudoku[y][x] = 0
				}
			}
			// Te recursive function returns here when none n is allowed.
			return 
		}
	}
	// This point is reached only when all boxes are different than 0.
	solved = true
}

func main() {
	fmt.Println("Sudoku to be solved:")
	printMatrix(sudoku)

	start := time.Now()
	solve(sudoku)
	elapsed := time.Since(start)

	fmt.Println("Solution:")
	printMatrix(sudoku)
	fmt.Printf("It took %s and %d iterations to solve the sudoku\n", elapsed, count)
}
