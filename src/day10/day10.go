package main

import (
	_ "embed"
	"github.com/drkennetz/aoc-2021/utils"
	"fmt"
	"log"
	"sort"
	"strings"
)

//..\..\inputs\day10\day10.txt
//go:embed ..\..\inputs\day10\day10.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		log.Fatalln("empty input file!")
	}
}

var scores = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var p2scores = map[string]int {
	"(": 1,
	"[": 2,
	"{": 3,
	"<": 4,
}

var closures = map[string]string {
	")": "(",
	"]": "[",
	"}": "{",
	">": "<",
}

func main() {
	parsed := parseInput(input)
	fmt.Println(balancedBrackets(1, parsed))
	fmt.Println(balancedBrackets(2, parsed))
}

func balancedBrackets(part int, input [][]string) int {
	part1Score := 0
	part2Scores := make([]int, 0)

	for _, s := range input {
		stack := make([]string, 0)
		CorruptLine := false
		for i := 0; i < len(s); i++ {
			//tmp := make([]string, 0)
			closed := false
			if _, ok := closures[string(s[i])]; ok {
				closed = true
			}
			if !closed {
				stack = append(stack, string(s[i]))
			} else {
				if len(stack) == 0 {
					panic("empty stack")
				}
				closure := utils.PopString(&stack)
				if closures[string(s[i])] != closure {
					part1Score += scores[string(s[i])]
					CorruptLine = true
				}
			}
		}

		if !CorruptLine {
			score := 0
			for j := len(stack) - 1; j >= 0; j-- {
				score *= 5
				score += p2scores[stack[j]]
			}
			part2Scores = append(part2Scores, score)
		}
	}
	if part == 1 {
		return part1Score
	}
	sort.Ints(part2Scores)
	return part2Scores[len(part2Scores)/2]
}

func parseInput(s string) [][]string {
	lines := strings.Split(s, "\r\n")
	output := make([][]string, 0)
	for _, v := range lines {
		tmp := []string{}
		for _, char := range v {
			tmp = append(tmp, string(char))
		}
		output = append(output, tmp)
	}
	return output
}