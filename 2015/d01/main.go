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
	if !filepath.IsAbs(inputFile) {
		inputFile = filepath.Join(path.Dir(filename), inputFile)
	}
	f, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	input := strings.Trim(string(f), "\n")
	return strings.Split(input, "\n")
}

func PartOne(input string) (count int) {
	dirs := map[rune]int{'(': 1, ')': -1}
	for _, p := range input {
		count += dirs[p]
	}
	return count
}

func PartTwo(input string) int {
	dirs := map[rune]int{'(': 1, ')': -1}
	count := 0
	for i, p := range input {
		count += dirs[p]
		if count < 0 {
			return i + 1
		}
	}
	return 0
}

func main() {
	lines := []string{"(())"}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = readLines(inputFile)
	}
	fmt.Println("PartOne: ", PartOne(lines[0]))
	fmt.Println("PartTwo: ", PartTwo(lines[0]))
}
