package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"strings"
)

func reverseCircular(s []int, start, end int) {
	for i, j := start, end-1; i < j; i, j = i+1, j-1 {
		ii, jj := i%len(s), j%len(s)
		s[ii], s[jj] = s[jj], s[ii]
	}
}

func P1(input string, size int) int {
	positions := make([]int, size)
	for i := 0; i < size; i++ {
		positions[i] = i
	}

	nums := utils.ToInts(strings.Split(input, ","))

	var pos, skip int

	for _, section := range nums {
		reverseCircular(positions, pos, pos+section)
		pos = pos + section + skip
		skip += 1
	}

	return positions[0] * positions[1]
}

func main() {
	lines := []string{"3,4,1,5"}
	size := 5
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
		size = 256
	}
	fmt.Println("Part 1 =>", P1(lines[0], size))
}
