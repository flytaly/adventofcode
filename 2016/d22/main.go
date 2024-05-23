package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"regexp"
)

type Node struct {
	x, y, size, used, avail, usePerc int
}

func (n Node) key() string {
	return fmt.Sprintf("%d_%d", n.x, n.y)
}

type NodesMap struct {
	Nodes      map[string]Node
	MaxX, MaxY int
}

func (n *NodesMap) Set(node Node) {
	n.Nodes[node.key()] = node
	if node.x > n.MaxX {
		n.MaxX = node.x
	}
	if node.y > n.MaxY {
		n.MaxY = node.y
	}
}

func (n NodesMap) Get(x, y int) Node {
	key := fmt.Sprintf("%d_%d", x, y)
	return n.Nodes[key]
}

func (n NodesMap) String() string {
	res := ""

	for x := 0; x <= n.MaxX; x++ {
		res += fmt.Sprintf("%7d", x)
	}
	res += "\n"

	for y := 0; y <= n.MaxY; y++ {
		res += fmt.Sprintf("%2d", y)
		for x := 0; x <= n.MaxX; x++ {
			node := n.Get(x, y)
			res += fmt.Sprintf("%7s", fmt.Sprintf("%d/%d", node.used, node.avail))
		}
		res += "\n"
	}
	return res
}

func parse(lines []string) (NodesMap, []Node) {
	nodeMap := NodesMap{Nodes: make(map[string]Node)}
	nodeList := []Node{}
	re := regexp.MustCompile(`/dev/grid/node-x(\d+)-y(\d+)\s+(\d+)T\s+(\d+)T\s+(\d+)T\s+(\d+)\%`)
	for _, line := range lines[2:] {
		n := utils.ToInts(re.FindStringSubmatch(line))
		node := Node{n[1], n[2], n[3], n[4], n[5], n[6]}
		nodeMap.Set(node)
		nodeList = append(nodeList, node)
	}
	return nodeMap, nodeList
}

func P1(lines []string) int {
	_, nodeList := parse(lines)

	count := 0
	for i := 0; i < len(nodeList); i++ {
		for j := 0; j < len(nodeList); j++ {
			a, b := nodeList[i], nodeList[j]
			if a.key() != b.key() && a.used > 0 && a.used <= b.avail {
				count++
			}
		}
	}

	return count
}

func P2(lines []string) int {
	nodeMap, _ := parse(lines)
	fmt.Println("\n", nodeMap)
	// can be solved visually by moving an empty cell to a target node bypassing high-capacity nodes
	// and then moving our data through this empty node, which would require 5 moves per step
	return 0
}

func main() {
	lines := []string{
		"root@ebhq-gridcenter# df -h",
		"Filesystem            Size  Used  Avail  Use%",
		"/dev/grid/node-x0-y0   10T    8T     2T   80%",
		"/dev/grid/node-x0-y1   11T    6T     5T   54%",
		"/dev/grid/node-x0-y2   32T   28T     4T   87%",
		"/dev/grid/node-x1-y0    9T    7T     2T   77%",
		"/dev/grid/node-x1-y1    8T    0T     8T    0%",
		"/dev/grid/node-x1-y2   11T    7T     4T   63%",
		"/dev/grid/node-x2-y0   10T    6T     4T   60%",
		"/dev/grid/node-x2-y1    9T    8T     1T   88%",
		"/dev/grid/node-x2-y2    9T    6T     3T   66%",
	}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("PartOne: ", P1(lines))
	fmt.Println("PartTwo: ", P2(lines))
}
