package day06

import (
	_ "embed"
	"fmt"
	"masonheverett/advent-of-go/util"
	"strings"
)

//go:embed input.txt
var input string

type Buckets []int

func (b Buckets) shift() {
	zeroDay := b[0]
	for i := 1; i < len(b); i++ {
		b[i-1] = b[i]
	}
	b[6] += zeroDay
	b[8] = zeroDay
}

func (b Buckets) total() int {
	count := 0
	for _, ct := range b {
		count += ct
	}
	return count
}

func newBuckets(fish []int) Buckets {
	b := make([]int, 9)
	for _, f := range fish {
		b[f]++
	}
	return b
}

func Solve() {
	fish := parseInput()
	part1(fish)
	part2(fish)
}

func parseInput() []int {
	numStrs := strings.Split(input, ",")
	fish := make([]int, len(numStrs))
	for i, str := range numStrs {
		fish[i] = util.DecStringToInt(str)
	}
	return fish
}

func part1(fish []int) {
	b := newBuckets(fish)
	for i := 0; i < 80; i++ {
		b.shift()
	}
	fmt.Println(b.total())
}

func part2(fish []int) {
	b := newBuckets(fish)
	for i := 0; i < 256; i++ {
		b.shift()
	}
	fmt.Println(b.total())
}
