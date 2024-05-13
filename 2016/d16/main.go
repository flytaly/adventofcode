package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func xnor(a, b int) int {
	return ^(a ^ b) & 1
}

func toB(a string) int {
	if a == "0" {
		return 0
	}
	return 1
}

func extend(input string, size int) string {
	if len(input) >= size {
		return input[:size]
	}
	b := ""
	for i := len(input) - 1; i >= 0; i-- { // reverse + inverse
		if input[i] == '1' {
			b += "0"
			continue
		}
		b += "1"
	}
	res := input + "0" + b
	return extend(res, size)
}

func checksum(input string, strLen int) string {
	if len(input) < strLen {
		input = extend(input, strLen)
	}
	oddLen := len(input)
	for ; oddLen%2 == 0; oddLen /= 2 {
	}
	parts, partLen := []string{}, len(input)/oddLen
	for i := 0; i < len(input); i += partLen {
		parts = append(parts, input[i:i+partLen])
	}
	var sb strings.Builder
	for _, part := range parts {
		r := 1
		for i := 0; i < len(part); i += 2 {
			a, b := part[i:i+1], part[i+1:i+2]
			r ^= xnor(toB(a), toB(b))
		}
		sb.WriteString(strconv.Itoa(r))
	}

	return sb.String()
}

func main() {
	input, length := "10000", 20
	if len(os.Args) >= 2 {
		input = os.Args[1]
		length = 272
	}
	ts := time.Now()
	fmt.Printf("PartOne(%s, %d): %s [%s] \n", input, length, checksum(input, length), time.Since(ts))
}
