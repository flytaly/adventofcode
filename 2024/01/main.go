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
	slices.Sort(a)
	slices.Sort(b)

	result := 0
	for _, value := range a {
		if index, found := slices.BinarySearch(b, value); found {
			count := 0
			for ; index < len(b) && b[index] == value; index++ {
				count++
			}
			result += value * count
		}
	}
	fmt.Println("Part 2:", result)
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
