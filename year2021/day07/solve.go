package day07

import (
	_ "embed"
	"fmt"
	"masonheverett/advent-of-go/util"
	"math"
	"strings"
)

//go:embed input.txt
var input string

func Solve() {
	pos := parseInput()
	part1(pos)
	part2(pos)
}

func parseInput() []int {
	numStrs := strings.Split(input, ",")
	pos := make([]int, len(numStrs))
	for i, str := range numStrs {
		pos[i] = util.DecStringToInt(str)
	}
	return pos
}

func part1(pos []int) {
	max := util.MaxInt(pos)
	min := util.MinInt(pos)
	least := 0
	for i := min; i <= max; i++ {
		cost := 0
		for _, p := range pos {
			cost += int(math.Abs(float64(p - i)))
		}
		if i == min || least > cost {
			least = cost
		}
	}
	fmt.Println(least)
}

func part2(pos []int) {
	max := util.MaxInt(pos)
	min := util.MinInt(pos)
	least := 0
	for i := min; i <= max; i++ {
		cost := 0
		for _, p := range pos {
			n := int(math.Abs(float64(p - i)))
			cost += n * (n + 1) / 2
		}
		if i == min || least > cost {
			least = cost
		}
	}
	fmt.Println(least)
}
