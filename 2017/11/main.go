package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"strings"
)

// about cube coordinate
// https://www.redblobgames.com/grids/hexagons/#coordinates

type Cube struct {
	q, r, s int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func P1(input string) int {
	pos := Cube{0, 0, 0}
	moves := strings.Split(input, ",")

	for _, move := range moves {
		switch move {
		case "n":
			pos.r--
			pos.s++
		case "s":
			pos.r++
			pos.s--
		case "ne":
			pos.q++
			pos.r--
		case "sw":
			pos.q--
			pos.r++
		case "nw":
			pos.q--
			pos.s++
		case "se":
			pos.q++
			pos.s--
		}

	}

	return max(abs(pos.q), abs(pos.r), abs(pos.s))
}

func main() {
	lines := []string{"se,sw,se,sw,sw"}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("Part 1 =>", P1(lines[0]))
}
