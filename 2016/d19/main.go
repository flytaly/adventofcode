package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func P1Loop(total int) int {
	elves := []int{}
	if total%2 != 0 {
		elves = append(elves, total)
	}
	for i := 1; i < total; i += 2 {
		elves = append(elves, i)
	}

	for len(elves) > 2 {
		next := []int{}
		if len(elves)%2 != 0 {
			next = append(next, elves[len(elves)-1])
		}
		for i := 0; i < len(elves)-1; i += 2 {
			next = append(next, elves[i])
		}
		elves = next
	}

	return elves[0]
}

// The Josephus Problem
// https://www.youtube.com/watch?v=uCsD3ZGzMgE
// 1  2 3  4 5 6 7  8 9 10 11 12 13 14 15
// 1  1 3  1 3 5 7  1 3 5  7  9  11 13 15
func P1Analitic(total int) int {
	// prevous power of 2
	pow2 := math.Floor(math.Log2(float64(total)))
	reset := int(math.Pow(2, pow2))
	return (total-reset)*2 + 1
}

func main() {
	input := 7
	if len(os.Args) > 1 {
		var err error
		input, err = strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
	}
	fmt.Printf("PartOne(%d): %d\n", input, P1Analitic(input))
}
