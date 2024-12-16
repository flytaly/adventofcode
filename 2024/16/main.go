package main

import (
	"aoc/utils"
	. "aoc/utils/grid"
	"cmp"
	"fmt"
	"image"
	"os"
	"slices"
)

func parse(lines []string) (grid Grid[string], start image.Point, end image.Point) {
	grid = NewStringGrid(lines)

	for p, v := range grid.PointsIter() {
		if v == "S" {
			start = p
		}
		if v == "E" {
			end = p
		}
	}

	return grid, start, end
}

type FrontTile struct {
	p    image.Point
	dir  image.Point
	cost int
}

func (t FrontTile) calcCost(dir image.Point) int {
	if t.dir.Eq(dir) {
		return t.cost + 1
	}
	return t.cost + 1001

}

type Node struct {
	p   image.Point
	dir image.Point
}

func PartOne(lines []string) (result int) {
	grid, start, end := parse(lines)

	frontier := []FrontTile{{p: start, cost: 0, dir: ToRight}}

	costs := map[Node]int{}

outer:
	for len(frontier) > 0 {
		slices.SortFunc(frontier, func(a, b FrontTile) int {
			return cmp.Compare(b.cost, a.cost)
		})
		current := frontier[len(frontier)-1]
		frontier = frontier[:len(frontier)-1]
		for _, dir := range []image.Point{ToTop, ToRight, ToBottom, ToLeft} {
			nextP := current.p.Add(dir)
			newCost := current.calcCost(dir)
			if nextP == end {
				result = newCost
				break outer
			}
			if grid.At(nextP) == "#" {
				continue
			}
			t := Node{p: nextP, dir: dir}
			if cost, ok := costs[t]; !ok || cost > newCost {
				costs[t] = newCost
				frontier = append(frontier, FrontTile{p: nextP, cost: newCost, dir: dir})
			}
		}
	}

	return result
}

func main() {
	lines := []string{
		"#################",
		"#...#...#...#..E#",
		"#.#.#.#.#.#.#.#.#",
		"#.#.#.#...#...#.#",
		"#.#.#.#.###.#.#.#",
		"#...#.#.#.....#.#",
		"#.#.#.#.#.#####.#",
		"#.#...#.#.#.....#",
		"#.#.#####.#.###.#",
		"#.#.#.......#...#",
		"#.#.###.#####.###",
		"#.#.#...#.....#.#",
		"#.#.#.#####.###.#",
		"#.#.#.........#.#",
		"#.#.#.#########.#",
		"#S#.............#",
		"#################",
	}

	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}

	fmt.Println("Part 1: ", PartOne(lines))
}
