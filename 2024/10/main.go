package main

import (
	"aoc/utils"
	. "aoc/utils/grid"
	"fmt"
	"image"
	"maps"
	"os"
)

var Pt = image.Pt

func mergeMaps[K comparable, V any](a, b map[K]V) map[K]V {
	res := maps.Clone(a)
	if res == nil {
		res = map[K]V{}
	}
	for k, v := range b {
		res[k] = v
	}
	return res
}

func PartOne(lines []string) int {
	heights := NewGrid[int](len(lines[0]), len(lines))
	for i, row := range lines {
		for j, col := range row {
			heights.Values[image.Point{j, i}] = int(col) - '0'
		}
	}
	grid := NewGrid[map[image.Point]bool](heights.Right+1, heights.Bottom+1)

	for height := 9; height >= 1; height-- {
		for point, h := range heights.PointsIter() {
			if h != height {
				continue
			}
			if h == 9 {
				grid.Set(point, map[image.Point]bool{point: true})
			}
			for _, d := range []image.Point{ToTop, ToRight, ToBottom, ToLeft} {
				if p2 := point.Add(d); heights.At(p2) == heights.At(point)-1 {
					grid.Set(p2, mergeMaps(grid.At(point), grid.At(p2)))
				}
			}
		}
	}

	count := 0
	for p, v := range heights.PointsIter() {
		if v == 0 {
			count += len(grid.At(p))
		}
	}

	return count
}

func PartTwo(lines []string) int {
	heights := NewGrid[int](len(lines[0]), len(lines))
	for i, row := range lines {
		for j, col := range row {
			heights.Values[image.Point{j, i}] = int(col) - '0'
		}
	}
	grid := NewGrid[int](heights.Right+1, heights.Bottom+1)

	for height := 9; height >= 1; height-- {
		for point, h := range heights.PointsIter() {
			if h != height {
				continue
			}
			if h == 9 {
				grid.Set(point, 1)
			}
			for _, d := range []image.Point{ToTop, ToRight, ToBottom, ToLeft} {
				if p2 := point.Add(d); heights.At(p2) == heights.At(point)-1 {
					grid.Set(p2, grid.At(point)+grid.At(p2))
				}
			}
		}
	}

	count := 0
	for p, v := range heights.PointsIter() {
		if v == 0 {
			count += grid.At(p)
		}
	}

	return count
}

func main() {
	lines := []string{
		"89010123",
		"78121874",
		"87430965",
		"96549874",
		"45678903",
		"32019012",
		"01329801",
		"10456732",
	}

	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}

	fmt.Println("Part 1: ", PartOne(lines))
	fmt.Println("Part 2: ", PartTwo(lines))
}
