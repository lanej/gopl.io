package main

import (
	"testing"
)

func BenchmarkManualJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		manualJoin()
	}
}

func BenchmarkLibraryJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		libraryJoin()
	}
}
