package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"regexp"
)

func minmax(nums []int) (int, int) {
	min := nums[0]
	max := nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

func P1(input []string) (sum int) {
	re := regexp.MustCompile(`\s`)
	for _, line := range input {
		nums := utils.ToInts(re.Split(line, -1))
		low, high := minmax(nums)
		sum += high - low
	}
	return sum
}

func main() {
	lines := []string{}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("Part1: ", P1(lines))
}
