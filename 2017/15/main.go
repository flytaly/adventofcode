package main

import (
	"fmt"
	"os"
	"strconv"
)

func P1(a, b int) (count int) {
	af, bf := 16807, 48271
	mask := (1 << 16) - 1 // latest 16 bits
	for i := 0; i < 40_000_000; i++ {
		a = (a * af) % 2147483647
		b = (b * bf) % 2147483647
		if a&mask == b&mask {
			count++
		}
	}

	return count
}

func main() {
	a := 65
	b := 8921
	if len(os.Args) > 2 {
		a, _ = strconv.Atoi(os.Args[1])
		b, _ = strconv.Atoi(os.Args[2])
	}
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("Part 1 =>", P1(a, b))
}
