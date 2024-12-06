package main

import (
	"aoc/utils"
	. "aoc/utils/grid"
	"fmt"
	"image"
	"os"
)

var nextDir = map[image.Point]image.Point{
	ToTop:    ToRight,
	ToRight:  ToBottom,
	ToBottom: ToLeft,
	ToLeft:   ToTop,
}

func getStartingPoint(grid Grid[string]) image.Point {
	for p, v := range grid.PointsIter() {
		if v == "^" {
			return p
		}
	}
	return image.Point{0, 0}
}

func PartOne(lines []string) {
	grid := NewStringGrid(lines)
	pos := getStartingPoint(grid)
	count := 1
	grid.Set(pos, "X")
	dir := ToTop

	for {
		nextPos := pos.Add(dir)
		if grid.At(nextPos) == "#" {
			dir = nextDir[dir]
			continue
		}

		if !grid.IsInside(nextPos) {
			break
		}

		if grid.At(nextPos) == "." {
			grid.Set(nextPos, "X")
			count++
		}

		pos = nextPos

	}
	fmt.Println(grid)
	fmt.Println("Part 1:", count)
}

func main() {
	lines := []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#.........",
		"......#...",
	}

	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}

	PartOne(lines)
}
