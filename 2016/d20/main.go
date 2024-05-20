package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Range struct {
	from int
	to   int
}

func parse(input []string) []Range {
	blocklist := []Range{}
	for _, l := range input {
		n := utils.ToInts(strings.Split(l, "-"))
		blocklist = append(blocklist, Range{from: n[0], to: n[1]})
	}

	sort.Slice(blocklist, func(i, j int) bool {
		return blocklist[i].from < blocklist[j].from
	})

	return blocklist
}

func P1(input []string) (result int) {
	for _, rng := range parse(input) {
		if rng.from > result {
			return result
		}
		if rng.to > result {
			result = rng.to + 1
		}
	}

	return result
}

func main() {
	lines := []string{"5-8", "0-2", "4-7"}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("PartOne: ", P1(lines))
	// fmt.Println("PartOne: ", P1(lines))
}
