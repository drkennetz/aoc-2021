package main

import (
	_ "embed"
	"fmt"
	"github.com/drkennetz/aoc-2021/utils"
	"log"
	"sort"
	"strings"
)

//go:embed ..\..\inputs\day09\day09.txt

var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		log.Fatalln("Empty input file")
	}
}

func main() {
	parsed := parseInput(input)
	fmt.Println(part1(parsed))
	fmt.Println(part2(parsed))
}

var diffs = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func part1(input [][]int) int {

	total := 0
	for i, row := range input {
		for j, val := range row {
			lowerThanNeighbors := true
			for _, d := range diffs {
				dr, dc := i+d[0], j+d[1]
				if (dr >= 0) && (dr < len(input)) && (dc >= 0) && (dc < len(input[0])) {
					// we are in bounds
					if input[dr][dc] <= val {
						// if higher or even, dr dc not a low point
						lowerThanNeighbors = false
						break
					}
				}
			}
			if lowerThanNeighbors {
				total += 1 + val
			}
		}
	}
	return total
}

func part2(input [][]int) int {

	var lowPoints [][2]int
	for r, rows := range input {
		for c, v := range rows {
			lowerThanAll := true
			for _, d := range diffs {
				dr, dc := r+d[0], c+d[1]
				if dr >= 0 && dr < len(input) && dc >= 0 && dc < len(input[0]) {
					if input[dr][dc] <= v {
						lowerThanAll = false
						break
					}
				}
			}

			if lowerThanAll {
				lowPoints = append(lowPoints, [2]int{r, c})
			}
		}
	}

	// go through all lowpoints and get basin sizes via helper func
	var basins []int
	for _, lp := range lowPoints {
		basins = append(basins, getBasinSize(input, lp[0], lp[1], map[[2]int]bool{}))
	}

	// return 3 largest basins multiplied together
	ans := 1
	sort.Ints(basins)
	for i := 0; i < 3; i++ {
		ans *= basins[len(basins)-1-i]
	}

	return ans
}


func getBasinSize(grid [][]int, r, c int, basinCoords map[[2]int]bool) int {
	// assume that every cell will be involved in one basin, just have to stop at nines
	if grid[r][c] == 9 {
		return 0
	}

	coord := [2]int{r, c}
	// stop if already seen
	if basinCoords[coord] {
		return 0
	}
	// mark as seen
	basinCoords[coord] = true

	for _, d := range diffs {
		dr, dc := r+d[0], c+d[1]
		if dr >= 0 && dr < len(grid) && dc >= 0 && dc < len(grid[0]) {
			// neat little recursion
			getBasinSize(grid, dr, dc, basinCoords)
		}
	}

	// final size of coords map is the basin size
	return len(basinCoords)
}

func parseInput(input string) [][]int {
	lines := strings.Split(input, "\r\n")
	output := make([][]int, 0)
	for _, v := range lines{
		tmp := []int{}
		for _, char := range v {
			tmp = append(tmp, utils.ToInt(string(char)))
		}
		output = append(output, tmp)
	}
	return output
}