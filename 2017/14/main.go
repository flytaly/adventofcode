package main

import (
	"aoc/2017/14/knothash"
	"fmt"
	"math/bits"
	"os"
	"strconv"
)

func P1(input string) (count int) {
	for i := 0; i < 128; i++ {
		for _, ch := range knothash.Hash(input + "-" + strconv.Itoa(i)) {
			hex, _ := strconv.ParseUint(string(ch), 16, 64)
			count += bits.OnesCount64(hex)
		}
	}
	return count
}

func main() {
	input := "flqrgnkx"
	if len(os.Args) > 1 {
		input = os.Args[1]
	}
	fmt.Println("Part 1 =>", P1(input))
}
