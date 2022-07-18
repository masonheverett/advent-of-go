package day05

import (
	_ "embed"
	"fmt"
	"masonheverett/advent-of-go/util"
	"strings"
)

//go:embed input.txt
var input string

type Point struct {
	x, y int
}

type Segment struct {
	start, end Point
}

type Grid map[Point]int

func (p Point) plot(g Grid) {
	g[p]++
}

func (s Segment) plot(g Grid, includeDiagonal bool) {
	// Just a Point
	if s.start == s.end {
		s.start.plot(g)
		return
	}
	// Vertical Segment
	if s.start.x == s.end.x {
		if s.start.y > s.end.y {
			Segment{s.end, s.start}.plot(g, includeDiagonal)
			return
		}
		for i := s.start.y; i <= s.end.y; i++ {
			Point{s.start.x, i}.plot(g)
		}
		return
	}
	// Horizontal Segment
	if s.start.y == s.end.y {
		if s.start.x > s.end.x {
			Segment{s.end, s.start}.plot(g, includeDiagonal)
			return
		}
		for i := s.start.x; i <= s.end.x; i++ {
			Point{i, s.start.y}.plot(g)
		}
		return
	}
	// Diagonal Segment
	if !includeDiagonal {
		return
	}
	if s.start.x > s.end.x {
		Segment{s.end, s.start}.plot(g, includeDiagonal)
		return
	}
	inc := 1
	if s.start.y > s.end.y {
		inc = -1
	}
	for i, j := s.start.x, s.start.y; i <= s.end.x; i, j = i+1, j+inc {
		Point{i, j}.plot(g)
	}
}

func (g Grid) dangerCount() int {
	count := 0
	for _, c := range g {
		if c > 1 {
			count++
		}
	}
	return count
}

func Solve() {
	segments := parseInput()
	part1(segments)
	part2(segments)
}

func parseInput() []Segment {
	lines := strings.Split(input, "\n")
	segments := make([]Segment, len(lines))
	for i, line := range lines {
		points := strings.Split(line, " -> ")
		start := strings.Split(points[0], ",")
		end := strings.Split(points[1], ",")
		segments[i].start.x = util.DecStringToInt(start[0])
		segments[i].start.y = util.DecStringToInt(start[1])
		segments[i].end.x = util.DecStringToInt(end[0])
		segments[i].end.y = util.DecStringToInt(end[1])
	}
	return segments
}

func part1(segments []Segment) {
	g := Grid(make(map[Point]int))
	for _, seg := range segments {
		seg.plot(g, false)
	}
	fmt.Println(g.dangerCount())
}

func part2(segments []Segment) {
	g := Grid(make(map[Point]int))
	for _, seg := range segments {
		seg.plot(g, true)
	}
	fmt.Println(g.dangerCount())
}
