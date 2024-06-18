package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"strings"
)

type Dir int

func (d Dir) String() string {
	return [...]string{"Up", "Right", "Down", "Left"}[d]
}

func (d Dir) turnLeft() Dir {
	return ((d-1)%4 + 4) % 4
}

func (d Dir) turnRight() Dir {
	return (d + 1) % 4
}

func (d Dir) toChar() rune {
	switch d {
	case Up:
		return '|'
	case Down:
		return '|'
	default:
		return '-'
	}
}

const (
	Up Dir = iota
	Right
	Down
	Left
)

type Pos struct {
	y, x int
}

func (p Pos) clone() Pos {
	return Pos{p.y, p.x}
}

func (p *Pos) moveTo(dir Dir) {
	switch dir {
	case Up:
		p.y--
	case Down:
		p.y++
	case Left:
		p.x--
	case Right:
		p.x++
	}

}

type Grid []string

func (g Grid) isInside(p Pos) bool {
	return p.x >= 0 && p.y >= 0 && p.y < len(g) && p.x < len(g[p.y])
}

func (g Grid) get(p Pos) rune {
	return rune(g[p.y][p.x])
}

func isLetter(c rune) bool {
	return c >= 'A' && c <= 'Z'
}

func countLetters(maze Grid) (count int) {
	for _, line := range maze {
		for _, cell := range line {
			if isLetter(cell) {
				count++
			}
		}
	}
	return count
}

func P1(maze Grid) (result string) {
	pos := Pos{0, strings.Index(maze[0], "|")}
	letters := countLetters(maze)
	for dir := Down; len(result) < letters; {
		nextPos := pos.clone()
		nextPos.moveTo(dir)
		if !maze.isInside(nextPos) {
			fmt.Println("DEAD END", string(maze.get(pos)), nextPos)
			break
		}

		pos = nextPos
		cell := maze.get(nextPos)
		if isLetter(cell) {
			result += string(cell)
		}

		if cell != '+' {
			continue
		}

		nextDir := dir.turnRight()
		nextPos.moveTo(nextDir)
		if maze.isInside(nextPos) && (maze.get(nextPos) == nextDir.toChar() || isLetter(maze.get(nextPos))) {
			dir = nextDir
			continue
		}
		dir = dir.turnLeft()
	}

	return result
}
func main() {
	lines := []string{
		"    |         ",
		"    |  +--+   ",
		"    A  |  C   ",
		"F---|----E|--+",
		"    |  |  |  D",
		"    +B-+  +--+",
	}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("Part 1 =>", P1(lines))
}
