package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"strconv"
)

type P struct {
	x, y int
}

func nextPerm(p []int) {
	for i := len(p) - 1; i >= 0; i-- {
		if i == 0 || p[i] < len(p)-i-1 {
			p[i]++
			return
		}
		p[i] = 0
	}
}

func getPerm[T any](orig []T, p []int) []T {
	result := append([]T{}, orig...)
	for i, v := range p {
		result[i], result[i+v] = result[i+v], result[i]
	}
	return result
}

func getPoints(grid []string) []P {
	pointsMap := map[int]P{}
	for y, row := range grid {
		for x, cell := range row {
			if cell == '#' || cell == '.' {
				continue
			}
			num, _ := strconv.Atoi(string(cell))
			pointsMap[num] = P{x, y}
		}
	}
	result := make([]P, len(pointsMap))
	for i := 0; i < len(pointsMap); i++ {
		result[i] = pointsMap[i]
	}
	return result
}

// find distance from P to every point with BFS
func calcCostMap(start P, grid []string) map[P]int {
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	costMap := map[P]int{start: 0}
	frontier := []P{start}
	count := 0

	for len(frontier) > 0 {
		nextFront := []P{}
		for _, currentPoint := range frontier {
			for _, d := range dirs {
				count++
				next := P{currentPoint.x + d[0], currentPoint.y + d[1]}
				if grid[next.y][next.x] == '#' {
					continue
				}
				if _, has := costMap[next]; has {
					continue
				}
				costMap[next] = costMap[currentPoint] + 1
				nextFront = append(nextFront, next)
			}
		}
		frontier = nextFront
	}

	return costMap
}

func solve(grid []string, isPartTwo bool) int {
	points := getPoints(grid)

	paths := map[P]map[P]int{}
	for _, start := range points {
		cost := calcCostMap(start, grid)
		paths[start] = map[P]int{}
		for _, end := range points {
			paths[start][end] = cost[end]
		}
	}

	result := -1
	var lastPoint P
	for p := make([]int, len(points[1:])); p[0] < len(p); nextPerm(p) {
		current, accum := points[0], 0
		for _, next := range getPerm(points[1:], p) {
			accum += paths[current][next]
			current = next
		}
		if result < 0 || accum < result {
			result = accum
			lastPoint = current
		}
	}

	if !isPartTwo {
		return result
	}
	return result + paths[points[0]][lastPoint]
}
func main() {
	lines := []string{
		"###########",
		"#0.1.....2#",
		"#.#######.#",
		"#4.......3#",
		"###########",
	}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("PartOne: ", solve(lines, false))
	fmt.Println("PartTwo: ", solve(lines, true))
}
