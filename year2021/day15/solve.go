package day15

import (
	_ "embed"
	"fmt"
	"masonheverett/advent-of-go/util"
	"math"
	"strings"
)

//go:embed input.txt
var input string

type Node struct {
	x, y int
}

type Edge struct {
	target Node
	weight int
}

type Graph struct {
	nodes       []Node
	adjacencies map[Node][]Edge
}

func newGraph(grid [][]int) *Graph {
	graph := &Graph{
		nodes:       make([]Node, 0),
		adjacencies: make(map[Node][]Edge),
	}
	for r := range grid {
		for c := range grid[r] {
			graph.addNode(Node{r, c})
			if r > 0 {
				graph.addAdjacency(Node{r, c}, Node{r - 1, c}, grid[r-1][c])
			}
			if r < len(grid)-1 {
				graph.addAdjacency(Node{r, c}, Node{r + 1, c}, grid[r+1][c])
			}
			if c > 0 {
				graph.addAdjacency(Node{r, c}, Node{r, c - 1}, grid[r][c-1])
			}
			if c < len(grid[r])-1 {
				graph.addAdjacency(Node{r, c}, Node{r, c + 1}, grid[r][c+1])
			}
		}
	}
	return graph
}

func (g *Graph) addNode(n Node) {
	g.nodes = append(g.nodes, n)
	g.adjacencies[n] = make([]Edge, 0)
}

func (g *Graph) addAdjacency(node1, node2 Node, weight int) {
	g.adjacencies[node1] = append(g.adjacencies[node1], Edge{node2, weight})
}

func (g *Graph) findPath(start, end Node) int {
	costs := make(map[Node]int)
	pq := PriorityQueue(make([]Edge, 0))
	costs[start] = 0
	for _, node := range g.nodes {
		if node != start {
			costs[node] = math.MaxInt
		}
	}
	pq.enqueue(Edge{start, 0})
	for len(pq) > 0 {
		shortest := pq.dequeue()
		current := shortest.target
		for _, neighbor := range g.adjacencies[current] {
			cost := costs[current] + neighbor.weight
			if cost < costs[neighbor.target] {
				costs[neighbor.target] = cost
				pq.enqueue(Edge{neighbor.target, cost})
			}
		}
	}
	return costs[end]
}

type PriorityQueue []Edge

func (p *PriorityQueue) enqueue(e Edge) {
	pq := *p
	if len(pq) == 0 {
		pq = append(pq, e)
	} else {
		added := false
		for i := 0; i < len(pq); i++ {
			if e.weight < pq[i].weight {
				pq = append(pq[:i+1], pq[i:]...)
				pq[i] = e
				added = true
				break
			}
		}
		if !added {
			pq = append(pq, e)
		}
	}
	*p = pq
}

func (p *PriorityQueue) dequeue() Edge {
	first := (*p)[0]
	*p = (*p)[1:]
	return first
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
		for j, ch := range strings.Split(line, "") {
			grid[i][j] = util.DecStringToInt(ch)
		}
	}
	return grid
}

func part1(grid [][]int) {
	graph := newGraph(grid)
	fmt.Println(graph.findPath(Node{0, 0}, Node{len(grid) - 1, len(grid[0]) - 1}))
}

func part2(grid [][]int) {
	fullGrid := make([][]int, len(grid)*5)
	for i := 0; i < 5; i++ {
		for r := range grid {
			fullGrid[i*len(grid)+r] = make([]int, len(grid[r])*5)
			for j := 0; j < 5; j++ {
				for c := range grid[r] {
					fullGrid[i*len(grid)+r][j*len(grid[r])+c] = ((grid[r][c] - 1 + i + j) % 9) + 1
				}
			}
		}
	}
	graph := newGraph(fullGrid)
	fmt.Println(graph.findPath(Node{0, 0}, Node{len(fullGrid) - 1, len(fullGrid[0]) - 1}))
}
