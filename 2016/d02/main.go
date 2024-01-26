package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	y int
	x int
}

func nextPos(pos Point, dir string) Point {
	switch dir {
	case "U":
		return Point{pos.y - 1, pos.x}
	case "D":
		return Point{pos.y + 1, pos.x}
	case "L":
		return Point{pos.y, pos.x - 1}
	case "R":
		return Point{pos.y, pos.x + 1}
	}
	return pos
}

func PartOne(lines []string) string {
	keypad := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	pos := Point{1, 1}
	result := ""

	isInside := func(pos Point) bool {
		return pos.y >= 0 && pos.y < len(keypad) && pos.x >= 0 && pos.x < len(keypad[0])
	}

	for _, line := range lines {
		dirs := strings.Split(line, "")
		for _, dir := range dirs {
			next := nextPos(pos, dir)
			if isInside(next) {
				pos = next
			}
		}
		result += fmt.Sprintf("%v", keypad[pos.y][pos.x])
	}

	return result
}

func PartTwo(lines []string) string {
	keypad := [][]string{
		{"_", "_", "1", "_", "_"},
		{"_", "2", "3", "4", "_"},
		{"5", "6", "7", "8", "9"},
		{"_", "A", "B", "C", "_"},
		{"_", "_", "D", "_", "_"},
	}
	pos := Point{2, 0}
	result := ""

	isInside := func(pos Point) bool {
		if pos.y >= 0 && pos.y < len(keypad) && pos.x >= 0 && pos.x < len(keypad[0]) {
			return keypad[pos.y][pos.x] != "_"
		}
		return false
	}

	for _, line := range lines {
		dirs := strings.Split(line, "")
		for _, dir := range dirs {
			next := nextPos(pos, dir)
			if isInside(next) {
				pos = next
			}
		}
		result += fmt.Sprintf("%v", keypad[pos.y][pos.x])
	}

	return result
}

func main() {
	lines := []string{}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("PartOne: ", PartOne(lines))
	fmt.Println("PartTwo: ", PartTwo(lines))
}
