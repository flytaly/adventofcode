package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
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

func parseInputOne(lines []string) (pairs [][]int) {
	var re = regexp.MustCompile(`\d+`)
	times := re.FindAllString(lines[0], -1)
	dist := re.FindAllString(lines[1], -1)
	for i := range times {
		t, _ := strconv.Atoi(times[i])
		d, _ := strconv.Atoi(dist[i])
		pairs = append(pairs, []int{t, d})
	}
	return pairs
}

func PartOne(lines []string) int {
	total := 0
	for _, pair := range parseInputOne(lines) {
		time, record := pair[0], pair[1]
		count := 0
		for i := 1; i < time; i++ {
			dist := (time - i) * i
			if record < dist {
				count++
			}
		}
		if count > 0 {
			total = max(total, 1) * count
		}
	}
	return total
}

func parseInputPart2(lines []string) (pair [2]int) {
	var re = regexp.MustCompile(`\d+`)
	times := re.FindAllString(lines[0], -1)
	dist := re.FindAllString(lines[1], -1)
	pair[0], _ = strconv.Atoi(strings.Join(times, ""))
	pair[1], _ = strconv.Atoi(strings.Join(dist, ""))
	return pair
}

func PartTwo(lines []string) int {
	parsed := parseInputPart2(lines)
	time, record := parsed[0], parsed[1]
	count := 0
	for i := 1; i < time; i++ {
		dist := (time - i) * i
		if record < dist {
			count++
		}
	}
	return count
}

func main() {
	lines := readLines()
	fmt.Println(PartOne(lines))
	fmt.Println(PartTwo(lines))
}
