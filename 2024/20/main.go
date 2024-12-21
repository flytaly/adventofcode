package main

import (
	"aoc/utils"
	. "aoc/utils/grid"
	"fmt"
	"image"
	"os"
	"slices"
)

func parse(input []string) (grid Grid[string], start image.Point, end image.Point) {
	grid = NewStringGrid(input)
	for p, v := range grid.PointsIter() {
		switch v {
		case "S":
			start = p
			grid.Set(p, ".")
		case "E":
			end = p
			grid.Set(p, ".")
		}
	}
	return grid, start, end
}

var neighbors = []image.Point{ToTop, ToRight, ToBottom, ToLeft}

func getDistMap(grid Grid[string], start image.Point, end image.Point) map[image.Point]int {
	dist := map[image.Point]int{start: 0}
	for curr := start; !curr.Eq(end); {
		for _, d := range neighbors {
			next := curr.Add(d)
			if _, seen := dist[next]; !seen && grid.At(next) == "." {
				dist[next] = dist[curr] + 1
				curr = next
				break
			}
		}
	}
	return dist
}

func PartOne(input []string, minSave int) (count int) {
	grid, start, end := parse(input)
	dist := getDistMap(grid, start, end)

	wallsChecked := map[image.Point]bool{}
	for pathCell, _ := range dist {
		for wall, v := range grid.Neighbs(pathCell) {
			if v != "#" || wallsChecked[wall] {
				continue
			}
			dists := []int{}
			for pathCell2, value := range grid.Neighbs(wall) {
				if value == "." {
					dists = append(dists, dist[pathCell2])
				}
			}
			if slices.Max(dists)-slices.Min(dists)-2 >= minSave {
				count++
			}
			wallsChecked[wall] = true
		}
	}

	return count
}

func PartTwo(input []string, minSave int) (count int) {
	grid, start, end := parse(input)
	dist := getDistMap(grid, start, end)

	for p1, d1 := range dist {
		for p2, d2 := range dist {
			manhattan := utils.Abs(p1.X-p2.X) + utils.Abs(p1.Y-p2.Y)
			if manhattan <= 20 && d2-d1-manhattan >= minSave {
				count++
			}
		}
	}

	return count
}

func main() {
	lines := []string{
		"###############",
		"#...#...#.....#",
		"#.#.#.#.#.###.#",
		"#S#...#.#.#...#",
		"#######.#.#.###",
		"#######.#.#...#",
		"#######.#.###.#",
		"###..E#...#...#",
		"###.#######.###",
		"#...###...#...#",
		"#.#####.#.###.#",
		"#.#...#.#.#...#",
		"#.#.#.#.#.#.###",
		"#...#...#...###",
		"###############",
	}

	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
		fmt.Println("Part 1:", PartOne(lines, 100))
		fmt.Println("Part 2:", PartTwo(lines, 100))
	} else {
		fmt.Println("Part 1:", PartOne(lines, 20))
		fmt.Println("Part 2:", PartTwo(lines, 50))
	}
}
