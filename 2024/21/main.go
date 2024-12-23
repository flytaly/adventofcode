package main

import (
	"aoc/utils"
	. "aoc/utils/grid"
	"fmt"
	"image"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type PathMap map[[2]string][]string

func (p PathMap) get(from, to string) []string {
	return p[[2]string{from, to}]
}

func (p *PathMap) set(from, to string, value []string) {
	(*p)[[2]string{from, to}] = value
}

var NKeypad = []string{
	"789",
	"456",
	"123",
	"#0A",
}

var DKeypad = []string{
	"#^A",
	"<v>",
}

func shortestPaths(grid Grid[string], start, end image.Point) (paths []string) {
	type Node struct {
		p    image.Point
		path []string
	}
	type Move struct {
		d  image.Point
		kp string
	}

	if start == end {
		return []string{"A"}
	}

	bestLen := utils.Abs(start.X-end.X) + utils.Abs(start.Y-end.Y)
	queue := []Node{{p: start, path: []string{}}}
	moves := []Move{{d: ToTop, kp: "^"}, {d: ToRight, kp: ">"}, {d: ToBottom, kp: "v"}, {d: ToLeft, kp: "<"}}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for _, move := range moves {
			neighb := node.p.Add(move.d)
			if grid.At(neighb) == "#" || grid.At(neighb) == "" {
				continue
			}
			next := Node{p: neighb, path: append(slices.Clone(node.path), move.kp)}
			if len(next.path) > bestLen {
				break
			}
			if neighb.Eq(end) {
				paths = append(paths, strings.Join(next.path, "")+"A")
				continue
			}
			queue = append(queue, next)
		}
	}
	return paths
}

// all paths between two points
func getAllPaths(grid Grid[string]) PathMap {
	allPaths := PathMap{}
	for start, startVal := range grid.PointsIter() {
		for end, endVal := range grid.PointsIter() {
			if startVal == "#" || endVal == "#" {
				continue
			}
			allPaths.set(startVal, endVal, shortestPaths(grid, start, end))
		}
	}
	return allPaths
}

func pathCombinations(pathMap PathMap, code string) (result []string) {
	var parts = [][]string{}
	for i := range len(code) - 1 {
		parts = append(parts, pathMap.get(code[i:i+1], code[i+1:i+2]))
	}

	for _, prod := range utils.CartesianProduct(parts...) {
		result = append(result, strings.Join(prod, ""))
	}

	return result
}

type DeepCode struct {
	code  string
	depth int
}

func nestedLen(cache map[DeepCode]int, dirMap PathMap, c DeepCode) (length int) {
	if res, ok := cache[c]; ok {
		return res
	}

	defer func() {
		cache[c] = length
	}()

	if c.depth == 0 {
		length = len(c.code)
		return length
	}

	code := "A" + c.code
	for i := range len(code) - 1 {
		minLen := math.MaxInt
		for _, path := range dirMap.get(code[i:i+1], code[i+1:i+2]) {
			minLen = min(minLen, nestedLen(cache, dirMap, DeepCode{path, c.depth - 1}))
		}
		length += minLen
	}

	return length
}

func Solve(input []string, depth int) (count int) {
	re := regexp.MustCompile(`^\d+`)

	numPaths := getAllPaths(NewStringGrid(NKeypad))
	dirPaths := getAllPaths(NewStringGrid(DKeypad))

	for _, code := range input {
		minLen := math.MaxInt
		cache := map[DeepCode]int{}
		for _, numPath := range pathCombinations(numPaths, "A"+code) {
			minLen = min(minLen, nestedLen(cache, dirPaths, DeepCode{numPath, depth}))
		}
		num, _ := strconv.Atoi(re.FindString(code))
		count += num * minLen
	}

	return count
}

func main() {
	lines := []string{
		"029A",
		"980A",
		"179A",
		"456A",
		"379A",
	}

	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("Part 1:", Solve(lines, 2))
	fmt.Println("Part 2:", Solve(lines, 25))
}
