package main

import (
	"aoc/utils"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strings"
)

func check(str string, prefixes map[byte][]string, memo map[string]bool) bool {
	if res, ok := memo[str]; ok {
		return res
	}
	for _, prefix := range prefixes[str[0]] {
		if after, ok := strings.CutPrefix(str, prefix); ok {
			if after == "" || check(after, prefixes, memo) {
				memo[str] = true
				return true
			}
		}
	}

	memo[str] = false
	return false
}

func PartOne(input []string) int {
	patterns := strings.Split(input[0], ", ")
	designs := append([]string{}, input[2:]...)

	prefixes := map[byte][]string{}
	for _, p := range patterns {
		prefixes[p[0]] = append(prefixes[p[0]], p)
	}

	for _, pat := range prefixes {
		slices.SortFunc(pat, func(a, b string) int {
			return cmp.Compare(len(b), len(a))
		})
	}

	count := 0
	memo := map[string]bool{}
	for _, d := range designs {
		if check(d, prefixes, memo) {
			count++
		}
	}

	return count
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

	fmt.Println("Part 1", PartOne(lines))
}
