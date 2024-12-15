package main

import (
	"aoc/utils"
	. "aoc/utils/grid"
	"fmt"
	"image"
	"os"
)

func parse(lines []string) (grid Grid[string], start image.Point, moves []image.Point) {
	var moveMap = map[rune]image.Point{'<': ToLeft, '>': ToRight, '^': ToTop, 'v': ToBottom}
	split := len(lines)
	for i, line := range lines {
		if line == "" {
			split = i
		}
		if i > split {
			for _, m := range line {
				moves = append(moves, moveMap[m])
			}
		}
	}

	grid = NewStringGrid(lines[:split])

	for p, v := range grid.PointsIter() {
		if v == "@" {
			start = p
			break
		}
	}

	return grid, start, moves
}

func PartOne(lines []string, time, cols, rows int) (result int) {
	grid, pos, moves := parse(lines)

outer:
	for _, move := range moves {
		next := pos.Add(move)
		for check := next; ; check = check.Add(move) {
			if grid.At(check) == "#" {
				continue outer
			}
			if grid.At(check) == "." {
				grid.Set(check, "O")
				grid.Set(pos, ".")
				grid.Set(next, "@")
				pos = next
				continue outer
			}
		}
	}

	for p, v := range grid.PointsIter() {
		if v == "O" {
			result += p.Y*100 + p.X
		}
	}

	return result
}

func main() {
	lines := []string{
		"##########",
		"#..O..O.O#",
		"#......O.#",
		"#.OO..O.O#",
		"#..O@..O.#",
		"#O#..O...#",
		"#O..O..O.#",
		"#.OO.O.OO#",
		"#....O...#",
		"##########",
		"",
		"<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^",
		"vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v",
		"><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<",
		"<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^",
		"^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><",
		"^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^",
		">^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^",
		"<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>",
		"^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>",
		"v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^",
	}

	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}

	fmt.Println("Part 1: ", PartOne(lines, 100, 101, 103))
}
