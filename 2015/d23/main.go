package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Instr struct {
	name  string
	reg   string
	value int
}

func parser(lines []string) []Instr {
	result := make([]Instr, len(lines))
	re := regexp.MustCompile(`(\w+) (\w+)?(?:, )?([+-]\w+)?`)
	for i, line := range lines {
		match := re.FindAllStringSubmatch(line, -1)
		var value int
		if match[0][3] != "" {
			value, _ = strconv.Atoi(match[0][3])
		}
		result[i] = Instr{name: match[0][1], reg: match[0][2], value: value}
	}

	return result
}

func program(ins []Instr, regs map[string]int) map[string]int {
	for i := 0; i < len(ins); i++ {
		reg, value := ins[i].reg, ins[i].value
		switch ins[i].name {
		case "hlf":
			regs[reg] /= 2
		case "tpl":
			regs[reg] *= 3
		case "inc":
			regs[reg] += 1
		case "jmp":
			i += value - 1
		case "jie":
			if regs[reg]%2 == 0 {
				i += value - 1
			}
		case "jio":
			if regs[reg] == 1 {
				i += value - 1
			}
		}
	}

	return regs
}

func PartOne(lines []string) int {
	regs := map[string]int{"a": 0, "b": 0}
	regs = program(parser(lines), regs)
	return regs["b"]
}

func PartTwo(lines []string) int {
	regs := map[string]int{"a": 1, "b": 0}
	regs = program(parser(lines), regs)
	return regs["b"]
}

func main() {
	lines := []string{
		"inc a",
		"jio a, +2",
		"tpl a",
		"inc a",
	}

	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("PartOne: ", PartOne(lines))
	fmt.Println("PartTwo: ", PartTwo(lines))
}
