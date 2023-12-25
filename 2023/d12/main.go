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

func parse(lines []string) ([]string, [][]int) {
	rows := []string{}
	numRows := [][]int{}
	for _, line := range lines {
		res := strings.Split(line, " ")
		rows = append(rows, res[0])
		nums := []int{}
		for _, n := range strings.Split(res[1], ",") {
			num, _ := strconv.Atoi(n)
			nums = append(nums, num)
		}
		numRows = append(numRows, nums)
	}
	return rows, numRows
}

func BtoI(b bool) int {
	if b {
		return 1
	}
	return 0
}

var memo = map[string]int{}

func count(row string, nums []int) (res int) {
	if len(row) == 0 {
		return BtoI(len(nums) == 0)
	}
	if len(nums) == 0 {
		return BtoI(!strings.Contains(row, "#"))
	}

	id := fmt.Sprintf("%v%v", row, nums)
	if v, has := memo[id]; has {
		return v
	}

	if row[0] == '.' || row[0] == '?' {
		res += count(row[1:], nums)
	}

	if row[0] == '#' || row[0] == '?' {
		n := nums[0]

		if n > len(row) || strings.Contains(row[:n], ".") || (n < len(row) && row[n] == '#') {
			return res
		}

		if n == len(row) && !strings.Contains(row, ".") {
			return res + BtoI(len(nums) == 1)
		}

		res += count(row[n+1:], nums[1:])
	}

	memo[id] = res
	return res
}

func PartOne(lines []string) {
	rows, numRows := parse(lines)

	total := 0
	for i := 0; i < len(rows); i++ {
		total += count(rows[i], numRows[i])
	}

	fmt.Println("Part 1:", total)
}

func unfoldRow(row string, n int) (res string) {
	return strings.TrimSuffix(strings.Repeat(row+"?", n), "?")
}

func unfoldNums(nums []int, n int) (res []int) {
	for i := 0; i < n; i++ {
		res = append(res, nums...)
	}
	return res
}

func PartTwo(lines []string) {
	rows, numRows := parse(lines)

	total := 0
	for i := 0; i < len(rows); i++ {
		r := unfoldRow(rows[i], 5)
		nn := unfoldNums(numRows[i], 5)
		total += count(r, nn)
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
