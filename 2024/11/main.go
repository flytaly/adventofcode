package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve(line string, blinks int) int {
	stones := map[string]int{}
	for _, n := range strings.Split(line, " ") {
		stones[n]++
	}

	for range blinks {
		next := map[string]int{}
		for stone, count := range stones {
			if len(stone)%2 == 0 {
				a, b := stone[:len(stone)/2], stone[len(stone)/2:]
				if b = strings.TrimLeft(b, "0"); b == "" {
					b = "0"
				}
				next[a] += count
				next[b] += count
				continue
			}
			if stone == "0" {
				next["1"] += count
				continue
			}
			num, _ := strconv.Atoi(stone)
			next[strconv.Itoa(num*2024)] += count
		}
		stones = next
	}
	// fmt.Println(stones)

	sum := 0
	for _, v := range stones {
		sum += v
	}
	return sum
}

func main() {
	lines := []string{"125 17"}

	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}

	fmt.Println("Part 1: ", Solve(lines[0], 25))
	fmt.Println("Part 2: ", Solve(lines[0], 75))
}
