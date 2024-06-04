package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type G map[string]*Node

type Node struct {
	name     string
	weight   int
	children G
	parent   *Node
}

func parseNode(str string) *Node {
	re := regexp.MustCompile(`(\w+)(?:\s\((\d+)\))`)
	match := re.FindAllStringSubmatch(str, -1)
	node := Node{name: match[0][1], children: make(G)}
	if len(match[0]) == 3 {
		node.weight, _ = strconv.Atoi(match[0][2])
	}
	return &node

}

func makeGraph(input []string) (nodes G, root string) {
	nodes = G{}
	for _, line := range input {
		split := strings.Split(line, " -> ")
		node := parseNode(split[0])
		if existed, exist := nodes[node.name]; exist {
			existed.weight = node.weight
			node = existed
		}
		nodes[node.name] = node

		if len(split) <= 1 {
			continue
		}
		children := strings.Split(split[1], ", ")
		for _, childName := range children {
			if childNode, exist := nodes[childName]; exist {
				childNode.parent = node
				node.children[childNode.name] = childNode
			} else {
				newChild := &Node{name: childName, parent: node, children: make(G)}
				node.children[newChild.name] = newChild
				nodes[childName] = newChild
			}
		}
	}
	for _, n := range nodes {
		if n.parent == nil {
			root = n.name
			break
		}
	}
	return nodes, root
}

func P1(input []string) string {
	_, root := makeGraph(input)
	return root
}

func allSame(nums map[string]int) bool {
	prev := -1
	for _, v := range nums {
		if prev != -1 && v != prev {
			return false
		}
		prev = v
	}
	return true

}

func getException(nums map[string]int) (string, int) {
	vals := map[int]int{}
	for _, n := range nums {
		vals[n]++
	}
	var exception int
	var others int
	for num, count := range vals {
		if count == 1 {
			exception = num
		} else {
			others = num
		}
	}
	diff := exception - others
	name := ""

	for n, val := range nums {
		if val == exception {
			name = n
		}
	}

	return name, diff
}

// recursively found the deepest unbalanced node and return it's value after rebalnce
func getWeight(g G, root string) (total int, children map[string]int, answer int) {
	node := g[root]
	total = node.weight

	childrenWeights := map[string]int{}
	for _, child := range node.children {
		w, _, res := getWeight(g, child.name)
		if res != 0 {
			answer = res
		}
		total += w
		childrenWeights[child.name] = w
	}

	if answer == 0 && !allSame(childrenWeights) {
		name, diff := getException(childrenWeights)
		answer = g[name].weight - diff
	}

	return total, childrenWeights, answer
}

func P2(input []string) int {
	nodes, root := makeGraph(input)
	_, _, result := getWeight(nodes, root)
	return result
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
	fmt.Println("Part 2 =>", P2(lines))
}
