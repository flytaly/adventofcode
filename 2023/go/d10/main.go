package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func readLines(inputFile string) []string {
	_, filename, _, _ := runtime.Caller(0)
	file := filepath.Join(path.Dir(filename), inputFile)
	f, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	input := strings.Trim(string(f), "\n")
	return strings.Split(input, "\n")
}

type Dir int

const (
	L Dir = iota
	U
	R
	D
)

type Dirs map[Dir]Dir

var Moves = map[Dir][]int{
	L: {-1, 0},
	U: {0, -1},
	R: {1, 0},
	D: {0, 1},
}

var Pipes = map[byte]Dirs{
	'.': {},
	'-': {L: L, R: R},
	'|': {U: U, D: D},
	'F': {L: D, U: R},
	'J': {R: U, D: L},
	'7': {R: D, U: L},
	'L': {L: U, D: R},
	'S': {L: L, U: U, R: R, D: D},
}

func findStart(maze []string) []int {
	for i, row := range maze {
		for j, val := range row {
			if val == 'S' {
				return []int{i, j}
			}
		}
	}
	return nil
}

func move(maze []string, pos []int, dir Dir) ([]int, Dir, bool) {
	row, col := pos[0], pos[1]
	pipe := maze[row][col]
	dirs := Pipes[pipe]
	next, has := dirs[dir]
	if !has {
		return pos, dir, true
	}
	rowNext, colNext := row+Moves[next][1], col+Moves[next][0]
	if rowNext < 0 || rowNext >= len(maze) || colNext < 0 || colNext >= len(maze[0]) {
		return pos, dir, true
	}

	end := maze[rowNext][colNext] == 'S'
	return []int{rowNext, colNext}, next, end
}

func PartOne(lines []string) {
	start := findStart(lines)
	for _, dir := range []Dir{L, U, R, D} {
		pos, dist := start, 0
		for end := false; !end; dist++ {
			pos, dir, end = move(lines, pos, dir)
		}
		if dist > 1 && pos[0] == start[0] && pos[1] == start[1] {
			fmt.Println("Part 1:", dist/2)
			return
		}
	}
}

func main() {
	var inputFile = "input.txt"
	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	}
	lines := readLines(inputFile)
	PartOne(lines)
}
