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
	inputS, inputI := FileReaderDay2("../../inputs/day02/day02.txt")
	fmt.Println(Part1(inputS, inputI))
	fmt.Println(Part2(inputS, inputI))
}

func Part1(strings []string, nums []int) int {
	vertical := 0
	horizontal := 0
	for i := 0; i < len(strings); i++ {
		if strings[i] == "forward" {
			horizontal += nums[i]
		} else if strings[i] == "down" {
			vertical += nums[i]
		} else {
			vertical -= nums[i]
		}
	}
	return vertical*horizontal
}

func Part2(strings []string, nums []int) int {
	vertical := 0
	horizontal := 0
	aim := 0
	for i := 0; i < len(strings); i++ {
		if strings[i] == "forward" {
			horizontal += nums[i]
			vertical += nums[i]*aim
		} else if strings[i] == "down" {
			aim += nums[i]
		} else {
			aim -= nums[i]
		}
	}
	return vertical*horizontal
}

func FileReaderDay2(s string) ([]string, []int) {
	outputS := make([]string, 0)
	outputI := make([]int, 0)
	file, err := os.Open(s)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		outputS = append(outputS, line[0])
		numI, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatalln(err)
		}
		outputI = append(outputI, numI)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return outputS, outputI
}
