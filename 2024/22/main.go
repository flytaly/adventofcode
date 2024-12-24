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

func PartTwo(lines []string) (maximum int) {
	const steps int = 2000
	var sequences = map[[4]int]int{}

	// It is faster to save all the data even though it is not needed.
	// Probably, it happens because of garbage collector.
	var seen = make([]map[[4]int]struct{}, len(lines))

	for i, n := range utils.ToInts(lines) {
		diffs := make([]int, steps)
		seen[i] = make(map[[4]int]struct{}, steps)
		prev := n
		for step := 0; step < steps; step++ {
			n = (n ^ n*64) % MOD
			n = (n ^ n/32) % MOD
			n = (n ^ n*2048) % MOD
			price := n % 10
			diffs[step] = price - prev
			prev = price
			if step < 3 {
				continue
			}
			seq := [4]int(diffs[step-3 : step+1])
			if _, ok := seen[i][seq]; !ok {
				seen[i][seq] = struct{}{}
				sequences[seq] += price
			}
		}
	}

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
