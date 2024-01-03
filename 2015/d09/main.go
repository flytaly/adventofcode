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

type Route struct {
	dest string
	dist int
}

func parser(lines []string) map[string][]Route {
	re := regexp.MustCompile(`(\w+) to (\w+) = (\d+)`)
	routes := map[string][]Route{}
	for _, line := range lines {
		m := re.FindAllStringSubmatch(line, 1)
		from, to := m[0][1], m[0][2]
		dist, _ := strconv.Atoi(m[0][3])
		routes[from] = append(routes[from], Route{to, dist})
		routes[to] = append(routes[to], Route{from, dist})
	}
	return routes
}

func dfs(name string, routes map[string][]Route, visited []string, minimize bool) int {
	neighbs := routes[name]
	result := 0

	check := func(dist int) bool {
		if result == 0 {
			return true
		}
		if !minimize {
			return dist > result
		}
		return dist < result
	}

	for _, r := range neighbs {
		if slices.Contains(visited, r.dest) {
			continue
		}
		dist := r.dist + dfs(r.dest, routes, append(visited, r.dest), minimize)
		if check(dist) {
			result = dist
		}
	}

	return result
}

func PartOne(lines []string) (distance int) {
	routes := parser(lines)
	distance = math.MaxInt

	for name := range routes {
		d := dfs(name, routes, []string{name}, true)
		if d < distance {
			distance = d
		}
	}

	return distance
}

func PartTwo(lines []string) (distance int) {
	routes := parser(lines)
	distance = 0

	for name := range routes {
		if d := dfs(name, routes, []string{name}, false); d > distance {
			distance = d
		}
	}

	return distance
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
