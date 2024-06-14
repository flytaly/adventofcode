package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func spin[T any](p []T, n int) {
	cpy := make([]T, len(p))
	copy(cpy, p)
	for i, ch := range cpy {
		p[(i+n)%len(p)] = ch
	}
}

func swap[T any](p []T, a, b int) {
	p[a], p[b] = p[b], p[a]
}

func P1(input string, size int) (result string) {
	programs := make([]byte, size)
	for i := 0; i < size; i++ {
		programs[i] = byte('a' + i)
	}

	ops := strings.Split(input, ",")
	for _, op := range ops {
		switch op[0] {
		case 's':
			n, _ := strconv.Atoi(op[1:])
			spin(programs, n)
		case 'x':
			split := strings.Split(op[1:], "/")
			pos := utils.ToInts(split)
			swap(programs, pos[0], pos[1])
		case 'p':
			split := strings.Split(op[1:], "/")
			pos1 := slices.Index(programs, split[0][0])
			pos2 := slices.Index(programs, split[1][0])
			swap(programs, pos1, pos2)
		}
	}

	return string(programs)
}

func main() {
	lines := []string{"s1,x3/4,pe/b"}
	size := 5
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
		size = 16
	}
	fmt.Println("Part 1 =>", P1(lines[0], size))
}
