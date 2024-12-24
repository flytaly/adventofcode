package main

import (
	"aoc/utils"
	"fmt"
	"os"
)

const MOD int = 16777216

func PartOne(lines []string) (sum int) {
	nums := utils.ToInts(lines)
	for _, n := range nums {
		for i := 0; i < 2000; i++ {
			n = (n ^ n*64) % MOD
			n = (n ^ n/32) % MOD
			n = (n ^ n*2048) % MOD
		}
		sum += n
	}

	return sum
}

func main() {
	lines := []string{
		"1",
		"10",
		"100",
		"2024",
	}

	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}

	fmt.Println("Part 1: ", PartOne(lines))
}
