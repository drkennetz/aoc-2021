package main

import (
	_ "embed"
	"fmt"
	"github.com/drkennetz/aoc-2021/utils"
	"log"
	"strings"
)

// ..\..\inputs\day06\day06.txt
//go:embed ..\..\inputs\day06\day06.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		log.Fatalln("Empty input file")
	}
}

func main() {
	parsed := parseInput(input)
	// part 1
	//fmt.Println(part(80, parsed))
	// part 2
	fmt.Println(part(256, parsed))
}

func part(days int, input []int) int64 {
	for d := 0; d < days; d++ {
		save := input[0] // this is where refresh happens
		for i := 0; i < len(input)-1; i++ {
			input[i] = input[i+1] // shift values because all non-zero days lose a day
		}
		input[8] = save  // each 0 adds a new 8
		input[6] += save // each 0 becomes a new 6 so we add to the current 6
	}
	var output int64
	// we add them all up and return!
	for _, v := range input {
		output += int64(v)
	}
	return output
}

func parseInput(input string) []int {
	line := strings.Split(input, ",")
	// all possible vals: 0,1,2,3,4,5,6,7,8
	parsed := make([]int, 9)
	// keep a running sum of all possible values in the input
	for _, v := range line {
		daysRemaining := utils.ToInt(v)
		parsed[daysRemaining]++
	}
	return parsed
}