package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/canizarex/sudoku-solver/sudoku"
)

var (
	instructions = "Please introduce the sudoku to be solved using %q as separator:\n\n"

	delimiter = flag.String("d", " ", "delimiter used to parse the input")
	fileName  = flag.String("f", "", "file containing a sudoku")
	fps       = flag.Uint("r", 120, "refresh rate for the verbose mode")
	sample    = flag.String("s", "", "sample to solve: [easy | mid | hardest]")
	verbose   = flag.Bool("v", false, "verbose output")
)

func parseInput(r io.Reader) [9][9]int {
	
	matrix := [9][9]int{}
	scanner := bufio.NewScanner(r)

	for i := 0; i < sudoku.Size; i++ {
		scanner.Scan()
		row := strings.Split(strings.TrimSpace(scanner.Text()), *delimiter)

		if len(row) != sudoku.Size {
			log.Fatal(fmt.Errorf("Error: The sudoku must have %d columns and %[1]d rows", sudoku.Size))
		}

		inner := [9]int{}
		for j, str := range row {
			n, _ := strconv.Atoi(str)
			inner[j] = n
		}
		matrix[i] = inner
	}
	return matrix
}

func main() {

	flag.Parse()

	var mySudoku *sudoku.Sudoku

	switch {
	case *sample != "":
		matrix, ok := samples[*sample]
		if ! ok {
			log.Fatalf("Couldn't found a sample called %q", *sample)
		}
		mySudoku = sudoku.New(*matrix)
	case *fileName != "":
		file, err := os.Open(*fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		fmt.Printf("Parsing file using %q as delimiter...\n", *delimiter)
		mySudoku = sudoku.New(parseInput(file))
	default:
		fmt.Printf(instructions, *delimiter)
		mySudoku = sudoku.New(parseInput(os.Stdin))
	}

	if *verbose {
		mySudoku.Verbose = true
		mySudoku.Fps = *fps
	}

	fmt.Printf("\nSudoku to be solved:\n")
	fmt.Print(mySudoku.Draw())

	start := time.Now()
	mySudoku.Solve()
	elapsed := time.Since(start)

	if !mySudoku.Solved {
		fmt.Println("The sudoku couldn't be solved")
		return
	}

	fmt.Printf("\nSolution:\n")
	fmt.Print(mySudoku.Draw())
	fmt.Printf("It took %s and %d iterations to solve the sudoku\n", elapsed, mySudoku.Count)
}
