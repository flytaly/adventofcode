package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse(lines []string) [][]int {
	result := make([][]int, len(lines))
	for i, line := range lines {
		split := strings.FieldsFunc(line, func(r rune) bool {
			return r == ' ' || r == ':'
		})
		result[i] = utils.ToInts(split)

	}
	return result
}

func isTrueEq(target int, nums []int, part2 bool) bool {
	if len(nums) == 1 {
		return nums[0] == target
	} else if nums[0] > target {
		return false
	}
	a, b := nums[0], nums[1]
	if isTrueEq(target, append([]int{a + b}, nums[2:]...), part2) {
		return true
	}
	if isTrueEq(target, append([]int{a * b}, nums[2:]...), part2) {
		return true
	}
	if !part2 {
		return false
	}
	concat, _ := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
	return isTrueEq(target, append([]int{concat}, nums[2:]...), part2)
}

func PartOne(lines []string) {
	acc := 0
	for _, nums := range parse(lines) {
		if isTrueEq(nums[0], nums[1:], false) {
			acc += nums[0]
		}
	}
	fmt.Println("Part  1:", acc)
}

func PartTwo(lines []string) {
	acc := 0
	for _, nums := range parse(lines) {
		if isTrueEq(nums[0], nums[1:], true) {
			acc += nums[0]
		}
	}
	fmt.Println("Part  2:", acc)
}

func main() {
	lines := []string{
		"190: 10 19",
		"3267: 81 40 27",
		"83: 17 5",
		"156: 15 6",
		"7290: 6 8 6 15",
		"161011: 16 10 13",
		"192: 17 8 14",
		"21037: 9 7 18 13",
		"292: 11 6 16 20",
	}

	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}

	PartOne(lines)
	PartTwo(lines)
}
