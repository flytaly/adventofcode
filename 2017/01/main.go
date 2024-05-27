package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"strconv"
)

func P1(input string) (count int) {
	for i := 0; i < len(input); i++ {
		if input[i] == input[(i+1)%len(input)] {
			n, _ := strconv.Atoi(input[i : i+1])
			count += n
		}
	}
	return count
}

func P2(input string) (count int) {
	inc := len(input) / 2
	for i := 0; i < len(input); i++ {
		if input[i] == input[(i+inc)%len(input)] {
			n, _ := strconv.Atoi(input[i : i+1])
			count += n
		}
	}
	return count
}

func main() {
	lines := []string{}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("Part1: ", P1(lines[0]))
	fmt.Println("Part2: ", P2(lines[0]))
}
