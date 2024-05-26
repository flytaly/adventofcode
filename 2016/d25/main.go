package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func execute(lines []string, regs map[string]int, limit int) string {
	instr := []string{}
	instr = append(instr, lines...)

	count := 0
	var sb strings.Builder
	for i := 0; i < len(instr); i++ {
		split := strings.Split(instr[i], " ")
		switch split[0] {
		case "cpy":
			n, err := strconv.Atoi(split[1])
			if err != nil {
				regs[split[2]] = regs[split[1]]
				continue
			}
			reg := split[2]
			if _, err := strconv.Atoi(reg); err != nil {
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
		case "out":
			count++
			if count > limit {
				return sb.String()
			}
			sb.WriteString(fmt.Sprintf("%d", regs[split[1]]))
		}
	}
	return ""
}

func find(input []string) int {
	rep := strings.Repeat("01", 15)
	fmt.Printf("searching... %s...\n", rep)
	for a := 1; a < 400; a++ {
		result := execute(input, map[string]int{"a": a}, 30)
		if result == rep {
			return a
		}
	}
	return -1
}

func main() {
	lines := []string{}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println(find(lines))
}
