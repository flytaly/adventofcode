package main

import (
	"aoc/utils"
	"cmp"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
)

func maxIndex[T cmp.Ordered](values []T) (index int) {
	maxVal := values[0]
	for i, v := range values {
		if v > maxVal {
			maxVal, index = v, i
		}
	}
	return index
}

func reallocate(nums *[]int, index int) {
	n := *nums
	value := n[index]
	n[index] = 0
	for i := 1; i <= value; i++ {
		n[(index+i)%len(n)] += 1
	}
}

func cycles(input []string) (int, int) {
	split := regexp.MustCompile(`\s`).Split(input[0], -1)
	nums := utils.ToInts(split)

	states := map[string]int{}

	for count := 1; ; count++ {
		index := maxIndex(nums)
		reallocate(&nums, index)
		bytes, _ := json.Marshal(nums)
		state := string(bytes)
		if _, has := states[state]; has {
			return count, count - states[state]
		}
		states[state] = count
	}
}

func main() {
	lines := []string{"0 2 7 0"}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	p1, p2 := cycles(lines)
	fmt.Println("Part 1 =>", p1)
	fmt.Println("Part 2 =>", p2)
}
