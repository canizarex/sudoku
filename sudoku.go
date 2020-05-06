package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

const size = 9

var verbose bool

type sudoku struct {
	board  [9][9]int
	count  int
	solved bool
}

func newSudoku(matrix [9][9]int) sudoku {
	return sudoku{matrix, 0, false}
}

func (s *sudoku) print() {
	vsep, hsep, xsep := "|", "―", "+"
	hline := fmt.Sprintf("%[1]s%[2]s%[1]s", " ", strings.Repeat(hsep, 23))
	hlinex := fmt.Sprintf("%[1]s%[2]s%[3]s%[2]s%[3]s%[2]s%[1]s", vsep, strings.Repeat(hsep, 7), xsep)

	for i, row := range s.board {
		if i == 0 {
			fmt.Printf("%s\n", hline)
		}
		for i, n := range row {
			switch {
			case i == 0:
				fmt.Printf("%s%2d", vsep, n)
			case (i+1)%3 == 0:
				fmt.Printf("%2d%2s", n, vsep)
			default:
				fmt.Printf("%2d", n)
			}
		}
		switch {
		case i == size-1:
			fmt.Printf("\n%s\n", hline)
		case (i+1)%3 == 0:
			fmt.Printf("\n%s\n", hlinex)
		default:
			fmt.Println()
		}
	}
}

func (s *sudoku) possible(y, x, n int) bool {

	// Check all the numbers in a given row
	for i := 0; i < size; i++ {
		if s.board[y][i] == n {
			return false
		}
	}

	// Check all the numbers in a given column
	for i := 0; i < size; i++ {
		if s.board[i][x] == n {
			return false
		}
	}

	// Check a 3:3 subgrid
	x0 := (x / 3) * 3
	y0 := (y / 3) * 3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if n == s.board[y0+i][x0+j] {
				return false
			}
		}
	}
	return true
}

func (s *sudoku) solve() {

	s.count++

	if verbose {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		s.print()
		fmt.Println(s.count)
	}

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			// Go ahead only if the box is empty (equals zero)
			if s.board[y][x] != 0 {
				continue
			}
			// For every n check if it is allowed in a given box
			// and if it is, call the function recursively to
			// start again.
			for n := 1; n < 10; n++ {
				if s.possible(y, x, n) {
					s.board[y][x] = n
					s.solve()
					// At this point the recursive function has returned
					// because there were no more possibilities so
					// it takes a step back and re-write the last written
					// box with a zero. To avoid undoing all the changes
					// once solved, we have to add a check first.
					if s.solved {
						return
					}
					s.board[y][x] = 0
				}
			}
			// Te recursive function returns here when none n is allowed.
			return
		}
	}
	// This point is reached only when all boxes are different than 0.
	s.solved = true
	return
}

func main() {

	mySudoku := newSudoku(hardest)

	args := os.Args[1:]

	for _, arg := range args {
		if arg == "-v" {
			verbose = true
			break
		}
	}

	fmt.Println("Sudoku to be solved:")
	mySudoku.print()

	start := time.Now()
	mySudoku.solve()
	elapsed := time.Since(start)

	fmt.Println("Solution:")
	mySudoku.print()
	fmt.Printf("It took %s and %d iterations to solve the sudoku\n", elapsed, mySudoku.count)
}
