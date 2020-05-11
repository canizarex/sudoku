package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const size = 9
const subGridSize = size / 3

var verbose bool

type sudoku struct {
	board  [9][9]int
	count  int
	solved bool
}

func newSudoku(matrix [9][9]int) sudoku {
	return sudoku{matrix, 0, false}
}

func clear() {
	fmt.Print("\033[2J \033[H")
}

func (s *sudoku) print() {
	
	var (
		hsepMult = size*2+subGridSize*2+1-2
		p string = "║" 
		hsep string = "═" 
		xsep string = "O"
		corner = []string{"╔", "╗", "╚", "╝"}
	)
	
	top := fmt.Sprintf("%s%s%s", corner[0], strings.Repeat(hsep, hsepMult), corner[1])
	bottom := fmt.Sprintf("%s%s%s", corner[2], strings.Repeat(hsep, hsepMult), corner[3])
	middle := fmt.Sprintf("%[1]s%[2]s%[3]s%[2]s%[3]s%[2]s%[1]s", p, strings.Repeat(hsep, 7), xsep)

	fmt.Printf("%s\n", top)
	for i, row := range s.board {
		for i, n := range row {
			switch {
			case i == 0:
				fmt.Printf("%s%2d", p, n)
			case (i+1)%subGridSize == 0:
				fmt.Printf("%2d%2s", n, p)
			default:
				fmt.Printf("%2d", n)
			}
		}
		switch {
		case i+1 == size:
			fmt.Printf("\n%s\n", bottom)
		case (i+1)%subGridSize == 0:
			fmt.Printf("\n%s\n", middle)
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
	x0 := (x / subGridSize) * subGridSize
	y0 := (y / subGridSize) * subGridSize

	for i := 0; i < subGridSize; i++ {
		for j := 0; j < subGridSize; j++ {
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
		clear()
		s.print()
		fmt.Println(s.count)
		time.Sleep(25 * time.Millisecond)
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
			for n := 1; n <= size; n++ {
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
}

func main() {

	mySudoku := newSudoku(mid)

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
