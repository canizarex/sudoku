package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/canizarex/sudoku-solver/sudoku"
)

var (
	verbose = flag.Bool("v", false, "Verbose output")
	sample  = flag.String("s", "easy", "Sample to solve: [easy | mid | hardest]")
	fName   = flag.String("f", "", "CSV file containing a sudoku")
)

func main() {

	flag.Parse()

	var mySudoku *sudoku.Sudoku

	if *fName == "" {
		mySudoku = sudoku.New(samples[*sample])
	} else {
		mySudoku = sudoku.New(parseCSV(*fName))
	}

	if *verbose {
		mySudoku.Verbose = true
	}

	fmt.Println("Sudoku to be solved:")
	fmt.Print(mySudoku.Draw())

	start := time.Now()
	mySudoku.Solve()
	elapsed := time.Since(start)

	if !mySudoku.Solved {
		fmt.Println("The sudoku couldn't be solved")
		return
	}

	fmt.Println("Solution:")
	fmt.Print(mySudoku.Draw())
	fmt.Printf("It took %s and %d iterations to solve the sudoku\n", elapsed, mySudoku.Count)
}
