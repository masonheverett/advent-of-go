package day03

import (
	"fmt"
	"log"
	"masonheverett/advent-of-go/util"
	"strconv"
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
	nums := lines
	var oxy, co2 uint64
	var err error
	for i := 0; i < len(nums[0]); i++ {
		nums = trim(nums, i, true)
		if len(nums) == 1 {
			oxy, err = strconv.ParseUint(nums[0], 2, 64)
			if err != nil {
				log.Fatal(err)
			}
			break
		}
	}
	nums = lines
	for i := 0; i < len(nums[0]); i++ {
		nums = trim(nums, i, false)
		if len(nums) == 1 {
			co2, err = strconv.ParseUint(nums[0], 2, 64)
			if err != nil {
				log.Fatal(err)
			}
			break
		}
	}
	fmt.Println(oxy * co2)
}

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
