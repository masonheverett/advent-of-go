package day09

import (
	_ "embed"
	"fmt"
	"masonheverett/advent-of-go/util"
	"sort"
	"strings"
)

//go:embed input.txt
var input string

type Point struct {
	x, y int
}

func Solve() {
	heightMap := parseInput()
	part1(heightMap)
	part2(heightMap)
}

func parseInput() [][]int {
	lines := strings.Split(input, "\n")
	heightMap := make([][]int, len(lines))
	for i, line := range lines {
		digits := strings.Split(line, "")
		heightMap[i] = make([]int, len(digits))
		for j, digit := range digits {
			heightMap[i][j] = util.DecStringToInt(digit)
		}
	}
	return heightMap
}

func part1(heightMap [][]int) {
	risk := 0
	for i, row := range heightMap {
		for j, val := range row {
			if (i == 0 || val < heightMap[i-1][j]) &&
				(j == 0 || val < row[j-1]) &&
				(i == len(heightMap)-1 || val < heightMap[i+1][j]) &&
				(j == len(row)-1 || val < row[j+1]) {
				risk += val + 1
			}
		}
	}
	fmt.Println(risk)
}

func part2(heightMap [][]int) {
	// Find the low points
	lowPoints := make([]Point, 0)
	for i, row := range heightMap {
		for j, val := range row {
			if (i == 0 || val < heightMap[i-1][j]) &&
				(j == 0 || val < row[j-1]) &&
				(i == len(heightMap)-1 || val < heightMap[i+1][j]) &&
				(j == len(row)-1 || val < row[j+1]) {
				lowPoints = append(lowPoints, Point{i, j})
			}
		}
	}
	basins := make([][]Point, len(lowPoints))
	// Breadth-first search for each basin
	for i, lp := range lowPoints {
		basin := make([]Point, 0)
		vstd := make(map[Point]bool)
		vstd[lp] = true
		q := make([]Point, 0)
		q = append(q, lp)
		for len(q) > 0 {
			pt := q[0]
			q = q[1:]
			basin = append(basin, pt)
			if pt.x > 0 && heightMap[pt.x-1][pt.y] < 9 && !vstd[Point{pt.x - 1, pt.y}] {
				vstd[Point{pt.x - 1, pt.y}] = true
				q = append(q, Point{pt.x - 1, pt.y})
			}
			if pt.y > 0 && heightMap[pt.x][pt.y-1] < 9 && !vstd[Point{pt.x, pt.y - 1}] {
				vstd[Point{pt.x, pt.y - 1}] = true
				q = append(q, Point{pt.x, pt.y - 1})
			}
			if pt.x < len(heightMap)-1 && heightMap[pt.x+1][pt.y] < 9 && !vstd[Point{pt.x + 1, pt.y}] {
				vstd[Point{pt.x + 1, pt.y}] = true
				q = append(q, Point{pt.x + 1, pt.y})
			}
			if pt.y < len(heightMap[pt.x])-1 && heightMap[pt.x][pt.y+1] < 9 && !vstd[Point{pt.x, pt.y + 1}] {
				vstd[Point{pt.x, pt.y + 1}] = true
				q = append(q, Point{pt.x, pt.y + 1})
			}
		}
		basins[i] = basin
	}
	sort.Slice(basins, func(i, j int) bool { return len(basins[i]) > len(basins[j]) })
	fmt.Println(len(basins[0]) * len(basins[1]) * len(basins[2]))
}
