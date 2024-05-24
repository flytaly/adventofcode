package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toggle(instr string) string {
	split := strings.Split(instr, " ")
	if len(split) == 2 { // one-argument instruction
		if split[0] == "inc" {
			return "dec " + strings.Join(split[1:], " ")
		}
		return "inc " + strings.Join(split[1:], " ")
	}
	if split[0] == "jnz" {
		return "cpy " + strings.Join(split[1:], " ")
	}
	return "jnz " + strings.Join(split[1:], " ")
}

func execute(lines []string, regs map[string]int) int {
	instr := []string{}
	instr = append(instr, lines...)

	for i := 0; i < len(instr); i++ {
		split := strings.Split(instr[i], " ")
		// fmt.Println(instr[i], regs, instr)
		switch split[0] {
		case "cpy":
			n, err := strconv.Atoi(split[1])
			if err != nil {
				regs[split[2]] = regs[split[1]]
				continue
			}
			reg := split[2]
			if _, err := strconv.Atoi(reg); err != nil { // skip numbers (invalid instruction)
				regs[reg] = n
			}
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
			jump, err := strconv.Atoi(split[2])
			if err != nil {
				jump = regs[split[2]]
			}
			i += (jump - 1)
		case "tgl":
			val, err := strconv.Atoi(split[1])
			skip := val
			if err != nil {
				skip = regs[split[1]]
			}
			if i+skip > 0 && i+skip < len(instr) {
				instr[i+skip] = toggle(instr[i+skip])
			}
		}
	}
	return regs["a"]
}

func main() {
	lines := []string{
		"cpy 2 a",
		"tgl a",
		"tgl a",
		"tgl a",
		"cpy 1 a",
		"dec a",
		"dec a",
	}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("PartOne: ", execute(lines, map[string]int{"a": 7}))
	fmt.Println("PartTwo: ", execute(lines, map[string]int{"a": 12}))
}
