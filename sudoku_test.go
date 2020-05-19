package main

import (
	"testing"

	"github.com/canizarex/go-exercises/sudoku-solver-go/sudoku"
)


func BenchmarkEasy(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mySudoku := sudoku.New(easy)
		mySudoku.Solve()
	}
}

func BenchmarkMid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mySudoku := sudoku.New(mid)
		mySudoku.Solve()
	}
}

func BenchmarkHardest(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mySudoku := sudoku.New(hardest)
		mySudoku.Solve()
	}

}
