package main

import (
	"aoc/utils"
	. "aoc/utils/grid"
	"fmt"
	"image"
	"os"
)

var dirs = []image.Point{ToTop, ToLeft, ToRight, ToBottom}

func regionBFS(grid Grid[string], visited map[image.Point]bool, start image.Point) (perimeter int, area int) {
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
		per, ar := regionBFS(grid, visited, next)
		perimeter += per
		area += ar
	}
	return perimeter, area
}

type Region struct {
	Per  int
	Area int
}

func PartOne(lines []string) int {
	grid := NewStringGrid(lines)
	visited := map[image.Point]bool{}
	regions := map[string][]Region{}

	// perimeters := map[string][]int{}
	// areas := map[string][]int{}

	for p, label := range grid.PointsIter() {
		if visited[p] {
			continue
		}
		per, ar := regionBFS(grid, visited, p)
		regions[label] = append(regions[label], Region{Per: per, Area: ar})
	}

	price := 0
	for _, islands := range regions {
		for _, reg := range islands {
			price += reg.Per * reg.Area
		}
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
}
