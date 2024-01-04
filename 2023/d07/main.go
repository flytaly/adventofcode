package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"

	"aoc/2023/d07/p1"
	"aoc/2023/d07/p2"
)

const inputFile = "input.txt"

func readLines() []string {
	_, filename, _, _ := runtime.Caller(0)
	file := filepath.Join(path.Dir(filename), inputFile)
	f, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	input := strings.Trim(string(f), "\n")
	return strings.Split(input, "\n")
}

func parseInput(lines []string) (pairs [][]string) {
	for _, l := range lines {
		sp := strings.Split(l, " ")
		pairs = append(pairs, sp)
	}
	return pairs
}

func PartOne(lines []string) int {
	parsed := parseInput(lines)

	sort.Slice(parsed, func(i, j int) bool {
		return !p1.IsStronger(parsed[i][0], parsed[j][0])
	})

	count := 0

	for index, data := range parsed {
		bid, _ := strconv.Atoi(data[1])
		count = count + (index+1)*bid
	}

	return count
}

func PartTwo(lines []string) int {
	parsed := parseInput(lines)

	// for _, v := range parsed {
	// 	fmt.Printf("%v -> %v\n", v[0], p2.HandStrenght(v[0]))
	// }

	sort.Slice(parsed, func(i, j int) bool {
		return !p2.IsStronger(parsed[i][0], parsed[j][0])
	})

	count := 0

	for index, data := range parsed {
		bid, _ := strconv.Atoi(data[1])
		count = count + (index+1)*bid
	}

	return count
}

func main() {
	lines := readLines()
	// fmt.Println(PartOne(lines))
	fmt.Println(PartTwo(lines))
}
