package main

import (
	"aoc/utils"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Inst struct {
	targetReg string
	action    string
	value     int
	condReg   string
	condSign  string
	condValue int
}

func parse(input []string) []Inst {
	conv := func(s string) int {
		value, _ := strconv.Atoi(s)
		return value
	}
	parsed := []Inst{}
	for _, line := range input {
		s := strings.Split(line, " ")
		parsed = append(parsed, Inst{s[0], s[1], conv(s[2]), s[4], s[5], conv(s[6])})
	}
	return parsed
}

func condition(regs map[string]int, inst Inst) bool {
	reg, val := inst.condReg, inst.condValue
	switch inst.condSign {
	case "==":
		return regs[reg] == val
	case "!=":
		return regs[reg] != val
	case "<":
		return regs[reg] < val
	case "<=":
		return regs[reg] <= val
	case ">":
		return regs[reg] > val
	case ">=":
		return regs[reg] >= val
	}

	return false
}

func operation(regs *map[string]int, inst Inst) {
	switch reg := inst.targetReg; inst.action {
	case "inc":
		(*regs)[reg] += inst.value
	case "dec":
		(*regs)[reg] -= inst.value
	}
}

func P1(input []string) (largest int) {
	parsed := parse(input)
	regs := map[string]int{}
	for _, instruction := range parsed {
		if condition(regs, instruction) {
			operation(&regs, instruction)
		}
	}

	largest = math.MinInt
	for _, value := range regs {
		if value > largest {
			largest = value
		}
	}

	return largest
}

func P2(input []string) (largest int) {
	parsed := parse(input)
	regs := map[string]int{}
	largest = math.MinInt

	for _, instruction := range parsed {
		if condition(regs, instruction) {
			operation(&regs, instruction)
			if v := regs[instruction.targetReg]; v > largest {
				largest = v
			}
		}
	}

	return largest
}

func main() {
	lines := []string{
		"b inc 5 if a > 1",
		"a inc 1 if b < 5",
		"c dec -10 if a >= 1",
		"c inc -20 if c == 10",
	}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("Part 1 =>", P1(lines))
	fmt.Println("Part 2 =>", P2(lines))
}
