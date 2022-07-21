package main

import (
	"flag"
	"fmt"
	"masonheverett/advent-of-go/year2021"
)

func main() {
	var year, day int
	flag.IntVar(&year, "y", 2015, "Specify year")
	flag.IntVar(&day, "d", 1, "Specify day")
	flag.Parse()
	fmt.Printf("\n🎄 Advent of Go, Year %v, Day %v 🎄\n\n", year, day)
	switch year {
	case 2015, 2016, 2017, 2018, 2019, 2020:
		fmt.Printf("Year %v not started yet\n", year)
	case 2021:
		year2021.SolveDay(day)
	}
}
