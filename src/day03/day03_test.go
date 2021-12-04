package main

import "testing"

var binaries = FileReaderDay3("../../inputs/day03/day03.txt")

func BenchmarkPart1(b *testing.B) {
	for n := 0; n <b.N; n++ {
		Part1(binaries)
	}
}

func BenchmarkPart2(b *testing.B) {
	for n :=0; n <b.N; n++ {
		Part2(binaries)
	}
}