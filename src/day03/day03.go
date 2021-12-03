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
	input, binaries := FileReaderDay3("../../inputs/day03/day03.txt")

	fmt.Println(Part1(input))
	fmt.Println(Part2(binaries))
}

func Part1(m map[int][]string) int64 {
	epsilonSlice, gammaSlice := findMostCommonByPosition(m)
	gamma := strings.Join(gammaSlice, "")
	epsilon := strings.Join(epsilonSlice, "")
	gammaInt, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 64)
	return gammaInt * epsilonInt
}

func Part2(s []string) int64 {

	co2String := filterScrubs("less", s, 0)[0]
	o2String := filterScrubs("more", s, 0)[0]
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
	if comp == "more" {
		if ones >= zeros {
			search = "1"
		} else {
			search = "0"
		}
	} else if comp == "less" {
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

func findMostCommonByPosition(m map[int][]string) ([]string, []string) {
	epsilonSlice := make([]string, len(m))
	gammaSlice := make([]string, len(m))
	for index, digits := range m {
		ones := 0
		zeros := 0
		for _, v := range digits {
			if v == "1" {
				ones++
			} else {
				zeros++
			}
		}
		if ones > zeros {
			epsilonSlice[index] = "0"
			gammaSlice[index] = "1"
		} else if zeros > ones {
			epsilonSlice[index] = "1"
			gammaSlice[index] = "0"
		} else {
			epsilonSlice[index] = "1"
			gammaSlice[index] = "1"
		}
	}
	return epsilonSlice, gammaSlice
}

func FileReaderDay3(s string) (map[int][]string, []string) {
	output := make(map[int][]string)
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
		for i, v := range line {
			output[i] = append(output[i], string(v))
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return output, binaries
}