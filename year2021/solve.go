package year2021

import (
	"fmt"
	"masonheverett/advent-of-go/year2021/day01"
	"masonheverett/advent-of-go/year2021/day02"
	"masonheverett/advent-of-go/year2021/day03"
	"masonheverett/advent-of-go/year2021/day04"
	"masonheverett/advent-of-go/year2021/day05"
	"masonheverett/advent-of-go/year2021/day06"
	"masonheverett/advent-of-go/year2021/day07"
	"masonheverett/advent-of-go/year2021/day08"
	"masonheverett/advent-of-go/year2021/day09"
	"masonheverett/advent-of-go/year2021/day10"
	"masonheverett/advent-of-go/year2021/day11"
	"masonheverett/advent-of-go/year2021/day12"
	"masonheverett/advent-of-go/year2021/day13"
)

func SolveDay(day int) {
	switch day {
	case 1:
		day01.Solve()
	case 2:
		day02.Solve()
	case 3:
		day03.Solve()
	case 4:
		day04.Solve()
	case 5:
		day05.Solve()
	case 6:
		day06.Solve()
	case 7:
		day07.Solve()
	case 8:
		day08.Solve()
	case 9:
		day09.Solve()
	case 10:
		day10.Solve()
	case 11:
		day11.Solve()
	case 12:
		day12.Solve()
	case 13:
		day13.Solve()
	case 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25:
		fmt.Printf("Day %v not completed yet\n", day)
	}
}
