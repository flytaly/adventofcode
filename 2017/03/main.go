package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func manhattan(a, b []int) int {
	return abs(a[0], b[0]) + abs(a[1], b[1])
}

func P1(input int) (steps int) {
	sideSize := int(math.Ceil(math.Sqrt(float64(input))))
	sideSize += (1 - sideSize%2) // make sure it's odd
	br := sideSize * sideSize    // next bottom-right corner value
	n := (sideSize - 1) / 2
	coords := []int{n, n} // br coords {1,1}, {3,3}, {5,5}...
	for i := br; ; i-- {  // move backward from bottom-right
		if i == input {
			return manhattan(coords, []int{0, 0})
		}
		if i > br-(sideSize-1) { //  left
			coords[0]--
		} else if i > br-(sideSize-1)*2 { //  up
			coords[1]--
		} else if i > br-(sideSize-1)*3 { // right
			coords[0]++
		} else if i > br-(sideSize-1)*4 { // down
			coords[1]++
		}
	}
}

func main() {
	input := 23
	if len(os.Args) > 1 {
		input, _ = strconv.Atoi(os.Args[1])
	}
	fmt.Println("input:", input)
	fmt.Println("Part 1 => ", P1(input))
}
