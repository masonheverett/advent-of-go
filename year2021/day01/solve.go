package day01

import (
	_ "embed"
	"fmt"
	"masonheverett/advent-of-go/util"
	"strings"
)

//go:embed input.txt
var input string

func Solve() {
	depths := parseInput()
	part1(depths)
	part2(depths)
}

func parseInput() []int {
	depths := make([]int, len(input))
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		depths[i] = util.DecStringToInt(line)
	}
	return depths
}

func part1(depths []int) {
	var count int
	for i, depth := range depths {
		if i == 0 {
			continue
		}
		if depths[i-1] < depth {
			count++
		}
	}
	fmt.Println(count)
}

func part2(depths []int) {
	var count int
	for i, depth := range depths {
		if i < 3 {
			continue
		}
		if depths[i-3] < depth {
			count++
		}
	}
	fmt.Println(count)
}
