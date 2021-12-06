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
	binaries := FileReaderDay3("../../inputs/day03/day03.txt")
	fmt.Println(Part1(binaries))
	fmt.Println(Part2(binaries))
}

func Part1(binaries []string) int64 {
	epsilon, gamma := findMostCommonByPosition(binaries)
	gammaInt, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 64)
	return gammaInt * epsilonInt
}

func Part2(s []string) int64 {
	co2String := filterScrubs("co2", s, 0)[0]
	o2String := filterScrubs("o2", s, 0)[0]
	co2Int, _ := strconv.ParseInt(co2String, 2, 64)
	o2Int, _ := strconv.ParseInt(o2String, 2, 64)
	return co2Int * o2Int
}

func filterScrubs(comp string, b []string, index int) []string {
	if len(b) == 1 {
		return b
	}
	ones := 0
	zeros := 0
	var tmpSlice []string
	for _, v := range b {
		if string(v[index]) == "1" {
			ones++
		} else {
			zeros++
		}
	}
	search := ""
	if comp == "o2" {
		if ones >= zeros {
			search = "1"
		} else {
			search = "0"
		}
	} else if comp == "co2" {
		if ones < zeros {
			search = "1"
		} else {
			search = "0"
		}
	}
	for _, binary := range b {
		if string(binary[index]) == search {
			tmpSlice = append(tmpSlice, binary)
		}
	}
	return filterScrubs(comp, tmpSlice, index+1)
}

// gamma is the most common bit across each index
// epsilon is the least common
func findMostCommonByPosition(binaries []string) (string, string) {
	var gamma, epsilon string
	for i:=0; i < len(binaries[0]); i++ {
		var zeroes, ones int
		for _, binary := range binaries {
			if binary[i] == '0' {
				zeroes++
			} else {
				ones++
			}
		}
		if zeroes > ones {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	return epsilon, gamma
}

func FileReaderDay3(s string) []string {
	var binaries []string
	file, err := os.Open(s)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		binaries = append(binaries, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return binaries
}