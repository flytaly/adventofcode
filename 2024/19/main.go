package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"strings"
)

func check(str string, prefixes map[byte][]string, memo map[string]int) int {
	if res, ok := memo[str]; ok {
		return res
	}
	for _, prefix := range prefixes[str[0]] {
		if after, ok := strings.CutPrefix(str, prefix); ok {
			if after == "" {
				memo[str]++
				continue
			}
			memo[str] += check(after, prefixes, memo)
		}
	}

	return memo[str]
}

func Solve(input []string) (valid int, sum int) {
	patterns := strings.Split(input[0], ", ")
	designs := append([]string{}, input[2:]...)

	prefixes := map[byte][]string{}
	for _, p := range patterns {
		prefixes[p[0]] = append(prefixes[p[0]], p)
	}

	memo := map[string]int{}
	for _, d := range designs {
		if res := check(d, prefixes, memo); res > 0 {
			valid++
			sum += res
		}
	}

	return valid, sum
}

func main() {
	lines := []string{
		"r, wr, b, g, bwu, rb, gb, br",
		"",
		"brwrr",
		"bggr",
		"gbbr",
		"rrbgbr",
		"ubwu",
		"bwurrg",
		"brgr",
		"bbrgwb",
	}

	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}

	p1, p2 := Solve(lines)
	fmt.Println("Part 1", p1)
	fmt.Println("Part 2", p2)
}
