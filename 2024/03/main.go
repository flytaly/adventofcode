package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"regexp"
)

func PartOne(lines []string) {
	result := 0

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	for _, l := range lines {
		matches := re.FindAllStringSubmatch(l, -1)
		for _, m := range matches {
			nums := utils.ToInts(m[1:])
			result += nums[0] * nums[1]
		}
	}

	fmt.Println("Part 1:", result)
}

func PartTwo(lines []string) {
	result := 0
	isEnabled := true
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	for _, l := range lines {
		matches := re.FindAllStringSubmatch(l, -1)
		for _, m := range matches {
			switch m[0] {
			case "don't()":
				isEnabled = false
			case "do()":
				isEnabled = true
			default:
				if isEnabled {
					nums := utils.ToInts(m[1:])
					result += nums[0] * nums[1]
				}
			}
		}
	}

	fmt.Println("Part 2:", result)
}

func main() {
	lines := []string{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"}

	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}

	PartOne(lines)
	PartTwo(lines)
}
