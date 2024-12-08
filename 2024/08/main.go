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

func PartOne(lines []string) {
	grid := NewStringGrid(lines)
	freq := map[string][]image.Point{}

	for i, p := range grid.PointsIter() {
		if p != "." {
			freq[p] = append(freq[p], i)
		}
	}

	isEmpty := func(p image.Point) bool {
		return grid.At(p) != "" && grid.At(p) != "#"
	}

	count := 0
	for _, coords := range freq {
		for pair := range Combi2(coords) {
			a, b := pair[0], pair[1]
			if anti := a.Add(a.Sub(b)); isEmpty(anti) {
				grid.Set(anti, "#")
				count++
			}
			if anti := b.Add(b.Sub(a)); isEmpty(anti) {
				grid.Set(anti, "#")
				count++
			}
		}
	}

	// fmt.Println(grid)

	fmt.Println("Part  1:", count)
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

	PartOne(lines)
}
