package main

import (
	"aoc/utils"
	"fmt"
	"math"
	"os"
)

func mostCommon(dict map[byte]int) byte {
	result, num := byte(0), 0
	for k, v := range dict {
		if v > num {
			num = v
			result = k
		}
	}
	return result
}

func leastCommon(dict map[byte]int) byte {
	result, num := byte(0), math.MaxInt
	for k, v := range dict {
		if v < num {
			num = v
			result = k
		}
	}
	return result
}

func PartOne(lines []string) (result string) {
	for i := 0; i < len(lines[0]); i++ {
		dict := map[byte]int{}
		for _, line := range lines {
			dict[line[i]]++
		}
		result += string(mostCommon(dict))
	}

	return result
}

func PartTwo(lines []string) (result string) {
	for i := 0; i < len(lines[0]); i++ {
		dict := map[byte]int{}
		for _, line := range lines {
			dict[line[i]]++
		}
		result += string(leastCommon(dict))
	}

	return result
}

func main() {
	lines := []string{""}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("PartOne: ", PartOne(lines))
	fmt.Println("PartTwo: ", PartTwo(lines))
}
