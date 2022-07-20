package day10

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed input.txt
var input string

type Brace rune

func (b Brace) isOpen() bool {
	return b == '(' || b == '[' || b == '{' || b == '<'
}

func (b Brace) closingPartner() Brace {
	return Brace(map[Brace]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}[b])
}

func (b Brace) score() int {
	return map[Brace]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}[b]
}

type Stack []Brace

func (s Stack) push(b Brace) Stack {
	return append(s, b)
}

func (s Stack) pop() (Brace, Stack) {
	return s[len(s)-1], s[:len(s)-1]
}

func (s Stack) score() int {
	bScore := map[Brace]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}
	total := 0
	for i := len(s) - 1; i >= 0; i-- {
		total *= 5
		total += bScore[s[i]]
	}
	return total
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
	total := 0
	for _, line := range lines {
		stack := Stack(make([]Brace, 0))
		for _, r := range line {
			br := Brace(r)
			if br.isOpen() {
				stack = stack.push(br)
				continue
			}
			var open Brace
			open, stack = stack.pop()
			if open.closingPartner() != br {
				total += br.score()
				break
			}
		}
	}
	fmt.Println(total)
}

func part2(lines []string) {
	scores := make([]int, 0)
	for _, line := range lines {
		stack := Stack(make([]Brace, 0))
		isCorrupt := false
		for _, r := range line {
			br := Brace(r)
			if br.isOpen() {
				stack = stack.push(br)
				continue
			}
			var open Brace
			open, stack = stack.pop()
			if open.closingPartner() != br {
				isCorrupt = true
				break
			}
		}
		if !isCorrupt {
			scores = append(scores, stack.score())
		}
	}
	sort.Ints(scores)
	fmt.Println(scores[(len(scores)-1)/2])
}
