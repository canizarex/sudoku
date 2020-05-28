package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/canizarex/sudoku-solver/sudoku"
)

func parseCSV(fileName string) [9][9]int {

	// TODO: validate the number of lines & columns.

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var matrix [9][9]int

	r := csv.NewReader(file)

	for i := 0; i < sudoku.Size; i++ {
		record, err := r.Read()

		switch {
		case err == io.EOF:
			msg := " The file must have 9 lines"
			log.Fatalf("Error: %s%s", err, msg)
		case err != nil:
			log.Fatal(err)
		}

		inner := [9]int{}
		for j, str := range record {
			n, _ := strconv.Atoi(str)
			inner[j] = n
		}
		matrix[i] = inner
	}
	return matrix
}
