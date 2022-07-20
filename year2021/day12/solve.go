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

func copyGraph(graph Graph) Graph {
	newNodes := make(map[string]*Node)
	for key, val := range graph.nodes {
		newNodes[key] = newNode()
		newNodes[key].blocked = val.blocked
		for _, neighbor := range val.neighbors {
			newNodes[key].addNeighbor(neighbor)
		}
	}
	return Graph{nodes: newNodes, doubleUsed: graph.doubleUsed}
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
	findPaths(graph, "start")
	fmt.Println(pathCount)
}

func part2(graph Graph) {
	pathCount = 0
	findPaths(graph, "start")
	fmt.Println(pathCount)
}

func findPaths(graph Graph, next string) {
	if next == "end" {
		pathCount++
		return
	}
	if graph.nodes[next].blocked && (next == "start" || graph.doubleUsed) {
		return
	}
	if strings.ToLower(next) == next {
		if next == "start" || !graph.nodes[next].blocked {
			graph.nodes[next].blocked = true
		} else {
			graph.doubleUsed = true
		}
	}
	for _, neighbor := range graph.nodes[next].neighbors {
		findPaths(copyGraph(graph), neighbor)
	}
}
