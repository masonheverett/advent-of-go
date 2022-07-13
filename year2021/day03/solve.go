package day03

import (
	"fmt"
	"masonheverett/advent-of-go/util"
)

func Solve() {
	lines := util.ReadLinesAsStrings("year2021/day03/input.txt")
	util.PrintHeader(2021, 3, 1)
	part1(lines)
	util.PrintHeader(2021, 3, 2)
	part2(lines)
}

func part1(lines []string) {
	zeroCounts := make([]int, len(lines[0]))
	for _, str := range lines {
		for i, bit := range str {
			if bit == '0' {
				zeroCounts[i]++
			}
		}
	}
	var gamma, epsilon uint
	for _, zc := range zeroCounts {
		gamma <<= 1
		epsilon <<= 1
		if zc < len(lines)/2 {
			gamma++
		} else {
			epsilon++
		}
	}
	fmt.Println(gamma * epsilon)
}

func part2(lines []string) {
	fmt.Println(lines)
}
