package day12

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

var pathCount int

type Node struct {
	neighbors []string
	blocked   bool
}

func newNode() *Node {
	return &Node{neighbors: make([]string, 0), blocked: false}
}

func (n *Node) addNeighbor(s string) {
	n.neighbors = append(n.neighbors, s)
}

type Graph struct {
	nodes      map[string]*Node
	doubleUsed bool
}

func (g *Graph) copy() *Graph {
	newNodes := make(map[string]*Node)
	for key, val := range g.nodes {
		newNodes[key] = newNode()
		newNodes[key].blocked = val.blocked
		for _, neighbor := range val.neighbors {
			newNodes[key].addNeighbor(neighbor)
		}
	}
	return &Graph{nodes: newNodes, doubleUsed: g.doubleUsed}
}

func (g *Graph) findPaths(next string) {
	if next == "end" {
		pathCount++
		return
	}
	if g.nodes[next].blocked && (next == "start" || g.doubleUsed) {
		return
	}
	if strings.ToLower(next) == next {
		if next == "start" || !g.nodes[next].blocked {
			g.nodes[next].blocked = true
		} else {
			g.doubleUsed = true
		}
	}
	for _, neighbor := range g.nodes[next].neighbors {
		g.copy().findPaths(neighbor)
	}
}

func Solve() {
	part1(parseInput())
	part2(parseInput())
}

func parseInput() Graph {
	lines := strings.Split(input, "\n")
	nodes := make(map[string]*Node)
	for _, line := range lines {
		edge := strings.Split(line, "-")
		if _, exists := nodes[edge[0]]; !exists {
			nodes[edge[0]] = newNode()
		}
		nodes[edge[0]].addNeighbor(edge[1])
		if _, exists := nodes[edge[1]]; !exists {
			nodes[edge[1]] = newNode()
		}
		nodes[edge[1]].addNeighbor(edge[0])
	}
	return Graph{nodes: nodes, doubleUsed: false}
}

func part1(graph Graph) {
	pathCount = 0
	graph.doubleUsed = true
	graph.findPaths("start")
	fmt.Println(pathCount)
}

func part2(graph Graph) {
	pathCount = 0
	graph.findPaths("start")
	fmt.Println(pathCount)
}
