package day04

import (
	_ "embed"
	"fmt"
	"masonheverett/advent-of-go/util"
	"strings"
)

//go:embed input.txt
var input string

type Board struct {
	board  [][]int
	marked [][]bool
}

func (b *Board) mark(num int) bool {
	for r, row := range b.board {
		for c, val := range row {
			if val == num {
				b.marked[r][c] = true
			}
		}
	}
	return b.won()
}

func (b *Board) won() bool {
	for r := 0; r < len(b.marked); r++ {
		fullRow, fullCol := true, true
		for c := 0; c < len(b.marked[r]); c++ {
			if !b.marked[r][c] {
				fullRow = false
			}
			if !b.marked[c][r] {
				fullCol = false
			}
		}
		if fullRow || fullCol {
			return true
		}
	}
	return false
}

func (b *Board) score() int {
	score := 0
	for r, row := range b.board {
		for c, val := range row {
			if !b.marked[r][c] {
				score += val
			}
		}
	}
	return score
}

func Solve() {
	nums, boards := parseInput()
	part1(nums, boards)
	part2(nums, boards)
}

func parseInput() ([]int, []Board) {
	sections := strings.Split(input, "\n\n")
	markStrs := strings.Split(sections[0], ",")
	nums := make([]int, len(markStrs))
	for i, str := range markStrs {
		nums[i] = util.DecStringToInt(str)
	}
	boards := make([]Board, len(sections)-1)
	for i, sec := range sections[1:] {
		board := [][]int{}
		for _, line := range strings.Split(sec, "\n") {
			line = strings.ReplaceAll(line, "  ", " ")
			for line[0] == ' ' {
				line = line[1:]
			}
			numStrs := strings.Split(line, " ")
			row := []int{}
			for _, numStr := range numStrs {
				row = append(row, util.DecStringToInt(numStr))
			}
			board = append(board, row)
		}
		marked := make([][]bool, len(board))
		for j := range marked {
			marked[j] = make([]bool, len(board[0]))
		}
		boards[i].board = board
		boards[i].marked = marked
	}
	return nums, boards
}

func part1(nums []int, boards []Board) {
	for _, num := range nums {
		for _, board := range boards {
			if board.mark(num) {
				fmt.Println(board.score() * num)
				return
			}
		}
	}
}

func part2(nums []int, boards []Board) {
	winners := make([]bool, len(boards))
	winCount := 0
	for _, num := range nums {
		for i, board := range boards {
			if !winners[i] {
				if board.mark(num) {
					winners[i] = true
					winCount++
					if winCount == len(boards) {
						fmt.Println(board.score() * num)
						return
					}
				}
			}
		}
	}
}
