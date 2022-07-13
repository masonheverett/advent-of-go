package day01

import (
	"fmt"
	"masonheverett/advent-of-go/util"
)

func Solve() {
	depths := util.ReadLinesAsInts("year2021/day01/input.txt")
	util.PrintHeader(2021, 1, 1)
	part1(depths)
	util.PrintHeader(2021, 1, 2)
	part2(depths)
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
