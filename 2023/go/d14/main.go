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

func countLoad(lines []string) (result int) {
	for i := 0; i < len(lines[0]); i++ {
		load := len(lines)
		for j, line := range lines {
			if line[i] == 'O' {
				result += load
				load -= 1
			}
			if line[i] == '#' {
				load = len(lines) - j - 1
			}

		}
	}

	return result
}

func PartOne(lines []string) {
	fmt.Println("Part 1:", countLoad(lines))
}

func main() {
	var inputFile = "input.txt"
	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	}
	lines := readLines(inputFile)
	PartOne(lines)
}
