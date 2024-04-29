package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"regexp"
)

func decompressV1(str string) (count int) {
	re := regexp.MustCompile(`^\((\d+)x(\d+)\)`)

	for i := 0; i < len(str); i++ {
		if !re.MatchString(str[i:]) {
			count++
			continue
		}
		match := re.FindStringSubmatch(str[i:])
		vals := utils.ToInts(match[1:])
		size, repeat := vals[0], vals[1]
		count += size * repeat
		i += len(match[0]) + size - 1
	}

	return count
}

func decompressV2(str string) (count int) {
	re := regexp.MustCompile(`^\((\d+)x(\d+)\)`)

	for i := 0; i < len(str); i++ {
		if !re.MatchString(str[i:]) {
			count++
			continue
		}
		match := re.FindStringSubmatch(str[i:])
		vals := utils.ToInts(match[1:])
		size, repeat := vals[0], vals[1]
		startIdx := i + len(match[0])
		count += decompressV2(str[startIdx:startIdx+size]) * repeat

		i = startIdx + size - 1
	}

	return count
}

func main() {
	lines := []string{}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("PartOne: ", decompressV1(lines[0]))
	fmt.Println("PartOne: ", decompressV2(lines[0]))
}
