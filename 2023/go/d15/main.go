package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func readLines(inputFile string) []string {
	_, filename, _, _ := runtime.Caller(0)
	file := filepath.Join(path.Dir(filename), inputFile)
	f, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	input := strings.Trim(string(f), "\n")
	return strings.Split(input, "\n")
}

func PartOne(lines []string) {
	total := 0
	for _, step := range strings.Split(lines[0], ",") {
		current := 0
		for _, v := range step {
			current += int(v)
			current *= 17
			current %= 256
		}
		total += current
	}

	fmt.Println("Part 1:", total)
}

func main() {
	var inputFile = "input.txt"
	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	}
	lines := readLines(inputFile)
	PartOne(lines)
}
