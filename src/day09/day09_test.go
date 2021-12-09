package main

import "testing"

var benchInput [][]int = parseInput(input)

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1(benchInput)
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2(benchInput)
	}
}