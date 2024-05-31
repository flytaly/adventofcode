package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"sort"
	"strings"
)

func isValid1(s string) bool {
	split := strings.Split(s, " ")
	seen := map[string]bool{}
	for _, word := range split {
		if seen[word] {
			return false
		}
		seen[word] = true
	}

	return true
}

func P1(input []string) (count int) {
	for _, pass := range input {
		if isValid1(pass) {
			count++
		}
	}
	return count
}

func isValid2(s string) bool {
	split := strings.Split(s, " ")
	seen := map[string]bool{}
	for _, word := range split {
		runes := []rune(word)
		sort.Slice(runes, func(i int, j int) bool { return runes[i] < runes[j] })
		word := string(runes)
		if seen[word] {
			return false
		}
		seen[word] = true
	}

	return true
}

func P2(input []string) (count int) {
	for _, pass := range input {
		if isValid2(pass) {
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
	fmt.Println("Part 1 => ", P1(lines))
	fmt.Println("Part 2 => ", P2(lines))
}
