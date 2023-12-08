package main

import (
	"fmt"
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

func parseDirs(l string) (dir []int) {
	for _, d := range l {
		if d == 'L' {
			dir = append(dir, 0)
			continue
		}
		dir = append(dir, 1)
	}
	return dir
}

func parseSteps(lines []string) map[string][]string {
	steps := make(map[string][]string)
	re := regexp.MustCompile(`[A-Z]+`)
	for _, l := range lines {
		r := re.FindAllString(l, -1)
		steps[r[0]] = r[1:]
	}
	return steps
}

func PartOne(lines []string) int {
	dir := parseDirs(lines[0])
	steps := parseSteps(lines[2:])

	count := 0
	for step := "AAA"; step != "ZZZ"; count++ {
		d := dir[count%len(dir)]
		step = steps[step][d]
	}
	return count
}

func main() {
	lines := readLines()
	fmt.Println(PartOne(lines))
}
