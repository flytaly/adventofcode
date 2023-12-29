package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

const MAX_VALUE = 100000000

func solve(input string, zeroes int) int {
	prefix := strings.Repeat("0", zeroes)
	for i := 0; i < MAX_VALUE; i++ {
		inp := input + strconv.Itoa(i)
		hash := md5.Sum([]byte(inp))
		if strings.HasPrefix(fmt.Sprintf("%x", hash), prefix) {
			return i
		}
	}
	return -1
}

func main() {
	// v := "abcdef"
	// fmt.Println(v, "=>", solve(v, 5))
	// v = "pqrstuv"
	// fmt.Println(v, "=>", solve(v, 5))
	v := "yzbqklnj"
	fmt.Printf("%s (5) => %d\n", v, solve(v, 5))
	fmt.Printf("%s (6) => %d\n", v, solve(v, 6))
}
