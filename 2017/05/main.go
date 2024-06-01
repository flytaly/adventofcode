package main

import (
	"aoc/utils"
	"fmt"
	"os"
)

func P1(input []string) (count int) {
	jumps := utils.ToInts(input)
	for pos := 0; pos < len(jumps); count++ {
		jump := jumps[pos]
		jumps[pos] += 1
		pos += jump
	}

	return count
}

func main() {
	lines := []string{
		"0",
		"3",
		"0",
		"1",
		"-3",
	}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("Part 1 => ", P1(lines))
}
