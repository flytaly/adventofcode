package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"strings"
)

func parse(input []string) (map[int]int, int) {
	layers := map[int]int{}
	maxLayer := 0

	for _, line := range input {
		nums := utils.ToInts(strings.Split(line, ": "))
		index, depth := nums[0], nums[1]
		if index > maxLayer {
			maxLayer = index
		}
		layers[index] = depth
	}

	return layers, maxLayer
}

// 0 1 2 => 0 1 2 1 0
func posAtStep(n, size int) int {
	cycle := size*2 - 2
	k := n % (cycle)
	if k < size {
		return k
	}
	return cycle - k
}

func P1(input []string) (result int) {
	layers, _ := parse(input)
	for depth, layerRange := range layers {
		pico := depth
		if posAtStep(pico, layerRange) == 0 {
			result += layerRange * depth
		}
	}
	return result
}

func P2(input []string) int {
	layers, _ := parse(input)
	for pico := 0; ; pico++ {
		wasCaught := false
		for depth, layerRange := range layers {
			if posAtStep(pico+depth, layerRange) == 0 {
				wasCaught = true
				break
			}
		}
		if !wasCaught {
			return pico
		}
	}
}

func main() {
	lines := []string{
		"0: 3",
		"1: 2",
		"4: 4",
		"6: 4",
	}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("Part 1 =>", P1(lines))
	fmt.Println("Part 2 =>", P2(lines))
}
