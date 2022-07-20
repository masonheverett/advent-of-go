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
	return &Node{
		neighbors: make([]string, 0),
		blocked:   false,
	}
}

func (n *Node) addNeighbor(s string) {
	n.neighbors = append(n.neighbors, s)
}

type Graph2 struct {
	nodes      map[string]*Node
	smallTwice bool
}

func Solve() {
	part1(parseInput())
	part2(parseInput())
}

func parseInput() map[string]*Node {
	lines := strings.Split(input, "\n")
	graph := make(map[string]*Node)
	for _, line := range lines {
		edge := strings.Split(line, "-")
		if _, exists := graph[edge[0]]; !exists {
			graph[edge[0]] = newNode()
		}
		graph[edge[0]].addNeighbor(edge[1])
		if _, exists := graph[edge[1]]; !exists {
			graph[edge[1]] = newNode()
		}
		graph[edge[1]].addNeighbor(edge[0])
	}
	return graph
}

func part1(graph map[string]*Node) {
	pathCount = 0
	findPaths1(graph, "start")
	fmt.Println(pathCount)
}

func part2(graph map[string]*Node) {
	pathCount = 0
	findPaths2(Graph2{graph, false}, "start")
	fmt.Println(pathCount)
}

func copyGraph1(graph map[string]*Node) map[string]*Node {
	newGraph := make(map[string]*Node)
	for key, val := range graph {
		newGraph[key] = newNode()
		newGraph[key].blocked = val.blocked
		for _, neighbor := range val.neighbors {
			newGraph[key].addNeighbor(neighbor)
		}
	}
	return newGraph
}

func findPaths1(graph map[string]*Node, next string) {
	if graph[next].blocked {
		return
	}
	if next == "end" {
		pathCount++
		return
	}
	if strings.ToLower(next) == next {
		graph[next].blocked = true
	}
	for _, neighbor := range graph[next].neighbors {
		findPaths1(copyGraph1(graph), neighbor)
	}
}

func copyGraph2(graph Graph2) Graph2 {
	newNodes := make(map[string]*Node)
	for key, val := range graph.nodes {
		newNodes[key] = newNode()
		newNodes[key].blocked = val.blocked
		for _, neighbor := range val.neighbors {
			newNodes[key].addNeighbor(neighbor)
		}
	}
	return Graph2{nodes: newNodes, smallTwice: graph.smallTwice}
}

func findPaths2(graph Graph2, next string) {
	if next == "start" && graph.nodes[next].blocked {
		return
	}
	if graph.nodes[next].blocked && graph.smallTwice {
		return
	}
	if next == "end" {
		pathCount++
		return
	}
	if strings.ToLower(next) == next {
		if next == "start" || !graph.nodes[next].blocked {
			graph.nodes[next].blocked = true
		} else {
			graph.smallTwice = true
		}
	}
	for _, neighbor := range graph.nodes[next].neighbors {
		findPaths2(copyGraph2(graph), neighbor)
	}
}
