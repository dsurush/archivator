package main

import (
	"testing"
)

/// competitive 3,511720497289684 times faster than consistently Archiver) for 100 files
/// competitive 3,666694080649804 times faster than consistently Archiver) for 200 files

func Benchmark_conc(b *testing.B) {
	strings := []string{"1.exe", "2.exe", "3.exe", "4.exe", "5.exe"}
	for i := 0; i < b.N; i++ {
			handleConc(strings)
	}
}

func Benchmark_seq(b *testing.B) {
	strings := []string{"1.exe", "2.exe", "3.exe", "4.exe", "5.exe"}
	for i := 0; i < b.N; i++ {
		handleSeq(strings)
	}
}