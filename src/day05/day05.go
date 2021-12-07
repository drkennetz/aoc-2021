package main

import (
	_"embed"
	"fmt"
	"github.com/drkennetz/aoc-2021/utils"
	"log"
	"strings"
)
// ..\..\inputs\day05\day05.txt
//go:embed ..\..\inputs\day05\day05.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		log.Fatalln("empty input file!")
	}
}

func main() {
	coords, grid := parseInput(input)
	//fmt.Println(part(1, coords, grid))
	fmt.Println(part(2, coords, grid))
}

func part(part int, coords [][4]int, grid [][]int) int {
	for _, c := range coords {
		// horizontal case
		if c[0] == c[2] {
			row := c[0]
			start, end := c[1], c[3]
			if c[1] > c[3] {
				start, end = end, start
			}
			for col := start; col <= end; col++ {
				grid[row][col]++
			}
			// vertical case
		} else if c[1] == c[3] {
			col := c[1]
			start, end := c[0], c[2]
			if c[0] > c[2] {
				start, end = end, start
			}
			for row := start; row <= end; row++ {
				grid[row][col]++
			}
		} else if part == 2 {
			// we can check on diags
			// if y1 > y2 we reverse the pairs
			if c[1] > c[3] {
				c = [4]int{c[2], c[3], c[0], c[1]}
			}
			// compare rows, will be going right because of the above ^
			// if going down and right (down has increasing indexes which is confusing)
			if c[0] < c[2] {
				// fill in missing points
				for row := c[0]; row <= c[2]; row++ {
					col := c[1] + row - c[0]
					grid[row][col]++
				}
			} else {
				// going up and right (up has decreasing indexes which is confusing)
				for row := c[0]; row >= c[2]; row-- {
					col := c[1] + c[0] - row
					grid[row][col]++
				}
			}
		}
	}

	// count up gt 2
	var output int
	for _, rows := range grid {
		for _, v := range rows {
			if v >= 2 {
				output++
			}
		}
	}
	return output
}


// parse the coordinates and create a grid
func parseInput(input string) ([][4]int, [][]int) {
	var coord [][4]int
	var maxRow, maxCol int
	for _, line := range strings.Split(input, "\n") {
		var x1, y1, x2, y2 int
		_, err := fmt.Sscanf(line, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		if err != nil {
			log.Fatalln("Error parsing line")
		}
		// compiler should sort this one out for me :D
		if utils.Max(x1, x2) > maxRow {
			maxRow = utils.Max(x1, x2)
		}
		if utils.Max(y1, y2) > maxCol {
			maxCol = utils.Max(y1, y2)
		}
		coord = append(coord, [4]int{x1, y1, x2, y2})
	}
	grid := make([][]int, maxRow+1)
	for i := range grid {
		grid[i] = make([]int, maxCol+1)
	}
	return coord, grid
}