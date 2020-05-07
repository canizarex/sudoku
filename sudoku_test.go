package main

import "testing"


func BenchmarkEasy(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mySudoku := newSudoku(easy)
		mySudoku.solve()
	}
}

func BenchmarkMid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mySudoku := newSudoku(mid)
		mySudoku.solve()
	}
}

func BenchmarkHardest(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mySudoku := newSudoku(hardest)
		mySudoku.solve()
	}

}
