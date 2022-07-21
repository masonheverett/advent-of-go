package day14

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Polymer struct {
	pairs       map[string]int
	first, last string
}

func (p *Polymer) step(rules map[string]string) {
	newPairs := make(map[string]int)
	for pair, count := range p.pairs {
		newPairs[string(pair[0])+rules[pair]] += count
		newPairs[rules[pair]+string(pair[1])] += count
	}
	p.pairs = newPairs
}

func (p *Polymer) elemCounts() map[string]int {
	counts := make(map[string]int)
	for pair, count := range p.pairs {
		counts[string(pair[0])] += count
		counts[string(pair[1])] += count
	}
	for elem := range counts {
		counts[elem] /= 2
	}
	counts[p.first] += 1
	counts[p.last] += 1
	return counts
}

func (p *Polymer) elemCountRange() int {
	elemCounts := p.elemCounts()
	max, min := elemCounts[p.first], elemCounts[p.first]
	for _, count := range elemCounts {
		if count > max {
			max = count
		}
		if count < min {
			min = count
		}
	}
	return max - min
}

func Solve() {
	part1(parseInput())
	part2(parseInput())
}

func parseInput() (Polymer, map[string]string) {
	sections := strings.Split(input, "\n\n")
	pairs := make(map[string]int)
	rules := make(map[string]string)
	for i := 0; i < len(sections[0])-1; i++ {
		pairs[sections[0][i:i+2]]++
	}
	for _, line := range strings.Split(sections[1], "\n") {
		pair := strings.Split(line, " -> ")
		rules[pair[0]] = pair[1]
	}
	first, last := sections[0][0], sections[0][len(sections[0])-1]
	return Polymer{pairs, string(first), string(last)}, rules
}

func part1(polymer Polymer, rules map[string]string) {
	for i := 0; i < 10; i++ {
		polymer.step(rules)
	}
	fmt.Println(polymer.elemCountRange())
}

func part2(polymer Polymer, rules map[string]string) {
	for i := 0; i < 40; i++ {
		polymer.step(rules)
	}
	fmt.Println(polymer.elemCountRange())
}
