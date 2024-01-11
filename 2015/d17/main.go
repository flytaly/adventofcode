package main

import (
	"aoc/utils"
	"fmt"
	"os"
)

func parser(lines []string) []int {
	return utils.ToInts(lines)
}

func combinations(values []int, expect int, minimize bool) int {
	var result int
	var combRecursive func(int, int, int)
	combRecursive = func(start int, k int, sum int) {
		if k == 0 {
			if sum == expect {
				result++
			}
			return
		}
		for i := start; i < len(values); i++ {
			nextSum := sum + values[i]
			if nextSum > expect {
				continue
			}
			combRecursive(i+1, k-1, nextSum)
		}
	}
	for r := 1; r <= len(values); r++ {
		combRecursive(0, r, 0)
		if minimize && result > 0 {
			return result
		}
	}
	return result
}

func PartOne(lines []string, total int) int {
	containers := parser(lines)
	return combinations(containers, total, false)
}

func PartTwo(lines []string, total int) int {
	containers := parser(lines)
	return combinations(containers, total, true)
}

func main() {
	lines := []string{"20", "15", "10", "5", "5"}
	total := 25
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
		total = 150
	}
	fmt.Println("PartOne: ", PartOne(lines, total))
	fmt.Println("PartTwo: ", PartTwo(lines, total))
}
