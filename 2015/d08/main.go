package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"strconv"
)

func PartOne(lines []string) (count int) {
	for _, line := range lines {
		unquoted, err := strconv.Unquote(line)
		if err != nil {
			panic(err)
		}
		count += len(line) - len(unquoted)
	}
	return count
}

func PartTwo(lines []string) (count int) {
	for _, line := range lines {
		quoted := strconv.Quote(line)
		count += len(quoted) - len(line)
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
