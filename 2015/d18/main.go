package main

import (
	"aoc/utils"
	"fmt"
	"os"
)

type LightMap map[complex128]bool

func ItoC(row, col int) complex128 {
	return complex(float64(row), float64(col))
}

func parser(lines []string) LightMap {
	lights := make(LightMap)
	for i, line := range lines {
		for j, v := range line {
			if v != '#' {
				continue
			}
			lights[ItoC(i, j)] = true
		}
	}

	return lights
}

func printMap(lights LightMap, n, m int) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if lights[ItoC(i, j)] {
				fmt.Print("#")
				continue
			}
			fmt.Print(".")
		}
		fmt.Println()
	}
}

func neighbors(pos complex128, lights LightMap) int {
	count := 0
	for _, d := range []complex128{
		-1 - 1i, // lt
		-1 + 0,  //t
		-1 + 1i, // rt
		0 + 1i,  // r
		1 + 1i,  // rb
		1 + 0i,  // b
		1 - 1i,  // lb
		0 - 1i,  // l
	} {
		if lights[pos+d] {
			count++
		}
	}
	return count
}

func nextStep(lights LightMap, n, m int) LightMap {
	next := make(LightMap)
	for row := 0; row < n; row++ {
		for col := 0; col < m; col++ {
			pos := ItoC(row, col)
			n := neighbors(pos, lights)
			if !lights[pos] {
				next[pos] = n == 3
				continue
			}
			next[pos] = n == 2 || n == 3
		}
	}
	return next
}

func countTrue(lights LightMap) int {
	result := 0
	for _, v := range lights {
		if v {
			result++
		}
	}
	return result
}

func PartOne(lines []string, steps int) int {
	n, m := len(lines), len(lines[0])
	lights := parser(lines)

	for i := 0; i < steps; i++ {
		lights = nextStep(lights, n, m)
	}

	return countTrue(lights)
}

func PartTwo(lines []string, steps int) int {
	n, m := len(lines), len(lines[0])
	lights := parser(lines)
	setCorners := func() {
		lights[ItoC(0, 0)] = true
		lights[ItoC(0, m-1)] = true
		lights[ItoC(n-1, 0)] = true
		lights[ItoC(n-1, m-1)] = true
	}

	setCorners()
	for i := 0; i < steps; i++ {
		lights = nextStep(lights, n, m)
		setCorners()
	}

	return countTrue(lights)
}

func main() {
	lines := []string{
		".#.#.#",
		"...##.",
		"#....#",
		"..#...",
		"#.#..#",
		"####..",
	}
	steps := 5
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
		steps = 100
	}
	fmt.Println("PartOne: ", PartOne(lines, steps))
	fmt.Println("PartTwo: ", PartTwo(lines, steps))
}
