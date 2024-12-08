package main

import (
	"aoc/utils"
	. "aoc/utils/grid"
	"fmt"
	"image"
	"iter"
	"os"
)

// Iterator for all combinations of 2 elements
func Combi2[T any](items []T) iter.Seq[[2]T] {
	return func(yield func([2]T) bool) {
		for i, it1 := range items {
			for _, it2 := range items[i+1:] {
				if !yield([2]T{it1, it2}) {
					return
				}
			}
		}
	}
}

func Solve(lines []string, isPartTwo bool) (count int) {
	grid := NewStringGrid(lines)
	freq := map[string][]image.Point{}

	for i, p := range grid.PointsIter() {
		if p != "." {
			freq[p] = append(freq[p], i)
		}
	}

	check := func(anti image.Point) {
		if grid.At(anti) != "#" {
			grid.Set(anti, "#")
			count++
		}
	}

	var createAnti func(point, diff image.Point)
	createAnti = func(point, diff image.Point) {
		anti := point.Add(diff)
		if !grid.IsInside(anti) {
			return
		}
		check(anti)
		if isPartTwo {
			createAnti(anti, diff)
		}
	}

	for _, coords := range freq {
		for pair := range Combi2(coords) {
			a, b := pair[0], pair[1]
			if isPartTwo {
				check(a)
				check(b)
			}
			createAnti(a, a.Sub(b))
			createAnti(b, b.Sub(a))
		}
	}

	return count
}

func main() {
	lines := []string{
		"............",
		"........0...",
		".....0......",
		".......0....",
		"....0.......",
		"......A.....",
		"............",
		"............",
		"........A...",
		".........A..",
		"............",
		"............",
	}

	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}

	fmt.Println("PartOne: ", Solve(lines, false))
	fmt.Println("PartTwo: ", Solve(lines, true))
}
