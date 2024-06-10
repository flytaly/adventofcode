package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	connections []int
}

func dfs(nodes map[int]Node, curr int, visited *map[int]bool) {
	node := nodes[curr]
	for _, neighb := range node.connections {
		if (*visited)[neighb] {
			continue
		}
		(*visited)[neighb] = true
		dfs(nodes, neighb, visited)
	}
}

func parse(input []string) map[int]Node {
	nodes := map[int]Node{}
	for _, line := range input {
		splits := strings.Split(line, " <-> ")
		value, _ := strconv.Atoi(splits[0])
		right := strings.Split(splits[1], ", ")
		nodes[value] = Node{
			connections: utils.ToInts(right),
		}
	}
	return nodes
}

func P1(input []string) int {
	visited := map[int]bool{}
	dfs(parse(input), 0, &visited)
	return len(visited)
}
func main() {
	lines := []string{
		"0 <-> 2",
		"1 <-> 1",
		"2 <-> 0, 3, 4",
		"3 <-> 2, 4",
		"4 <-> 2, 3, 6",
		"5 <-> 6",
		"6 <-> 4, 5",
	}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("Part 1 =>", P1(lines))
}
