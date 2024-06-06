package main

import (
	"aoc/utils"
	"fmt"
	"os"
)

func P1(input string) int {
	isGarbage := false
	nested, sum := 0, 0
	for i := 0; i < len(input); i++ {
		if input[i] == '!' {
			i++
			continue
		}

		if isGarbage {
			isGarbage = input[i] != '>'
			continue
		}

		switch input[i] {
		case '<':
			isGarbage = true
		case '{':
			nested++
		case '}':
			sum += nested
			nested--

		}

	}

	return sum
}

func main() {
	lines := []string{
		"{{<a!>},{<a!>},{<a!>},{<ab>}}",
	}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("Part 1 =>", P1(lines[0]))
}
