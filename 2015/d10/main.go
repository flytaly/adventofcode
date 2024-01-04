package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"strings"
)

func lookAndSay(nums []int) (numsNext []int) {
	prev, count := nums[0], 0
	for _, curr := range nums {
		if prev == curr {
			count++
			continue
		}
		numsNext = append(numsNext, count, prev)
		prev = curr
		count = 1
	}
	return append(numsNext, count, prev)
}

func PartOne(line string) int {
	nums := utils.ToInts(strings.Split(line, ""))
	for i := 0; i < 40; i++ {
		nums = lookAndSay(nums)
	}
	return len(nums)
}

func PartTwo(line string) int {
	nums := utils.ToInts(strings.Split(line, ""))
	for i := 0; i < 50; i++ {
		nums = lookAndSay(nums)
	}
	return len(nums)
}

func main() {
	lines := []string{"111221"}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("PartOne: ", PartOne(lines[0]))
	fmt.Println("PartTwo: ", PartTwo(lines[0]))
}
