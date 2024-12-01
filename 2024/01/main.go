package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"regexp"
	"slices"
)

func splitLists(lists []string) ([]int, []int) {
	a := make([]int, len(lists))
	b := make([]int, len(lists))

	re := regexp.MustCompile(`\s+`)

	for i, v := range lists {
		split := re.Split(v, -1)
		nums := utils.ToInts(split)
		a[i], b[i] = nums[0], nums[1]
	}

	return a, b
}

func PartOne(lists []string) {
	a, b := splitLists(lists)

	slices.Sort(a)
	slices.Sort(b)

	res := 0
	for i := range a {
		res += utils.Abs(a[i] - b[i])
	}

	fmt.Println("Part 1:", res)
}

func PartTwo(lists []string) {
	a, b := splitLists(lists)

	res := 0
	for _, v1 := range a {
		count := 0
		for _, v2 := range b {
			if v1 == v2 {
				count++
			}
		}
		res += v1 * count
	}

	fmt.Println("Part 2:", res)
}

func main() {
	lines := []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	PartOne(lines)
	PartTwo(lines)
}
