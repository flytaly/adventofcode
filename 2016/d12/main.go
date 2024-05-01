package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PartOne(lines []string) int {
	regs := map[string]int{}
	for i := 0; i < len(lines); i++ {
		split := strings.Split(lines[i], " ")
		switch split[0] {
		case "cpy":
			n, err := strconv.Atoi(split[1])
			if err != nil {
				regs[split[2]] = regs[split[1]]
				continue
			}
			regs[split[2]] = n
		case "inc":
			regs[split[1]]++
		case "dec":
			regs[split[1]]--
		case "jnz":
			val, err := strconv.Atoi(split[1])
			if err == nil && val == 0 {
				continue
			}
			if err != nil && regs[split[1]] == 0 {
				continue
			}
			n, _ := strconv.Atoi(split[2])
			i += (n - 1)
		}
	}
	return regs["a"]
}

func main() {
	lines := []string{}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("PartOne: ", PartOne(lines))
}
