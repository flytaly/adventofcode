package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parse(lines []string) (map[string][]string, [][]string) {
	rules := map[string][]string{}
	updates := [][]string{}
	split := slices.Index(lines, "")
	for _, line := range lines[:split] {
		vals := strings.Split(line, "|")
		rules[vals[0]] = append(rules[vals[0]], vals[1])
	}
	for _, v := range lines[split+1:] {
		updates = append(updates, strings.Split(v, ","))
	}
	return rules, updates
}

func isItersect[T comparable](a, b []T) bool {
	for _, v := range a {
		if slices.Contains(b, v) {
			return true
		}
	}
	return false
}

func isCorrectUpdate(rules map[string][]string, update []string) bool {
	for i, v := range slices.Backward(update) {
		if isItersect(rules[v], update[:i]) {
			return false
		}
	}
	return true
}

func getMiddle(s []string) int {
	num, _ := strconv.Atoi(s[(len(s)-1)/2])
	return num
}

func PartOne(lines []string) {
	rules, updates := parse(lines)
	result := 0

	for _, update := range updates {
		if isCorrectUpdate(rules, update) {
			result += getMiddle(update)
		}
	}

	fmt.Println("Part 1:", result)
}

func PartTwo(lines []string) {
	rules, updates := parse(lines)
	result := 0

	for _, update := range updates {
		if !isCorrectUpdate(rules, update) {
			slices.SortFunc(update, func(a, b string) int {
				if slices.Contains(rules[a], b) {
					return -1
				}
				return 1
			})
			result += getMiddle(update)
		}
	}

	fmt.Println("Part 2:", result)
}

func main() {
	lines := []string{
		"47|53",
		"97|13",
		"97|61",
		"97|47",
		"75|29",
		"61|13",
		"75|53",
		"29|13",
		"97|29",
		"53|29",
		"61|53",
		"97|53",
		"61|29",
		"47|13",
		"75|47",
		"97|75",
		"47|61",
		"75|61",
		"47|29",
		"75|13",
		"53|13",
		"",
		"75,47,61,53,29",
		"97,61,53,29,13",
		"75,29,13",
		"75,97,47,61,53",
		"61,13,29",
		"97,13,75,29,47",
	}

	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}

	PartOne(lines)
	PartTwo(lines)
}
