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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func containsInRow(row string, ch rune) bool {
	for _, v := range row {
		if v == ch {
			return true
		}
	}
	return false
}

func containsInCol(table []string, j int, ch byte) bool {
	for _, row := range table {
		if row[j] == ch {
			return true
		}
	}

	return false
}

func getEmpties(space []string) ([]int, []int) {
	rows, cols := []int{}, []int{}

	for i, row := range space {
		if !containsInRow(row, '#') {
			rows = append(rows, i)
		}
	}

	for j := range space[0] {
		if !containsInCol(space, j, '#') {
			cols = append(cols, j)
		}
	}

	return rows, cols
}

type galaxy struct {
	r int
	c int
}

func getGalaxies(space []string) []galaxy {
	res := []galaxy{}
	for c, l := range space {
		for r, v := range l {
			if v == '#' {
				res = append(res, galaxy{c, r})
			}
		}
	}
	return res
}

func getDist(space []string) int {
	p := 0

	rows, cols := getEmpties(space)

	dist := func(g1, g2 galaxy) int {
		pad := 0
		for _, row := range rows {
			if row < max(g1.r, g2.r) && row > min(g1.r, g2.r) {
				pad += 1
			}
		}
		for _, col := range cols {
			if col < max(g1.c, g2.c) && col > min(g1.c, g2.c) {
				pad += 1
			}
		}
		return abs(g1.r-g2.r) + abs(g1.c-g2.c) + pad
	}

	galaxies := getGalaxies(space)

	for i, g1 := range galaxies {
		for _, g2 := range galaxies[i+1:] {
			p += dist(g1, g2)
		}
	}

	return p
}

func PartOne(lines []string) {
	fmt.Println("Part 1:", getDist(lines))
}

func main() {
	var inputFile = "input.txt"
	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	}
	lines := readLines(inputFile)
	PartOne(lines)
}
