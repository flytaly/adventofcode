package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func parse(lines []string) [][]int {
	res := [][]int{}
	re := regexp.MustCompile(`\d+`)
	for _, line := range lines {
		m := re.FindAllStringSubmatch(line, -1)
		nums := make([]int, 3)
		for i := 0; i < 3; i++ {
			nums[i], _ = strconv.Atoi(m[i][0])
		}
		res = append(res, nums)
	}
	return res
}

func isPossible(n []int) bool {
	for i := 0; i < 3; i++ {
		if n[i] >= n[(i+1)%3]+n[(i+2)%3] {
			return false
		}
	}
	return true
}

func PartOne(lines []string) (count int) {
	input := parse(lines)
	for _, nums := range input {
		if isPossible(nums) {
			count++
		}
	}
	return count
}

func PartTwo(lines []string) (count int) {
	input := parse(lines)
	for i := 0; i < len(input)-2; i += 3 {
		for j := 0; j < 3; j++ {
			triangle := []int{input[i][j], input[i+1][j], input[i+2][j]}
			if isPossible(triangle) {
				count++
			}
		}

	}
	return count
}

func main() {
	lines := []string{}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("PartOne: ", PartOne(lines))
	fmt.Println("PartTwo: ", PartTwo(lines))
}
