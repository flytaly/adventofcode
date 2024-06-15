package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

func P1(step int) int {
	buffer := []int{0}
	current := 0
	for i := 1; i < 2018; i++ {
		current = (current+step)%len(buffer) + 1
		buffer = slices.Insert(buffer, current, i)
	}
	return buffer[current+1]
}

func P2(step int) int {
	// track only the second value
	var secondVal, pos int
	bufferLen := 1
	for i := 1; i < 50_000_000; i++ {
		pos = (pos+step)%bufferLen + 1
		bufferLen += 1
		if pos == 1 {
			secondVal = i
		}
	}
	return secondVal
}

func main() {
	input := "3"
	if len(os.Args) > 1 {
		input = os.Args[1]
	}
	step, _ := strconv.Atoi(input)
	fmt.Println("Part 1 =>", P1(step))
	fmt.Println("Part 2 =>", P2(step))
}
