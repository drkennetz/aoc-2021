package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := FileReaderDay1("../../inputs/day01/day01.txt")
	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
	fmt.Println(Part2Refactor(input))
}

// Part1 checks to see if next value is greater than current value and returns the count of those increases
func Part1(intslice []int) int {
	current := 0
	next := 1
	increasing := 0
	for current < (len(intslice) -1) {
		if intslice[current] < intslice[next] {
			increasing++
		}
		current++
		next++
	}
	return increasing
}

func Part1Refactor(intslice []int) int {
	increasing := 0
	for i:=1; i < len(intslice); i++ {
		if intslice[i] > intslice[i-1] {increasing++}
	}
	return increasing
}

// Part2 checks to see if sliding windows of 3 are less than next sliding windows of 3
func Part2(intslice []int) int {
	currentStart, currentStop := 0, 2
	nextStart, nextStop := 1, 3
	increasing := 0
	for currentStop < (len(intslice) - 1) {
		if miniSum(intslice[currentStart:currentStop+1]) < miniSum(intslice[nextStart:nextStop+1]) {
			increasing++
		}
		currentStart++; currentStop++; nextStart++; nextStop++
	}
	return increasing
}

func Part2Refactor(intslice []int) int{
	increasing := 0
	for i := 3; i < len(intslice); i++ {
		curr, prev := intslice[i], intslice[i-3]
		if curr > prev {increasing++}
	}
	return increasing
}
// helper func
func miniSum(intslice []int) int {
	res := 0
	for _, v := range intslice {
		res += v
	}
	return res
}

// FileReaderDay1 reads the Day1 input data into a []int
func FileReaderDay1(s string) []int {
	parsed := make([]int, 0)

	file, err := os.Open(s)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			log.Fatalln(err)
		}
		parsed = append(parsed, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return parsed
}
