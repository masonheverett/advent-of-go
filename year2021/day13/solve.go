package day13

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

type Fold struct {
	axis  string
	value int
}

func Solve() {
	part1(parseInput())
	part2(parseInput())
}

func parseInput() (map[Point]bool, []Fold) {
	sections := strings.Split(input, "\n\n")
	points := make(map[Point]bool)
	for _, line := range strings.Split(sections[0], "\n") {
		pair := strings.Split(line, ",")
		points[Point{
			x: util.DecStringToInt(pair[0]),
			y: util.DecStringToInt(pair[1]),
		}] = true
	}
	foldLines := strings.Split(sections[1], "\n")
	folds := make([]Fold, len(foldLines))
	for i, line := range foldLines {
		words := strings.Split(line, " ")
		pair := strings.Split(words[2], "=")
		folds[i].axis = pair[0]
		folds[i].value = util.DecStringToInt(pair[1])
	}
	return points, folds
}

func part1(points map[Point]bool, folds []Fold) {
	fmt.Println(len(foldPaper(points, folds[0])))
}

func part2(points map[Point]bool, folds []Fold) {
	for _, fold := range folds {
		points = foldPaper(points, fold)
	}
	printPaper(points)
}

func foldPaper(points map[Point]bool, fold Fold) map[Point]bool {
	newPoints := make(map[Point]bool)
	for point := range points {
		switch {
		case fold.axis == "x" && point.x > fold.value:
			newPoints[Point{2*fold.value - point.x, point.y}] = true
		case fold.axis == "y" && point.y > fold.value:
			newPoints[Point{point.x, 2*fold.value - point.y}] = true
		default:
			newPoints[point] = true
		}
	}
	return newPoints
}

func printPaper(points map[Point]bool) {
	maxX, maxY := 0, 0
	for point := range points {
		if point.x > maxX {
			maxX = point.x
		}
		if point.y > maxY {
			maxY = point.y
		}
	}
	grid := make([][]rune, maxY+1)
	for i := range grid {
		grid[i] = make([]rune, maxX+1)
		for j := range grid[i] {
			grid[i][j] = ' '
		}
	}
	for point := range points {
		grid[point.y][point.x] = '#'
	}
	for _, row := range grid {
		fmt.Println(string(row))
	}
}
