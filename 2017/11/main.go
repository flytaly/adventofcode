package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"strings"
)

// about cube coordinates
// https://www.redblobgames.com/grids/hexagons/#coordinates

type Cube struct {
	q, r, s int
}

func (c *Cube) moveTo(dir string) {
	switch dir {
	case "n":
		c.r--
		c.s++
	case "s":
		c.r++
		c.s--
	case "ne":
		c.q++
		c.r--
	case "sw":
		c.q--
		c.r++
	case "nw":
		c.q--
		c.s++
	case "se":
		c.q++
		c.s--
	}
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
		pos.moveTo(move)
	}

	return max(abs(pos.q), abs(pos.r), abs(pos.s))
}

func P2(input string) int {
	pos := Cube{0, 0, 0}
	moves := strings.Split(input, ",")
	maxDist := 0

	for _, move := range moves {
		pos.moveTo(move)
		dist := max(abs(pos.q), abs(pos.r), abs(pos.s))
		if dist > maxDist {
			maxDist = dist
		}
	}

	return maxDist
}

func main() {
	lines := []string{"se,sw,se,sw,sw"}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("Part 1 =>", P1(lines[0]))
	fmt.Println("Part 2 =>", P2(lines[0]))
}
