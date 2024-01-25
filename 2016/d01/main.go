package main

import (
	"aoc/utils"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func manhattanC(x complex128) int {
	r := math.Abs(real(x))
	i := math.Abs(imag(x))
	return int(r) + int(i)
}

func Solve(input string, isPartTwo bool) int {
	instructions := strings.Split(input, ", ")

	pos := 0 + 0i
	currentDir := 0. + 1i
	visitMap := map[complex128]struct{}{}

	for _, inst := range instructions {
		switch inst[0] {
		case 'R':
			currentDir *= complex(0., -1.)
		case 'L':
			currentDir *= complex(0., 1.)
		}
		amount, _ := strconv.ParseFloat(inst[1:], 64)
		if !isPartTwo {
			pos += currentDir * complex(amount, 0.)
			continue
		}
		for i := 1; i <= int(amount); i++ {
			pos += currentDir
			if _, has := visitMap[pos]; has {
				return manhattanC(pos)
			}
			visitMap[pos] = struct{}{}
		}
	}

	return manhattanC(pos)
}

func main() {
	lines := []string{"R8, R4, R4, R8"}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("PartOne: ", Solve(lines[0], false))
	fmt.Println("PartTwo: ", Solve(lines[0], true))
}
