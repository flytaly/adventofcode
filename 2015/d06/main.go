package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"regexp"
)

type Coord struct {
	row int
	col int
}

type instruction struct {
	op   string
	from Coord
	to   Coord
}

func parse(lines []string) []instruction {
	re := regexp.MustCompile(`(turn on|toggle|turn off) (\d+),(\d+) through (\d+),(\d+)`)
	instructions := make([]instruction, len(lines))
	for i, line := range lines {
		match := (re.FindAllStringSubmatch(line, -1))[0]
		ints := utils.ToInts(match[2:])
		instructions[i] = instruction{
			op:   match[1],
			from: Coord{col: ints[0], row: ints[1]},
			to:   Coord{col: ints[2], row: ints[3]},
		}
	}
	return instructions
}

func fill[T any](instructions []instruction, ops map[string]func(T) T) [][]T {
	grid := make([][]T, 1000)
	for _, inst := range instructions {
		from, to := inst.from, inst.to
		for i := from.col; i <= to.col; i++ {
			if len(grid[i]) == 0 {
				grid[i] = make([]T, 1000)
			}
			for j := from.row; j <= to.row; j++ {
				grid[i][j] = ops[inst.op](grid[i][j])
			}
		}
	}
	return grid
}

func countBool(grid [][]bool) (res int) {
	for _, row := range grid {
		for _, val := range row {
			if val {
				res++
			}
		}
	}
	return res
}

func PartOne(lines []string) int {
	instructions := parse(lines)
	grid := fill(
		instructions,
		map[string]func(bool) bool{
			"turn on":  func(x bool) bool { return true },
			"toggle":   func(x bool) bool { return !x },
			"turn off": func(x bool) bool { return false },
		})
	return countBool(grid)
}

func sum(grid [][]int) (res int) {
	for _, row := range grid {
		for _, val := range row {
			res += val
		}
	}
	return res
}

func PartTwo(lines []string) int {
	instructions := parse(lines)
	grid := fill(
		instructions,
		map[string]func(int) int{
			"turn on":  func(x int) int { return x + 1 },
			"toggle":   func(x int) int { return x + 2 },
			"turn off": func(x int) int { return max(0, x-1) },
		})
	return sum(grid)
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
