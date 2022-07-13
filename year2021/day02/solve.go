package day02

import (
	"fmt"
	"log"
	"masonheverett/advent-of-go/util"
	"strconv"
)

type command struct {
	dir  string
	dist int
}

func Solve() {
	slices := util.ReadLinesAsStringSlices("year2021/day02/input.txt")
	commands := parseCommands(slices)
	util.PrintHeader(2021, 2, 1)
	part1(commands)
	util.PrintHeader(2021, 2, 2)
	part2(commands)
}

func parseCommands(slices [][]string) []command {
	commands := make([]command, len(slices))
	for i, slice := range slices {
		dist, err := strconv.Atoi(slice[1])
		if err != nil {
			log.Fatal(err)
		}
		commands[i] = command{slice[0], dist}
	}
	return commands
}

func part1(commands []command) {
	var hrz, vrt int
	for _, cmd := range commands {
		switch cmd.dir {
		case "forward":
			hrz += cmd.dist
		case "up":
			vrt -= cmd.dist
		case "down":
			vrt += cmd.dist
		}
	}
	fmt.Println(hrz * vrt)
}

func part2(commands []command) {
	var hrz, vrt, aim int
	for _, cmd := range commands {
		switch cmd.dir {
		case "forward":
			hrz += cmd.dist
			vrt += aim * cmd.dist
		case "up":
			aim -= cmd.dist
		case "down":
			aim += cmd.dist
		}
	}
	fmt.Println(hrz * vrt)
}
