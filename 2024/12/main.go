package main

import (
	"aoc/utils"
	. "aoc/utils/grid"
	"fmt"
	"image"
	"os"
)

var dirs = []image.Point{ToTop, ToLeft, ToRight, ToBottom}

func regionDFS(grid Grid[string], visited map[image.Point]bool, start image.Point) (perimeter int, area int) {
	visited[start] = true
	area = 1
	for _, d := range dirs {
		next := start.Add(d)
		if grid.At(next) != grid.At(start) {
			perimeter++
			continue
		}
		if visited[next] {
			continue
		}
		per, ar := regionDFS(grid, visited, next)
		perimeter += per
		area += ar
	}
	return perimeter, area
}

type Region struct {
	Per  int
	Area int
}

func PartOne(lines []string) (price int) {
	grid := NewStringGrid(lines)
	visited := map[image.Point]bool{}
	regions := map[string][]Region{}

	for p, label := range grid.PointsIter() {
		if visited[p] {
			continue
		}
		per, ar := regionDFS(grid, visited, p)
		regions[label] = append(regions[label], Region{Per: per, Area: ar})
	}

	for _, islands := range regions {
		for _, reg := range islands {
			price += reg.Per * reg.Area
		}
	}

	return price
}

// count the number of vertices around a point, which can be a corner of a polygon
func countVertices[T comparable](grid Grid[T], p image.Point) (vertices int) {
	dd := []image.Point{ToLeft, ToTop, ToRight, ToBottom}
	for i := range dd {
		d1, d2 := dd[i], dd[(i+1)%4]
		id := grid.At(p)
		adj1, adj2, diag := p.Add(d1), p.Add(d2), p.Add(d1).Add(d2)
		// outer corner
		if grid.At(adj1) != id && grid.At(adj2) != id {
			vertices++
			continue
		}
		// inner corner
		if grid.At(adj1) == id && grid.At(adj2) == id && grid.At(diag) != id {
			vertices++
		}
	}

	return vertices
}

// In a polygon, the number of vertices equals the number of edges.
func regionBFS[T comparable](grid Grid[T], visited map[image.Point]bool, start image.Point) (edges int, area int) {
	queue := []image.Point{start}
	visited[start] = true
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		area++
		for _, d := range dirs {
			next := p.Add(d)
			if grid.At(next) == grid.At(p) && !visited[next] {
				visited[next] = true
				queue = append(queue, next)
			}
		}
		edges += countVertices(grid, p)
	}
	return edges, area
}

func PartTwo(lines []string) (price int) {
	type Polygon struct {
		Edges int
		Area  int
	}
	grid := NewStringGrid(lines)
	visited := map[image.Point]bool{}
	regions := []Polygon{}
	for p, _ := range grid.PointsIter() {
		if visited[p] {
			continue
		}
		edges, ar := regionBFS(grid, visited, p)
		regions = append(regions, Polygon{Edges: edges, Area: ar})
	}

	for _, reg := range regions {
		price += reg.Edges * reg.Area
	}

	return price
}

func main() {
	lines := []string{
		"RRRRIICCFF",
		"RRRRIICCCF",
		"VVRRRCCFFF",
		"VVRCCCJFFF",
		"VVVVCJJCFE",
		"VVIVCCJJEE",
		"VVIIICJJEE",
		"MIIIIIJJEE",
		"MIIISIJEEE",
		"MMMISSJEEE",
	}

	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}

	fmt.Println("Part 1: ", PartOne(lines))
	fmt.Println("Part 2: ", PartTwo(lines))
}
