package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Node struct {
	name  string
	value int
	next  *Node
	prev  *Node
}

func parseNode(str string) *Node {
	re := regexp.MustCompile(`(\w+)(?:\s\((\d+)\))`)
	match := re.FindAllStringSubmatch(str, -1)
	node := Node{name: match[0][1]}
	if len(match[0]) == 3 {
		node.value, _ = strconv.Atoi(match[0][2])
	}
	return &node

}

func P1(input []string) string {
	nodes := map[string]*Node{}
	for _, line := range input {
		split := strings.Split(line, " -> ")
		node := parseNode(split[0])
		if _, exist := nodes[node.name]; !exist {
			nodes[node.name] = node
		}
		if len(split) <= 1 {
			continue
		}
		children := strings.Split(split[1], ", ")
		for _, child := range children {
			if childNode, exist := nodes[child]; exist {
				childNode.prev = node
				node.next = childNode
			} else {
				childNode := &Node{name: child, prev: node}
				nodes[child] = childNode
			}
		}
	}
	for _, n := range nodes {
		if n.prev == nil {
			return n.name
		}
	}
	return ""
}

func main() {
	lines := []string{
		"pbga (66)",
		"xhth (57)",
		"ebii (61)",
		"havc (66)",
		"ktlj (57)",
		"fwft (72) -> ktlj, cntj, xhth",
		"qoyq (66)",
		"padx (45) -> pbga, havc, qoyq",
		"tknk (41) -> ugml, padx, fwft",
		"jptl (61)",
		"ugml (68) -> gyxo, ebii, jptl",
		"gyxo (61)",
		"cntj (57)",
	}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("Part 1 =>", P1(lines))
}
