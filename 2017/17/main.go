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

func main() {
	input := "3"
	if len(os.Args) > 1 {
		input = os.Args[1]
	}
	step, _ := strconv.Atoi(input)
	fmt.Println("Part 1 =>", P1(step))
}
