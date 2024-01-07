package main

import (
	"aoc/utils"
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
)

type RelationMap = map[string]map[string]int

func parser(lines []string) RelationMap {
	re := regexp.MustCompile(`(\w+) would (gain|lose) (\d+) .* next to (\w+)\.`)
	rel := RelationMap{}
	for _, line := range lines {
		m := re.FindAllStringSubmatch(line, 1)
		from, to := m[0][1], m[0][4]
		value, _ := strconv.Atoi(m[0][3])
		if m[0][2] == "lose" {
			value = -value
		}
		if rel[from] == nil {
			rel[from] = map[string]int{}
		}
		rel[from][to] = value
	}
	return rel
}

func dfs(name string, relations RelationMap, visited []string) int {
	neighbs := relations[name]
	result := math.MinInt

	check := func(happ int) bool {
		if result == math.MinInt {
			return true
		}
		return happ > result
	}

	if len(neighbs) == len(visited)-1 { // add last person + first person to close the loop
		first := visited[0]
		return relations[name][first] + relations[first][name]
	}

	for target, happiness := range neighbs {
		if slices.Contains(visited, target) {
			continue
		}
		total := happiness + relations[target][name] + dfs(target, relations, append(visited, target))
		if check(total) {
			result = total
		}
	}
	if result == math.MinInt {
		return 0
	}
	return result
}

func PartOne(lines []string) int {
	relations := parser(lines)
	for name := range relations { // doesn't matter which person would be first
		return dfs(name, relations, []string{name})
	}
	return 0
}

func PartTwo(lines []string) int {
	relations := parser(lines)
	relations["Me"] = map[string]int{}
	for name := range relations {
		if name == "Me" {
			continue
		}
		relations[name]["Me"] = 0
		relations["Me"][name] = 0
	}
	name := "Me"
	return dfs(name, relations, []string{name})
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
