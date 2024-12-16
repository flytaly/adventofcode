package main

import (
	"aoc/utils"
	. "aoc/utils/grid"
	"cmp"
	"fmt"
	"image"
	"math"
	"os"
	"slices"
	"time"
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
	path []image.Point
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

func Solve(lines []string) (int, int) {
	grid, start, end := parse(lines)

	frontier := []FrontTile{{p: start, cost: 0, dir: ToRight, path: []image.Point{start}}}
	costs := map[Node]int{}
	paths := [][]image.Point{}

	finalCost := math.MaxInt

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
			if nextP == end && finalCost >= newCost {
				finalCost = newCost
				paths = append(paths, append(current.path, end))
				continue outer
			}
			if grid.At(nextP) == "#" {
				continue
			}
			t := Node{p: nextP, dir: dir}
			if cost, ok := costs[t]; !ok || cost >= newCost {
				costs[t] = newCost

				path := append(slices.Clone(current.path), nextP)
				fr := FrontTile{p: nextP, cost: newCost, dir: dir, path: path}
				frontier = append(frontier, fr)
			}
		}
	}

	unique := make(map[image.Point]struct{})
	for _, path := range paths {
		for _, p := range path {
			grid.Set(p, "O")
			unique[p] = struct{}{}
		}
	}

	return finalCost, len(unique)
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

	ts := time.Now()
	p1, p2 := Solve(lines)
	fmt.Println("Part 1: ", p1)
	fmt.Println("Part 2: ", p2)
	fmt.Println("Time: ", time.Since(ts))
}
