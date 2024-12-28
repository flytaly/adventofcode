package main

import (
	"aoc/utils"
	. "aoc/utils/grid"
	"fmt"
	"os"
	"time"
)

func parse(lines []string) (locks, keys []Grid[string]) {
	locks, keys = []Grid[string]{}, []Grid[string]{}

	schema := []string{}
	for _, line := range append(lines, "") {
		if line != "" {
			schema = append(schema, line)
			continue
		}
		if schema[0][0] == '#' {
			locks = append(locks, NewStringGrid(schema))
		} else {
			keys = append(keys, NewStringGrid(schema))
		}
		schema = []string{}
	}
	return locks, keys
}

func Solve(lines []string) (result int) {
	locks, keys := parse(lines)

	for _, lock := range locks {
		for _, key := range keys {
			overlap := false
			for p := range lock.PointsIter() {
				if lock.At(p) == "#" && key.At(p) == "#" {
					overlap = true
					break
				}
			}
			if !overlap {
				result++
			}
		}
	}

	return result
}

func main() {
	inputFile := "example.txt"
	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	}

	lines := utils.ReadLines(inputFile)

	ts := time.Now()
	fmt.Printf("Part 1: %d [%v] \n", Solve(lines), time.Since(ts))
}
