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

func parse(lines []string) map[complex128]string {
	grid := map[complex128]string{}
	for i, row := range lines {
		for j, val := range row {
			grid[complex(float64(i), float64(j))] = string(val)
		}
	}

	return grid
}

func clone[M map[K]V, K comparable, V any](m M) M {
	n := make(M, len(m))
	for k, v := range m {
		n[k] = v
	}
	return n
}

func PartOne(lines []string) {
	grid := parse(lines)
	n, m := float64(len(lines)), float64(len(lines[0]))
	reached := map[complex128]bool{}

	for num, v := range grid {
		if v == "S" {
			reached[num] = true
		}
	}

	set := func(coords map[complex128]bool, x complex128) {
		if real(x) >= 0 && real(x) < n && imag(x) >= 0 && imag(x) < m && grid[x] != "#" {
			coords[x] = true
		}
	}

	for i := 0; i < 64; i++ {
		next := clone(reached)
		for coord := range grid {
			if reached[coord] {
				next[coord] = false
				set(next, coord-1)
				set(next, coord+1)
				set(next, coord-1i)
				set(next, coord+1i)
			}
		}
		reached = next
	}

	count := 0
	for _, v := range reached {
		if v {
			count++
		}
	}

	fmt.Println("Part 1:", count)
}

func main() {
	var inputFile = "input.txt"
	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	}
	lines := readLines(inputFile)
	PartOne(lines)
}
