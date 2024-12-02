package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"regexp"
)

var re = regexp.MustCompile(`\s+`)

func parseLine(line string) []int {
	split := re.Split(line, -1)
	return utils.ToInts(split)
}

func isSafe(nums []int) bool {
	var allIncreasing bool
	for i := 1; i < len(nums); i++ {
		prev, current := nums[i-1], nums[i]

		diff := utils.Abs(current - prev)
		if diff == 0 || diff > 3 {
			return false
		}

		increasing := current > prev
		if i > 1 && allIncreasing != increasing {
			return false
		}
		allIncreasing = increasing
	}
	return true
}

func PartOne(lines []string) {
	count := 0
	for _, l := range lines {
		if isSafe(parseLine(l)) {
			count++
		}
	}
	fmt.Println("Part 1:", count)
}

func isSafe2(nums []int) bool {
	if isSafe(nums) {
		return true
	}
	fixedNums := make([]int, len(nums)-1)
	for removeIndex := 0; removeIndex < len(nums); removeIndex++ {
		pointer := 0
		for i := range fixedNums {
			if pointer == removeIndex {
				pointer++
			}
			fixedNums[i] = nums[pointer]
			pointer++
		}
		if isSafe(fixedNums) {
			return true
		}
	}

	return false
}

func PartTwo(lines []string) {
	count := 0
	for _, l := range lines {
		nums := parseLine(l)
		if isSafe2(nums) {
			count++
			continue
		}
	}
	fmt.Println("Part 2:", count)
}

func main() {
	lines := []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}

	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	PartOne(lines)
	PartTwo(lines)
}
