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

func nextA(a int) int {
	af := 16807
	for {
		a = (a * af) % 2147483647
		if a%4 == 0 {
			return a
		}
	}
}

func nextB(b int) int {
	bf := 48271
	for {
		b = (b * bf) % 2147483647
		if b%8 == 0 {
			return b
		}
	}
}

func P2(a, b int) (count int) {
	mask := (1 << 16) - 1
	for i := 1; i < 5_000_000; i++ {
		a, b = nextA(a), nextB(b)
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
	fmt.Println("Part 2 =>", P2(a, b))
}
