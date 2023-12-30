package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"slices"
)

func PartOne(lines []string) (count int) {
	vowels := []rune("aeiou")
	disallowed := []string{"ab", "cd", "pq", "xy"}

	isNice := func(line string) bool {
		vowelsNum, hasTwice := 0, false
		line = line + " " //  add space to not check boundaries
		for i := 0; i < len(line)-1; i++ {
			if slices.Contains(vowels, rune(line[i])) {
				vowelsNum++
			}
			if slices.Contains(disallowed, line[i:i+2]) {
				return false
			}
			hasTwice = hasTwice || line[i] == line[i+1]
		}
		return hasTwice && vowelsNum >= 3
	}

	for _, line := range lines {
		if isNice(line) {
			count++
		}
	}

	return count
}

func PartTwo(lines []string) (count int) {
	hasPair := func(line string) bool {
		pairs := map[string]int{} // first position of the pair
		for i := 0; i < len(line)-1; i++ {
			p := line[i : i+2]
			if _, ok := pairs[p]; !ok {
				pairs[p] = i
				continue
			}
			if pairs[p] < i-1 {
				return true
			}
		}
		return false
	}

	hasRepeatBetween := func(line string) bool {
		for i := 0; i < len(line)-2; i++ {
			if line[i] == line[i+2] {
				return true
			}
		}
		return false
	}

	for _, line := range lines {
		if hasPair(line) && hasRepeatBetween(line) {
			count++
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
	fmt.Println("PartOne: ", PartOne(lines))
	fmt.Println("PartTwo: ", PartTwo(lines))
}
