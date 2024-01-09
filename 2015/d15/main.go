package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"regexp"
)

func parser(lines []string) [][]int {
	re := regexp.MustCompile(`\w+: \w+ (-?\d+), \w+ (-?\d+), \w+ (-?\d+), \w+ (-?\d+), \w+ (-?\d+)`)
	res := [][]int{}
	for _, line := range lines {
		m := re.FindAllStringSubmatch(line, 1)
		nums := utils.ToInts(m[0][1:])
		res = append(res, nums)

	}
	return res
}

func calories(a, b, c, d int, data [][]int) bool {
	total := a*data[0][4] + b*data[1][4] + c*data[2][4] + d*data[3][4]
	return total == 500
}

func countScore(data [][]int, withCalories bool) int {
	//just brute force
	res := 0
	for a := 0; a <= 100; a++ {
		for b := 0; b <= 100; b++ {
			if a+b > 100 {
				continue
			}
			for c := 0; c <= 100; c++ {
				if a+b+c > 100 {
					continue
				}
				for d := 0; d <= 100; d++ {
					if a+b+c+d != 100 {
						continue
					}
					if withCalories && !calories(a, b, c, d, data) {
						continue
					}
					s1 := max(0, a*data[0][0]+b*data[1][0]+c*data[2][0]+d*data[3][0])
					s2 := max(0, a*data[0][1]+b*data[1][1]+c*data[2][1]+d*data[3][1])
					s3 := max(0, a*data[0][2]+b*data[1][2]+c*data[2][2]+d*data[3][2])
					s4 := max(0, a*data[0][3]+b*data[1][3]+c*data[2][3]+d*data[3][3])
					prod := s1 * s2 * s3 * s4
					if prod > res {
						res = prod
					}
				}
			}
		}
	}

	return res
}

func PartOne(lines []string) int {
	return countScore(parser(lines), false)
}

func PartTwo(lines []string) int {
	return countScore(parser(lines), true)
}

func main() {
	lines := []string{}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	if len(lines) == 0 {
		return
	}
	fmt.Println("PartOne: ", PartOne(lines))
	fmt.Println("PartTwo: ", PartTwo(lines))
}
