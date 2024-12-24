package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"time"
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

func PartTwo(lines []string) (sum int) {
	const steps int = 2000
	sequences := map[[4]int]int{}

	for _, n := range utils.ToInts(lines) {
		diffs := make([]int, steps)
		seen := map[[4]int]struct{}{}
		prev := n
		for i := 0; i < steps; i++ {
			n = (n ^ n*64) % MOD
			n = (n ^ n/32) % MOD
			n = (n ^ n*2048) % MOD
			price := n % 10
			diffs[i] = price - prev
			prev = price
			if i < 3 {
				continue
			}
			seq := [4]int(diffs[i-3 : i+1])
			if _, ok := seen[seq]; !ok {
				seen[seq] = struct{}{}
				sequences[seq] += price
			}
		}
	}
	maximum := 0
	for _, v := range sequences {
		if v > maximum {
			maximum = v
		}
	}

	return maximum
}

func main() {
	lines := []string{
		"1",
		"2",
		"3",
		"2024",
	}

	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}

	fmt.Println("Part 1:", PartOne(lines))
	ts := time.Now()
	fmt.Printf("Part 2: %d [%v]\n", PartTwo(lines), time.Since(ts))
}
