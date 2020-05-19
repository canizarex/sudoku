package sudoku

import (
	"fmt"
	"strings"
	"time"
)

const (
	Size    = 9
	BoxSize = 3
)

var clearScr = fmt.Sprint("\033[2J \033[H")

type Sudoku struct {
	Grid    [9][9]int
	Count   int
	Solved  bool
	Verbose bool
}

func New(matrix [9][9]int) *Sudoku {
	s := new(Sudoku)
	s.Grid = matrix
	return s
}

func (s *Sudoku) Draw() string {

	screen := make([]byte, 0, 605)

	var (
		n1     int      = 2*(Size+BoxSize) - 1
		n2     int      = 2*BoxSize + 1
		vSep   string   = "║"
		hSep   string   = "═"
		xSep   string   = "O"
		corner []string = []string{"╔", "╗", "╚", "╝"}
	)

	top := fmt.Sprintf("%s%s%s\n", corner[0], strings.Repeat(hSep, n1), corner[1])
	bottom := fmt.Sprintf("%s%s%s", corner[2], strings.Repeat(hSep, n1), corner[3])
	middle := fmt.Sprintf("%[1]s%[2]s%[3]s%[2]s%[3]s%[2]s%[1]s", vSep, strings.Repeat(hSep, n2), xSep)

	screen = append(screen, top...)
	for i, row := range s.Grid {
		for i, n := range row {
			switch {
			case i == 0:
				screen = append(screen, fmt.Sprintf("%s%2d", vSep, n)...)
			case (i+1)%BoxSize == 0:
				screen = append(screen, fmt.Sprintf("%2d%2s", n, vSep)...)
			default:
				screen = append(screen, fmt.Sprintf("%2d", n)...)
			}
		}
		switch {
		case i+1 == Size:
			screen = append(screen, fmt.Sprintf("\n%s\n", bottom)...)
		case (i+1)%BoxSize == 0:
			screen = append(screen, fmt.Sprintf("\n%s\n", middle)...)
		default:
			screen = append(screen, '\n')
		}
	}
	return string(screen)
}

func (s *Sudoku) possible(y, x, n int) bool {

	// Check all the numbers in a given row
	for i := 0; i < Size; i++ {
		if s.Grid[y][i] == n {
			return false
		}
	}

	// Check all the numbers in a given column
	for i := 0; i < Size; i++ {
		if s.Grid[i][x] == n {
			return false
		}
	}

	// Check a 3:3 box
	x0 := (x / BoxSize) * BoxSize
	y0 := (y / BoxSize) * BoxSize

	for i := 0; i < BoxSize; i++ {
		for j := 0; j < BoxSize; j++ {
			if n == s.Grid[y0+i][x0+j] {
				return false
			}
		}
	}
	return true
}

func (s *Sudoku) Solve() {

	s.Count++

	if s.Verbose {
		fmt.Printf("%s%s\n%d\n", clearScr, s.Draw(), s.Count)
		time.Sleep(10 * time.Millisecond)
	}

	for y := 0; y < Size; y++ {
		for x := 0; x < Size; x++ {
			// Go ahead only if the cell is empty (equals zero)
			if s.Grid[y][x] != 0 {
				continue
			}
			// For every n check if it is possible in a given cell
			// and if it is, call the function recursively to
			// start again.
			for n := 1; n <= Size; n++ {
				if s.possible(y, x, n) {
					s.Grid[y][x] = n
					s.Solve()
					// At this point the recursive function has returned
					// because there are no more possibilities so
					// it takes a step back and re-write the last written
					// cell with a zero. To avoid undoing all the changes
					// once solved, it's necessary to check if it's already
					// solved.
					if s.Solved {
						return
					}
					s.Grid[y][x] = 0
				}
			}
			// Te recursive function returns here when none n is allowed.
			return
		}
	}
	// This point is reached only when all the cells are different than 0.
	s.Solved = true
}
