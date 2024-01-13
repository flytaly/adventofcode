package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func parser(lines []string) map[string]string {
	res := make(map[string]string)

	for _, line := range lines {
		split := strings.Split(line, " => ")
		res[split[1]] = split[0] // swap because values are unique
	}

	return res
}

func PartOne(lines []string) int {
	replace := parser(lines[:len(lines)-2])
	text := lines[len(lines)-1]
	memo := make(map[string]struct{})

	for key, val := range replace {
		re := regexp.MustCompile(val)
		for _, match := range re.FindAllStringIndex(text, -1) {
			i := match[0]
			replaced := text[:i] + key + text[i+len(val):]
			memo[replaced] = struct{}{}
		}
	}

	return len(memo)
}

func dfs(text string, replaceMap map[string]string, count int) int {
	for key, val := range replaceMap {
		if !strings.Contains(text, key) {
			continue
		}
		replaced := strings.Replace(text, key, val, 1)
		count += 1
		res := dfs(replaced, replaceMap, count)
		if res != -1 {
			return res
		}
		count--
	}
	if text == "e" {
		return count
	}

	return -1
}

// non-deterministic solution, so should be launched multiple times
func PartTwo(lines []string) int {
	replace := parser(lines[:len(lines)-2])
	text := lines[len(lines)-1]

	return dfs(text, replace, 0)
}

func main() {
	lines := []string{
		"e => H",
		"e => O",
		"H => HO",
		"H => OH",
		"O => HH",
		"",
		"HOHOHO",
	}

	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("PartOne: ", PartOne(lines))
	fmt.Println("PartTwo: ", PartTwo(lines))
}
