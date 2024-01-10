package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Memories map[string]int

func parser(lines []string) []Memories {
	re := regexp.MustCompile(`[:,] `)

	res := []Memories{}
	for _, line := range lines {
		split := re.Split(line, -1)
		m := Memories{}
		for i := 1; i < len(split); i = i + 2 {
			n, _ := strconv.Atoi(split[i+1])
			m[split[i]] = n
		}
		res = append(res, m)
	}
	return res
}

var things = map[string]int{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

func PartOne(lines []string) int {
	memList := parser(lines)
	for sueIndex, memories := range memList {
		fit := true
		for memory, memNum := range memories {
			if expectNum, has := things[memory]; has && expectNum != memNum {
				fit = false
			}
		}
		if fit {
			return sueIndex + 1
		}
	}
	return 0
}

func PartTwo(lines []string) int {
	memList := parser(lines)

	var test = func(memory string, num int) bool {
		switch expect := things[memory]; memory {
		case "cats":
			fallthrough
		case "trees":
			return num > expect
		case "pomeranians":
			fallthrough
		case "goldfish":
			return num < expect
		default:
			return num == expect
		}
	}

outer:
	for sueIndex, memories := range memList {
		for memory, memNum := range memories {
			if !test(memory, memNum) {
				continue outer
			}
		}
		return sueIndex + 1
	}

	return 0
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
