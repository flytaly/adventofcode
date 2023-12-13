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

func isSymmetrical(pattern []string, center int) bool {
	for i := 0; i < center; i++ {
		t, b := center-i-1, center+i
		if t < 0 || b > len(pattern)-1 {
			continue
		}
		if pattern[t] != pattern[b] {
			return false
		}
	}
	return true
}

func count(pattern []string) int {
	for point := 1; point < len(pattern); point++ {
		if isSymmetrical(pattern, point) {
			return point
		}
	}
	return 0
}

func transpose(pattern []string) []string {
	rot := []string{}
	for i := 0; i < len(pattern[0]); i++ {
		row := ""
		for j := len(pattern) - 1; j >= 0; j-- {
			row += string(pattern[j][i])
		}
		rot = append(rot, row)
	}
	return rot
}

func PartOne(lines []string) {
	pattern := []string{}
	total := 0

	for i, l := range lines {
		if l != "" {
			pattern = append(pattern, l)
			if i < len(lines)-1 {
				continue
			}
		}
		value := count(pattern) * 100
		if value == 0 {
			value = count(transpose(pattern))
		}
		total += value
		pattern = []string{}
	}

	fmt.Println("Part 1:", total)
}

func count2(pattern []string, prevline int) int {
	for point := 1; point < len(pattern); point++ {
		if isSymmetrical(pattern, point) && point != prevline {
			return point
		}
	}
	return 0
}

func fixAndCount(pattern []string) int {
	swap := func(v rune) string {
		if v == '.' {
			return "#"
		}
		return "."
	}
	prevMirr := count(pattern)
	for i, l := range pattern {
		for j, v := range l {
			pattern[i] = l[:j] + swap(v) + l[j+1:]
			newMirr := count2(pattern, prevMirr)
			pattern[i] = l
			if newMirr != 0 && newMirr != prevMirr {
				return newMirr
			}
		}
	}

	return 0
}

func PartTwo(lines []string) {
	pattern := []string{}
	total := 0

	for i, l := range lines {
		if l != "" {
			pattern = append(pattern, l)
			if i < len(lines)-1 {
				continue
			}
		}
		v := fixAndCount(pattern) * 100
		if v == 0 {
			v = fixAndCount(transpose(pattern))
		}
		total += v
		pattern = []string{}
	}

	fmt.Println("Part 2:", total)
}

func main() {
	var inputFile = "input.txt"
	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	}
	lines := readLines(inputFile)
	PartOne(lines)
	PartTwo(lines)
}
