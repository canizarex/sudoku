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

var (
		// Chars to clear the screen and move the cursor in *nix
		clearScr = fmt.Sprint("\033[2J \033[H")

		// Related to draw():
		// Multiplier to draw the top and bottom bar
		m1 int = 2*(Size+BoxSize) - 1
		// Multiplier to draw the middle bar
		m2 int = 2*BoxSize + 1

		// Different sepatators to represent a sodoku
		vSep   string   = "║"
		vSep2  string   = "|"
		hSep   string   = "═"
		hSep2  string   = "-"
		xSep   string   = "+"
		corner []string = []string{"╔", "╗", "╚", "╝"}
)

type Sudoku struct {
	Grid            [9][9]int
	row, col, Count int
	Fps             uint
	Solved, Verbose bool
}

func New(matrix [9][9]int) *Sudoku {
	return &Sudoku{Grid: matrix}
}

func (s *Sudoku) Draw() string {

	screenCap := 605 // Max number of bytes needed to draw the sudoku
	screen := make([]byte, 0, screenCap)

	top := fmt.Sprintf("%s%s%s\n", corner[0], strings.Repeat(hSep, m1), corner[1])
	bottom := fmt.Sprintf("%s%s%s", corner[2], strings.Repeat(hSep, m1), corner[3])
	middle := fmt.Sprintf("%[1]s%[2]s%[3]s%[2]s%[3]s%[2]s%[1]s", vSep, strings.Repeat(hSep2, m2), xSep)

	screen = append(screen, top...)
	for i, row := range s.Grid {
		for j, n := range row {
			switch {
			case j == 0:
				screen = append(screen, fmt.Sprintf("%s%2d", vSep, n)...)
			case (j + 1) == Size:
				screen = append(screen, fmt.Sprintf("%2d%2s", n, vSep)...)
			case (j+1)%BoxSize == 0:
				screen = append(screen, fmt.Sprintf("%2d%2s", n, vSep2)...)
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

func (s *Sudoku) isPossible(row, col, n int) bool {

	// Check all the numbers in a given row
	for i := 0; i < Size; i++ {
		if s.Grid[row][i] == n {
			return false
		}
	}

	// Check all the numbers in a given column
	for i := 0; i < Size; i++ {
		if s.Grid[i][col] == n {
			return false
		}
	}

	// Check a 3:3 box
	c0 := (col / BoxSize) * BoxSize
	r0 := (row / BoxSize) * BoxSize

	for i := 0; i < BoxSize; i++ {
		for j := 0; j < BoxSize; j++ {
			if n == s.Grid[r0+i][c0+j] {
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
		time.Sleep(time.Second / time.Duration(s.Fps))
	}

	for row := s.row; row < Size; row++ {
		for col := s.col; col < Size; col++ {
			// Check if the cell is empty (equals zero)
			if s.Grid[row][col] != 0 {
				continue
			}
			// For every n check if it is possible in a given cell
			// and if it is, mark the position and call the function
			// recursively to start again.
			for n := 1; n <= Size; n++ {
				if s.isPossible(row, col, n) {
					s.Grid[row][col] = n
					s.row, s.col = row, col
					s.Solve()
				}
			}
			if s.Solved {
				return
			}
			s.Grid[row][col] = 0
			// The recursive function returns here when none n is allowed.
			return
		}
		// Reset the position of s.col before going to the next s.row
		s.col = 0
	}
	// This point is reached only when all the cells are different than 0.
	s.Solved = true
}
