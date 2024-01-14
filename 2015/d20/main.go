package main

import (
	"fmt"
)

func FirstFit(s []int, x int) int {
	for i, v := range s {
		if v >= x {
			return i
		}
	}
	return -1
}

func PartOne(input int) int {
	houses := make([]int, input/10)
	for i := 1; i < input/10; i++ {
		for j := i; j < input/10; j += i {
			houses[j] += i * 10
		}
	}
	return FirstFit(houses, input)
}

func PartTwo(input int) int {
	houses := make([]int, input/10-240000)
	for i := 1; i < input/11; i++ {
		for j := i; j < input/11 && j < i+i*50; j += i {
			houses[j] += i * 11
		}
	}
	return FirstFit(houses, input)
}

func main() {
	input := 29000000
	fmt.Println("PartOne: ", input, "->", PartOne(input))
	fmt.Println("PartTwo: ", input, "->", PartTwo(input))
}
