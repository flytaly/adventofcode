package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Registers struct {
	A int
	B int
	C int
}

func parse(lines []string) ([]int, Registers) {
	r := Registers{}
	r.A, _ = strconv.Atoi(strings.Split(lines[0], " ")[2])
	r.B, _ = strconv.Atoi(strings.Split(lines[1], " ")[2])
	r.C, _ = strconv.Atoi(strings.Split(lines[2], " ")[2])
	progStr := strings.Split(lines[4], " ")[1]
	prog := utils.ToInts(strings.Split(progStr, ","))
	return prog, r
}

func execute(prog []int, rr Registers) (string, Registers) {
	output := []string{}
	for pointer := 0; pointer < len(prog); {
		literal := prog[pointer+1]
		combo := []int{0, 1, 2, 3, rr.A, rr.B, rr.C}[literal]
		switch prog[pointer] {
		case 0:
			rr.A = rr.A >> combo
		case 1:
			rr.B ^= literal
		case 2:
			rr.B = combo % 8
		case 3:
			if rr.A != 0 {
				pointer = literal
				continue
			}
		case 4:
			rr.B ^= rr.C
		case 5:
			output = append(output, strconv.Itoa(combo%8))
		case 6:
			rr.B = rr.A >> combo
		case 7:
			rr.C = rr.A >> combo
		}
		pointer += 2
	}

	return strings.Join(output, ","), rr
}

func PartOne(lines []string) string {
	program, registers := parse(lines)
	res, _ := execute(program, registers)
	return res
}

func main() {
	lines := []string{
		"Register A: 729",
		"Register B: 0",
		"Register C: 0",
		"",
		"Program: 0,1,5,4,3,0",
	}

	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}

	fmt.Println("Part 1:", PartOne(lines))
}
