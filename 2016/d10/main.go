package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"regexp"
)

type Destination string

const (
	Bot    Destination = "bot"
	Output Destination = "output"
)

type Op struct {
	lowDest  Destination
	lowId    int
	highDest Destination
	highId   int
}

func parse(lines []string) (chips map[int][]int, ops map[int]Op) {
	loadRe := regexp.MustCompile(`^value (\d+) goes to bot (\d+)`)
	instructionRe := regexp.MustCompile(`^bot (\d+) gives low to (bot|output) (\d+) and high to (bot|output) (\d+)`)

	chips = make(map[int][]int)
	ops = make(map[int]Op)

	for _, line := range lines {
		if loadRe.MatchString(line) {
			ints := utils.ToInts(loadRe.FindStringSubmatch(line)[1:])
			value, bot := ints[0], ints[1]
			chips[bot] = append(chips[bot], value)
		} else if instructionRe.MatchString(line) {
			match := instructionRe.FindStringSubmatch(line)
			nums := utils.ToInts(match[1:])
			bot := nums[0]
			ops[bot] = Op{
				lowDest:  Destination(match[2]),
				lowId:    nums[2],
				highDest: Destination(match[4]),
				highId:   nums[4],
			}
		}
	}
	return chips, ops
}

func Process(lines []string, num1, num2 int, isPartOne bool) int {
	botChips, ops := parse(lines)
	outputs := map[int]int{}

	check := func(bot, l, h int) (bool, int) {
		if isPartOne {
			if l == min(num1, num2) && h == max(num1, num2) {
				return true, bot
			}
			return false, 0
		}

		if outputs[0] != 0 && outputs[1] != 0 && outputs[2] != 0 {
			return true, outputs[0] * outputs[1] * outputs[2]
		}
		return false, 0
	}

	for count := 0; count != len(botChips); {
		for bot, chips := range botChips {
			if len(chips) != 2 {
				count++
				continue
			}
			op := ops[bot]
			l, h := min(chips[0], chips[1]), max(chips[0], chips[1])
			if op.lowDest == Bot && len(botChips[op.lowId]) < 2 {
				botChips[op.lowId] = append(botChips[op.lowId], l)
				botChips[bot] = []int{h}
			}
			if op.highDest == Bot && len(botChips[op.highId]) < 2 {
				botChips[op.highId] = append(botChips[op.highId], h)
				botChips[bot] = []int{l}
			}
			if op.lowDest == Output {
				outputs[op.lowId] = l
				botChips[bot] = []int{h}
			}
			if op.highDest == Output {
				outputs[op.highId] = h
				botChips[bot] = []int{l}
			}
			if isEnd, value := check(bot, l, h); isEnd {
				return value
			}
		}
	}

	return 0
}

func main() {
	lines := []string{}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("PartOne: ", Process(lines, 17, 61, true))
	fmt.Println("PartTwo: ", Process(lines, 17, 61, false))
}
