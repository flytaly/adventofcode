package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"slices"
	"strings"
)

func PartOne(lines []string) (count int) {
	for _, l := range lines {
		nums := utils.ToInts(strings.Split(l, "x"))
		w, h, l := nums[0], nums[1], nums[2]
		a, b, c := w*h, w*l, h*l
		count += 2*a + 2*b + 2*c + min(a, b, c)
	}
	return count
}

func PartTwo(lines []string) (count int) {
	for _, l := range lines {
		nums := utils.ToInts(strings.Split(l, "x"))
		slices.Sort(nums)
		w, h, l := nums[0], nums[1], nums[2]
		count += 2*w + 2*h + w*h*l
	}
	return count
}

func main() {
	lines := []string{}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("PartOne: ", PartOne(lines))
	fmt.Println("PartTwo: ", PartTwo(lines))
}
