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

func PartOne(lines []string) (result int) {
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

func expandGrid(g Grid[string]) Grid[string] {
	grid := NewGrid[string]((g.Right+1)*2, g.Bottom+1)
	for p, v := range g.PointsIter() {
		v1, v2 := v, v
		if v == "@" {
			v2 = "."
		} else if v == "O" {
			v1, v2 = "[", "]"
		}
		grid.Set(image.Pt(p.X*2, p.Y), v1)
		grid.Set(image.Pt(p.X*2+1, p.Y), v2)
	}
	return grid
}

func PartTwo(lines []string) (result int) {
	grid1, pos, moves := parse(lines)
	grid := expandGrid(grid1)
	pos = image.Pt(pos.X*2, pos.Y)

	pair := map[string]image.Point{
		"[": image.Pt(1, 0),
		"]": image.Pt(-1, 0),
	}

	grid.Set(pos, ".")
outer:
	for _, move := range moves {
		transfer := map[image.Point]string{pos: "."}
		checkQ := []image.Point{pos.Add(move)}
		for len(checkQ) > 0 {
			check := checkQ[0]
			checkQ = checkQ[1:]
			if _, ok := transfer[check]; ok {
				continue
			}
			switch cell := grid.At(check); cell {
			case "#":
				continue outer
			case "[", "]":
				transfer[check] = cell
				matching := check.Add(pair[cell])
				transfer[matching] = grid.At(matching)
				checkQ = append(checkQ, check.Add(move), matching.Add(move))
			}
		}

		for p, v := range transfer {
			if _, ok := transfer[p.Sub(move)]; !ok {
				grid.Set(p, ".") // clear behind the first boxes
			}
			grid.Set(p.Add(move), v)
		}

		pos = pos.Add(move)
	}

	grid.Set(pos, "@")
	fmt.Println(grid)

	for p, v := range grid.PointsIter() {
		if v == "[" {
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

	fmt.Println("Part 1: ", PartOne(lines))
	fmt.Println("Part 2: ", PartTwo(lines))
}
