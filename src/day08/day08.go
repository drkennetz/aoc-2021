package main

import (
	_"embed"
	"fmt"
	"log"
	"strings"
)

//..\..\inputs\day08\day08.txt
//go:embed test.txt

var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		log.Fatalln("Empty input file")
	}
}

func main() {
	fmt.Println(parseInput(input))
}

// 7:3, 1:2, 4:4, 8:7
func parseInput (input string) ([][]string, [][]string) {
	lineOut := make([][]string, 0)
	lineIn := make([][]string, 0)
	for _, line := range input {
		var tmpOut string = strings.Split(string(line), " | ")[1]
		var tmpSlOut []string = strings.Split(tmpOut, " ")
		lineOut = append(lineOut, tmpSlOut)
		var tmpIn string = strings.Split(string(line), " | ")[0]
		var tmpSlIn []string = strings.Split(tmpIn, " ")
		lineIn = append(lineIn, tmpSlIn)
	}
	return lineOut, lineIn
}