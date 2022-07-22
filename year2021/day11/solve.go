package day11

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

func Solve() {
	part1(parseInput())
	part2(parseInput())
}

func parseInput() [][]int {
	lines := strings.Split(input, "\n")
	grid := make([][]int, len(lines))
	for i, line := range lines {
		grid[i] = make([]int, len(line))
		for j, n := range strings.Split(line, "") {
			grid[i][j] = util.DecStringToInt(n)
		}
	}
	return grid
}

func part1(grid [][]int) {
	flashCount := 0
	for i := 0; i < 100; i++ {
		// Load up a queue of increases to make
		q := make([]Point, 0)
		for x, row := range grid {
			for y := range row {
				q = append(q, Point{x, y})
			}
		}
		// Process the queue, adding to it as flashes occur
		for len(q) > 0 {
			pt := q[0]
			q = q[1:]
			// Bounds check
			if pt.x < 0 || pt.y < 0 || pt.x >= len(grid) || pt.y >= len(grid[pt.x]) {
				continue
			}
			// Already flashed, move on
			if grid[pt.x][pt.y] == 10 {
				continue
			}
			// Increase the point
			grid[pt.x][pt.y]++
			// Flash (aka load up the queue with neighbors)
			if grid[pt.x][pt.y] == 10 {
				flashCount++
				q = append(q,
					Point{pt.x - 1, pt.y - 1},
					Point{pt.x - 1, pt.y},
					Point{pt.x - 1, pt.y + 1},
					Point{pt.x, pt.y - 1},
					Point{pt.x, pt.y + 1},
					Point{pt.x + 1, pt.y - 1},
					Point{pt.x + 1, pt.y},
					Point{pt.x + 1, pt.y + 1},
				)
			}
		}
		// Set flashed points to zero
		for x, row := range grid {
			for y := range row {
				if grid[x][y] == 10 {
					grid[x][y] = 0
				}
			}
		}
	}
	fmt.Println(flashCount)
}

func part2(grid [][]int) {
	stepCount := 1
	for {
		// Load up a queue of increases to make
		q := make([]Point, 0)
		for x, row := range grid {
			for y := range row {
				q = append(q, Point{x, y})
			}
		}
		// Process the queue, adding to it as flashes occur
		for len(q) > 0 {
			pt := q[0]
			q = q[1:]
			// Bounds check
			if pt.x < 0 || pt.y < 0 || pt.x >= len(grid) || pt.y >= len(grid[pt.x]) {
				continue
			}
			// Already flashed, move on
			if grid[pt.x][pt.y] == 10 {
				continue
			}
			// Increase the point
			grid[pt.x][pt.y]++
			// Flash (aka load up the queue with neighbors)
			if grid[pt.x][pt.y] == 10 {
				q = append(q,
					Point{pt.x - 1, pt.y - 1},
					Point{pt.x - 1, pt.y},
					Point{pt.x - 1, pt.y + 1},
					Point{pt.x, pt.y - 1},
					Point{pt.x, pt.y + 1},
					Point{pt.x + 1, pt.y - 1},
					Point{pt.x + 1, pt.y},
					Point{pt.x + 1, pt.y + 1},
				)
			}
		}
		// Set flashed points to zero
		flashCount := 0
		for x, row := range grid {
			for y := range row {
				if grid[x][y] == 10 {
					grid[x][y] = 0
					flashCount++
				}
			}
		}
		if flashCount == len(grid)*len(grid[0]) {
			fmt.Println(stepCount)
			return
		}
		stepCount++
	}
}
