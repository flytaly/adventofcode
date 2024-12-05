package main

import (
	"aoc/utils"
	G "aoc/utils/grid"
	"fmt"
	"image"
	"os"
	"slices"
)

func PartOne(lines []string) {
	grid := G.NewRuneGrid(lines)

	target := []rune("XMAS")
	count := 0
	for p := range grid.PointsIter() {
		for _, dir := range G.Dirs {
			span := grid.Span(p, dir, 4)
			if slices.Equal(span, target) {
				count++
			}
		}
	}
	fmt.Println("Part 1:", count)
}

func PartTwo(lines []string) {
	grid := G.NewRuneGrid(lines)

	count := 0

	isMAS := func(r []rune) bool {
		s := string(r)
		return s == "MAS" || s == "SAM"
	}

	diag1 := []image.Point{{-1, -1}, {1, 1}}
	diag2 := []image.Point{{1, -1}, {-1, 1}}

	for p := range grid.PointsIter() {
		d1 := grid.Span(p.Add(diag1[0]), diag1[1], 3)
		d2 := grid.Span(p.Add(diag2[0]), diag2[1], 3)
		if isMAS(d1) && isMAS(d2) {
			count++
		}
	}
	fmt.Println("Part 2:", count)
}

func main() {
	lines := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}

	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}

	PartOne(lines)
	PartTwo(lines)
}
