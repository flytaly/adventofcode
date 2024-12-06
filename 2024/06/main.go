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

walk:
	for dir := ToTop; ; {
		nextPos := pos.Add(dir)

		switch grid.At(nextPos) {
		case "":
			break walk
		case ".":
			grid.Set(nextPos, "X")
			count++
		case "#":
			dir = nextDir[dir]
			continue walk
		}
		pos = nextPos
	}

	// fmt.Println(grid)
	fmt.Println("Part 1:", count)
}

func move(grid Grid[string], pos *image.Point, dir *image.Point) {
	for {
		nextPos := pos.Add(*dir)
		cell := grid.At(nextPos)
		if cell == "#" {
			*dir = nextDir[*dir]
			continue
		}
		*pos = nextPos
		return
	}
}

// Tortoise and Hare Algorithm
func hasCycle(grid Grid[string], start image.Point) bool {
	posT, dirT := start, ToTop
	posH, dirH := start, ToTop
	for {
		move(grid, &posH, &dirH)
		move(grid, &posH, &dirH)
		move(grid, &posT, &dirT)
		if !grid.IsInside(posH) {
			return false
		}
		if posH.Eq(posT) && dirH.Eq(dirT) {
			return true
		}
	}
}

func PartTwo(lines []string) {
	grid := NewStringGrid(lines)
	start := getStartingPoint(grid)

	count := 0
	checked := map[image.Point]bool{}
	for pos, dir := start, ToTop; ; {
		move(grid, &pos, &dir)
		if !grid.IsInside(pos) {
			break
		}
		if _, ok := checked[pos]; ok {
			continue
		}
		if cell := grid.At(pos); cell != "#" {
			grid.Set(pos, "#")
			if hasCycle(grid, start) {
				count++
			}
			grid.Set(pos, cell)
			checked[pos] = true
		}
	}
	fmt.Println("Part 2:", count)
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
	PartTwo(lines)
}
