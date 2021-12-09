package main

import (
	_ "embed"
	"fmt"
	"github.com/drkennetz/aoc-2021/utils"
	"log"
	"strings"
)

//go:embed ..\..\inputs\day04\day04.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		log.Fatalln("empty input file!")
	}
}

func main() {

	ans := part1(input)
	fmt.Println("Output:", ans)
	//ans2 := part2(input)
	//fmt.Println("Output:", ans2)
}

func part1(input string) int {
	nums, boards := parseInput(input)
	for n := 0; n < len(nums); n++ {
		for b := 0; b < len(boards); b++ {
			didWin := boards[b].PickNum(nums[n])
			if didWin {
				return boards[b].Score() * nums[n]
			}
		}
	}
	panic("a board should've won and returned from the loop")
}

func part2(input string) int {
	nums, boards := parseInput(input)
	lastWinningScore := -1
	alreadyWon := map[int]bool{}
	for n := 0; n < len(nums); n++ {
		for b := 0; b < len(boards); b++ {
			if alreadyWon[b] {
				continue
			}
			didWin := boards[b].PickNum(nums[n])
			if didWin {
				lastWinningScore = boards[b].Score() * nums[n]
				alreadyWon[b] = true
			}
		}

	}
	return lastWinningScore
}

// BoardState maintains a parsed board and a boolean matrix of cells that have
// been picked/marked
type BoardState struct {
	board  [5][5]int
	picked [5][5]bool
}

// NewBoardState Creates empty [][]boolean of picked for corresponding board values
func NewBoardState(board [5][5]int) BoardState {
	picked := [5][5]bool{}
	for i := range picked {
		picked[i] = [5]bool{}
	}
	return BoardState{
		board:  board,
		picked: picked,
	}
}

func (b *BoardState) PickNum(num int) bool {
	for r:=0; r < len(b.board); r++ {
		for c:=0; c < len(b.board); c++ {
			if b.board[r][c] == num {
				b.picked[r][c] = true
			}
		}
	}
	for i := 0; i < len(b.board); i++ {
		isFullRow, isFullCol := true, true
		for j := 0; j < len(b.board); j++ {
			// check row at index i
			if !b.picked[i][j] {
				isFullRow = false
			}
			// check col at index j
			if !b.picked[j][i] {
				isFullCol = false
			}
		}
		if isFullRow || isFullCol {
			// returns true if is winning board
			return true
		}
	}

	// false for incomplete board
	return false
}

func (b *BoardState) Score() int {
	var score int

	for r, rows := range b.board {
		for c, v := range rows {
			// adds up all the non-picked/marked cells
			if !b.picked[r][c] {
				score += v
			}
		}
	}

	return score
}

func parseInput(input string) (nums []int, boards []BoardState) {
	lines := strings.Split(input, "\r\n")

	for _, v := range strings.Split(lines[0], ",") {
		nums = append(nums, utils.ToInt(v))
	}
	tmpidx := 0
	b := [5][5]int{}
	for index, grid := range lines[1:] {
		if grid == "" {
			tmpidx = 0
			if index != 0{
				boards = append(boards, NewBoardState(b))
			}
			continue
		}
		grid = strings.ReplaceAll(grid, "  ", " ")
		grid = strings.TrimLeft(grid, " ")
		row := [5]int{}
		for idx, line := range strings.Split(grid, " ") {
			line = strings.TrimLeft(line, " ")
			line = strings.ReplaceAll(line, "  ", " ")
			for line[0] == ' ' {
				line = line[1:]
			}
			parts := strings.Split(line, " ")[0]
			newparts := utils.ToInt(parts)
			row[idx] = newparts
		}

		b[tmpidx] = row
		tmpidx++
	}
	boards = append(boards, NewBoardState(b))
	return nums, boards
}
