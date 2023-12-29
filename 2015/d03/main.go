package main

import (
	"aoc/utils"
	"fmt"
	"os"
)

var dirs = map[rune]complex128{'>': 0 + 1i, '<': 0 - 1i, '^': -1 + 0i, 'v': 1 + 0i}

func PartOne(lines []string) int {
	loc := 0 + 0i
	visited := map[complex128]int{loc: 1}
	for _, move := range lines[0] {
		loc += dirs[move]
		visited[loc] += 1
	}
	return len(visited)
}

func PartTwo(lines []string) int {
	santa, robo := 0+0i, 0+0i
	visited := map[complex128]int{santa: 1}
	for i, move := range lines[0] {
		if i%2 == 0 {
			santa += dirs[move]
			visited[santa] += 1
			continue
		}
		robo += dirs[move]
		visited[robo] += 1
	}
	return len(visited)
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
