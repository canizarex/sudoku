package main

import (
	"encoding/csv"
	"log"
	"os"
	"io"
	"strconv"
	_ "fmt"
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

	for i := 0; i < size; i++ {
		record, err := r.Read()

		switch {
		case err == io.EOF:
			msg := " The file must have 9 lines"
			log.Fatalf("Error: %s%s", err, msg)
		case err != nil:
			log.Fatal(err)
		}

		inner := [9]int{}
		for i, str := range record {
			n, _ := strconv.Atoi(str)
			inner[i] = n
		}
		matrix[i] = inner
	}
	return matrix
}
