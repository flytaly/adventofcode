package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func P1(input []string) int {
	ops := make([][]string, len(input))

	for i, op := range input {
		ops[i] = strings.Split(op, " ")
	}

	regs := map[string]int{}

	toValue := func(numOrRegister string) int {
		if numOrRegister[0] >= ('a') && numOrRegister[0] <= ('z') {
			return regs[numOrRegister]
		}
		num, _ := strconv.Atoi(numOrRegister)
		return num
	}

	var lastPlayed int
	for i := 0; i < len(ops); i++ {
		switch op := ops[i]; op[0] {
		case "snd":
			lastPlayed = regs[op[1]]
		case "set":
			regs[op[1]] = toValue(op[2])
		case "add":
			regs[op[1]] += toValue(op[2])
		case "mul":
			regs[op[1]] *= toValue(op[2])
		case "mod":
			regs[op[1]] %= toValue(op[2])
		case "rcv":
			if toValue(op[1]) > 0 {
				return lastPlayed
			}
		case "jgz":
			if toValue(op[1]) > 0 {
				i += toValue(op[2]) - 1
			}
		}
	}

	return -1
}

func main() {
	lines := []string{
		"set a 1",
		"add a 2",
		"mul a a",
		"mod a 5",
		"snd a",
		"set a 0",
		"rcv a",
		"jgz a -1",
		"set a 1",
		"jgz a -2",
	}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("Part 1 =>", P1(lines))
}
