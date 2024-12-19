package main

import (
	"aoc/utils"
	. "aoc/utils/grid"
	"fmt"
	"image"
	"os"
	"strings"
)

func findPath(grid Grid[string]) int {
	start := image.Pt(0, 0)
	end := image.Pt(grid.Right, grid.Bottom)
	queue := []image.Point{start}
	dist := map[image.Point]int{start: 0}
	for i := 0; len(queue) > 0; i++ {
		current := queue[0]
		queue = queue[1:]
		for _, d := range []image.Point{ToLeft, ToTop, ToRight, ToBottom} {
			neighb := current.Add(d)
			if dist[neighb] != 0 || grid.At(neighb) != "." {
				continue
			}
			distN := dist[current] + 1
			if neighb == end {
				return distN
			}
			dist[neighb] = distN
			queue = append(queue, neighb)
		}
	}

	return -1
}

func PartOne(input []string, size int, steps int) int {
	grid := NewGrid[string](size, size)
	grid.Fill(".")
	for _, bytes := range input[:steps] {
		sp := utils.ToInts(strings.Split(bytes, ","))
		grid.Set(image.Pt(sp[0], sp[1]), "#")
	}

	return findPath(grid)
}

func main() {
	lines := []string{
		"5,4",
		"4,2",
		"4,5",
		"3,0",
		"2,1",
		"6,3",
		"2,4",
		"1,5",
		"0,6",
		"3,3",
		"2,6",
		"5,1",
		"1,2",
		"5,5",
		"2,5",
		"6,5",
		"1,4",
		"0,4",
		"6,4",
		"1,1",
		"6,1",
		"1,0",
		"0,5",
		"1,6",
		"2,0",
	}

	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
		fmt.Println("Part 1:", PartOne(lines, 71, 1024))
	} else {
		fmt.Println("Part 1:", PartOne(lines, 7, 12))
	}

}
