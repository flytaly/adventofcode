package d23

import (
	"fmt"
	"slices"
	"time"

	"k8s.io/apimachinery/pkg/util/sets"
)

var dirs = []C{-1, -1i, 1, 1i}

func getNodes(grid map[C]string) []C {
	nodes := []C{}
	for cell, val := range grid {
		if val == "#" {
			continue
		}
		neighbors := 0
		for _, dir := range dirs {
			pos := cell + dir
			if grid[pos] != "" && grid[pos] != "#" {
				neighbors++
			}
		}
		if neighbors > 2 {
			nodes = append(nodes, cell)
		}
	}
	return nodes
}

type C = complex128
type P struct {
	pos  C
	dist int
}
type CGraph = map[C]sets.Set[P]
type CSet = sets.Set[C]

func dfsMax(p C, end C, graph CGraph, visited CSet) int {
	if p == end {
		return 0
	}

	res := -1
	visited.Insert(p)
	for node := range graph[p] {
		if visited.Has(node.pos) {
			continue
		}

		if d := dfsMax(node.pos, end, graph, visited); d != -1 {
			res = max(res, d+node.dist)
		}
	}
	visited.Delete(p)

	return res
}

func PartTwo(lines []string) {
	grid, startPos, endPos := parse(lines)

	nodes := []C{startPos, endPos}
	nodes = append(nodes, getNodes(grid)...)

	// fmt.Println("Graph nodes", nodes)

	graph := CGraph{}
	for _, node := range nodes {
		graph[node] = sets.New[P]()
		visited := CSet{}
		visited.Insert(node)
		path := []P{{pos: node, dist: 0}}

		for len(path) > 0 {
			point := path[len(path)-1]
			path = path[:len(path)-1]

			if point.dist > 0 && slices.Contains(nodes, point.pos) {
				graph[node].Insert(point)
				continue
			}

			for _, dir := range dirs {
				pos := point.pos + dir
				if grid[pos] == "" || grid[pos] == "#" || visited.Has(pos) {
					continue
				}
				visited.Insert(pos)
				path = append(path, P{pos: pos, dist: point.dist + 1})
			}
		}
	}

	ts := time.Now()
	result := dfsMax(startPos, endPos, graph, CSet{})

	fmt.Printf("Part 2: %d (%s)\n", result, time.Since(ts))
}
