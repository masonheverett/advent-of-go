package day03

import (
	_ "embed"
	"fmt"
	"masonheverett/advent-of-go/util"
	"strings"
)

//go:embed input.txt
var input string

func trim(nums []string, ndx int, most bool) []string {
	zeros, ones := make([]string, 0), make([]string, 0)
	for _, num := range nums {
		if num[ndx] == '0' {
			zeros = append(zeros, num)
		} else {
			ones = append(ones, num)
		}
	}
	if most {
		if len(ones) >= len(zeros) {
			return ones
		}
		return zeros
	}
	if len(zeros) <= len(ones) {
		return zeros
	}
	return ones
}

func Solve() {
	lines := parseInput()
	part1(lines)
	part2(lines)
}

func parseInput() []string {
	return strings.Split(input, "\n")
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
	nums := lines
	var oxy, co2 int
	for i := 0; i < len(nums[0]); i++ {
		nums = trim(nums, i, true)
		if len(nums) == 1 {
			oxy = util.BinStringToInt(nums[0])
			break
		}
	}
	nums = lines
	for i := 0; i < len(nums[0]); i++ {
		nums = trim(nums, i, false)
		if len(nums) == 1 {
			co2 = util.BinStringToInt(nums[0])
			break
		}
	}
	fmt.Println(oxy * co2)
}
