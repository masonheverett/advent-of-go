package day02

import (
	_ "embed"
	"fmt"
	"masonheverett/advent-of-go/util"
	"strings"
)

//go:embed input.txt
var input string

type Command struct {
	dir  string
	dist int
}

func Solve() {
	commands := parseInput()
	part1(commands)
	part2(commands)
}

func parseInput() []Command {
	lines := strings.Split(input, "\n")
	commands := make([]Command, len(lines))
	for i, line := range lines {
		pair := strings.Split(line, " ")
		commands[i] = Command{pair[0], util.DecStringToInt(pair[1])}
	}
	return commands
}

func part1(commands []Command) {
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

func part2(commands []Command) {
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
