package main

import (
	"log"
	"testing"
)

var S, I = FileReaderDay2("../../inputs/day02/day02.txt")

func TestPart1(t *testing.T) {
	result1 := 2150351
	actual1 := Part1(S, I)
	if actual1 != result1 {
		log.Fatalln("FAIL! expected 2150351 got ", actual1)
	}
}

func TestPart2(t *testing.T) {
	result2 := 1842742223
	actual2 := Part2(S, I)
	if actual2 != result2 {
		log.Fatalln("FAIL! expected 1842742223 got ", actual2)
	}
}

func BenchmarkPart1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Part1(S, I)
	}
}

func BenchmarkPart2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Part2(S, I)
	}
}