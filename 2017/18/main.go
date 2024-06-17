package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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

type Regs struct {
	regs map[string]int
}

func (r Regs) get(numOrRegister string) int {
	if numOrRegister[0] >= ('a') && numOrRegister[0] <= ('z') {
		return r.regs[numOrRegister]
	}
	num, _ := strconv.Atoi(numOrRegister)
	return num
}

func program(ops [][]string, id int, in, out, res chan int) {
	regs := Regs{regs: map[string]int{"p": id}}
	count := 0
	for i := 0; i < len(ops); i++ {
		fmt.Println(id, ops[i])
		switch op := ops[i]; op[0] {
		case "snd":
			out <- regs.get(op[1])
			count++
		case "set":
			regs.regs[op[1]] = regs.get(op[2])
		case "add":
			regs.regs[op[1]] += regs.get(op[2])
		case "mul":
			regs.regs[op[1]] *= regs.get(op[2])
		case "mod":
			regs.regs[op[1]] %= regs.get(op[2])
		case "rcv":
			select {
			case regs.regs[op[1]] = <-in:
			case <-time.After(500 * time.Millisecond):
				res <- count
				return
			}
		case "jgz":
			if regs.get(op[1]) > 0 {
				i += regs.get(op[2]) - 1
			}
		}
	}
	res <- count
}

func P2(input []string) int {
	ops := make([][]string, len(input))

	for i, op := range input {
		ops[i] = strings.Split(op, " ")
	}

	chA, chB := make(chan int, 1000), make(chan int, 1000)
	result := make(chan int)

	go program(ops, 0, chA, chB, make(chan int))
	go program(ops, 1, chB, chA, result)

	return <-result
}

func main() {
	lines := []string{
		"snd 1",
		"snd 2",
		"snd p",
		"rcv a",
		"rcv b",
		"rcv c",
		"rcv d",
	}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("Part 1 =>", P1(lines))
	fmt.Println("Part 2 =>", P2(lines))
}
