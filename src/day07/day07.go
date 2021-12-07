package main

import (
	_ "embed"
	"fmt"
	"github.com/drkennetz/aoc-2021/utils"
	"log"
	"math"
	"sort"
	"strings"
)

// switch these two to switch between dummy case and real case
// ..\..\inputs\day07\day07.txt
//go:embed ..\..\inputs\day07\day07.txt

var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		log.Fatalln("Empty input file")
	}
}

func main() {
	input := parseInput(input)
	fmt.Println(calcFuelAtStep(5, 16))
	fmt.Println(findFuelPart1(input))
	fmt.Println(findFuelPart2(input))
}

// https://www.google.com/search?q=exponential+increase+formula&rlz=1C1GCEA_enUS845US845&oq=exponential+increase&aqs=chrome.0.0i512j69i57j0i512l8.3136j0j7&sourceid=chrome&ie=UTF-8
func findFuelPart2(input []int) int {
	min := math.MaxInt32
	for i := 0; i < len(input); i++ {
		cost := 0
		for _, v := range input {
			cost +=  calcFuelAtStep(v, i)
		}
		if cost < min {
			min = cost
		}
	}
	return min
}

func calcFuelAtStep(i, j int) int {
	lossTotal := 0
	current := utils.Abs(i, j)
	for i := 0; i < current; i++ {
		lossTotal += i+1
	}
	return lossTotal
}

// The least amount of fuel will be found by normalization using median
func findFuelPart1(input []int) int {
	middleIdx := len(input)/2
	middleNum := input[middleIdx]
	total := 0
	for _, v := range input {
		total += utils.Abs(middleNum, v)
	}
	return total
}

func parseInput(input string) []int {
	parsed := strings.Split(input, ",")
	output := make([]int, 0)
	for _, v := range parsed {
		output = append(output, utils.ToInt(v))
	}
	// we sort it because the distance from the middle will be the shortest distance
	sort.Ints(output)
	return output
}