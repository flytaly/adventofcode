package main

import (
	"aoc/utils"
	"fmt"
	"os"
)


func PartOne(input string) (count int) {
	dirs := map[rune]int{'(': 1, ')': -1}
	for _, p := range input {
		count += dirs[p]
	}
	return count
}

func PartTwo(input string) int {
	dirs := map[rune]int{'(': 1, ')': -1}
	count := 0
	for i, p := range input {
		count += dirs[p]
		if count < 0 {
			return i + 1
		}
	}
	return 0
}

func main() {
	lines := []string{"(())"}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("PartOne: ", PartOne(lines[0]))
	fmt.Println("PartTwo: ", PartTwo(lines[0]))
}
