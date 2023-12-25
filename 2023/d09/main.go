package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
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

func parse(lines []string) (result [][]int) {
	for _, line := range lines {
		vals := strings.Split(line, " ")
		nums := []int{}
		for _, v := range vals {
			num, _ := strconv.Atoi(v)
			nums = append(nums, num)
		}
		result = append(result, nums)

	}
	return result
}

type Dir int

const (
	End Dir = iota
	Start
)

func extrapolate(nums []int, dir Dir) int {
	next := make([]int, len(nums)-1)

	allZero := true
	for i, n := range nums[1:] {
		next[i] = n - nums[i]
		if next[i] != 0 {
			allZero = false
		}
	}
	if allZero {
		return 0
	}

	nextValue := extrapolate(next, dir)

	if dir == End {
		return next[len(next)-1] + nextValue
	}

	return next[0] - nextValue
}

func PartOne(lines []string) {
	numRows := parse(lines)
	result := 0
	for _, nums := range numRows {
		result += extrapolate(nums, End) + nums[len(nums)-1]
	}
	fmt.Println("Part 1:", result)
}

func PartTwo(lines []string) {
	numRows := parse(lines)
	result := 0
	for _, nums := range numRows {
		result += nums[0] - extrapolate(nums, Start)
	}
	fmt.Println("Part 2:", result)
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
