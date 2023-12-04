package main

import (
	"fmt"
	"math"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
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

func parseInput(line string) ([]string, []string) {
	cols := regexp.MustCompile(`[:|]`).Split(line, -1)
	var re = regexp.MustCompile(`\d+`)
	return re.FindAllString(cols[1], -1),
		re.FindAllString(cols[2], -1)
}

func intersection(a []string, b []string) (matches int) {
	for _, n := range b {
		for _, w := range a {
			if n == w {
				matches++
			}
		}
	}
	return matches
}

func PartOne(lines []string) int {
	total := 0

	for _, line := range lines {
		win, have := parseInput(line)
		matches := intersection(win, have)
		total += int(math.Pow(2, float64(matches-1)))
	}

	return total
}

func main() {
	lines := readLines()
	fmt.Println(PartOne(lines))
}
