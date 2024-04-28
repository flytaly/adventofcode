package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"regexp"
)

type Point struct {
	X int
	Y int
}

type P = Point

type Grid struct {
	points map[P]bool
	w      int
	h      int
}

func (g Grid) String() string {
	res := ""
	for y := 0; y < g.h; y++ {
		for x := 0; x < g.w; x++ {
			if g.points[P{x, y}] {
				res += "#"
			} else {
				res += "."
			}
		}
		res += "\n"
	}
	return res
}

func (g Grid) count() int {
	count := 0
	for _, v := range g.points {
		if v {
			count++
		}
	}
	return count
}

func rect(grid Grid, w, h int) {
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			grid.points[P{i, j}] = true
		}
	}
}

func shiftCol(grid Grid, col, by int) {
	column := make(map[P]bool, grid.h)
	for row := 0; row < grid.h; row++ {
		p := P{col, row}
		if grid.points[p] {
			column[P{col, (row + by) % grid.h}] = grid.points[p]
		}
	}

	for row := 0; row < grid.h; row++ {
		grid.points[P{col, row}] = column[P{col, row}]
	}
}

func shiftRow(grid Grid, row, by int) {
	rowMap := make(map[P]bool, grid.w)
	for col := 0; col < grid.w; col++ {
		p := P{col, row}
		if grid.points[p] {
			rowMap[P{(col + by) % grid.w, row}] = grid.points[p]
		}
	}

	for col := 0; col < grid.w; col++ {
		grid.points[P{col, row}] = rowMap[P{col, row}]
	}
}

type Op struct {
	op     string
	values []int
}

func parseOperations(lines []string) []Op {
	res := []Op{}
	rectRe := regexp.MustCompile(`^rect (\d+)x(\d+)$`)
	colRe := regexp.MustCompile(`^rotate column x=(\d+) by (\d+)$`)
	rowRe := regexp.MustCompile(`^rotate row y=(\d+) by (\d+)$`)
	for _, op := range lines {
		if rectRe.MatchString(op) {
			matches := rectRe.FindStringSubmatch(op)
			res = append(res, Op{op: "rect", values: utils.ToInts(matches[1:])})
		}
		if colRe.MatchString(op) {
			matches := colRe.FindStringSubmatch(op)
			res = append(res, Op{op: "column", values: utils.ToInts(matches[1:])})
		}
		if rowRe.MatchString(op) {
			matches := rowRe.FindStringSubmatch(op)
			res = append(res, Op{op: "row", values: utils.ToInts(matches[1:])})
		}
	}
	return res
}

func lightGrid(lines []string, w, h int) Grid {
	grid := Grid{points: map[P]bool{}, w: w, h: h}
	ops := parseOperations(lines)
	for _, op := range ops {
		switch op.op {
		case "rect":
			rect(grid, op.values[0], op.values[1])
		case "column":
			shiftCol(grid, op.values[0], op.values[1])
		case "row":
			shiftRow(grid, op.values[0], op.values[1])
		}
	}

	return grid
}

func PartOne(lines []string) (count int) {
	grid := lightGrid(lines, 50, 6)
	return grid.count()
}


func main() {
	lines := []string{}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("PartOne: ", PartOne(lines))
}
