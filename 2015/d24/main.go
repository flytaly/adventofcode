package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"slices"
)

func parser(lines []string) []int {
	n := utils.ToInts(lines)
	slices.SortFunc(n, func(a, b int) int { return b - a })
	return n
}

func sum(nums []int) (sum int) {
	for _, v := range nums {
		sum += v
	}
	return sum
}

func prod(nums []int) int {
	prod := 1
	for _, v := range nums {
		prod *= v
	}
	return prod
}

// This works even without checking whether the other elements can be divided into equal groups.
// A more general solution would be more difficult.
func solve(nums []int, target int) []int {
	var result []int

	var backtrack func(nums []int, current []int, start int)
	backtrack = func(nums []int, curr []int, start int) {
		if (len(result) > 0) && (len(curr) > len(result)) {
			return
		}
		if sum(curr) == target && (len(result) != len(curr) || prod(result) > prod(curr)) {
			result = slices.Clone(curr)
		}
		for i := start; i < len(nums); i++ {
			backtrack(nums, append(curr, nums[i]), i+1)
		}
	}

	backtrack(nums, []int{}, 0)

	return result
}

func PartOne(lines []string) int {
	nums := parser(lines)
	target := sum(nums) / 3

	return prod(solve(nums, target))
}

func PartTwo(lines []string) int {
	nums := parser(lines)
	target := sum(nums) / 4

	return prod(solve(nums, target))
}

func main() {
	lines := []string{"1", "2", "3", "4", "5", "7", "8", "9", "10", "11"}

	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("PartOne: ", PartOne(lines))
	fmt.Println("PartTwo: ", PartTwo(lines))
}
