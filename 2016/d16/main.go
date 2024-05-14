package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func xnor(a, b byte) byte {
	return ^(a ^ b) & 1
}

// Use a slice of bytes, because string manipulation is not efficient.
func extend(input string, size int) []byte {
	res := make([]byte, size)

	for i := 0; i < len(input); i++ {
		if input[i] == '1' {
			res[i] = 1
			continue
		}
		res[i] = 0
	}
	// pointer points to the end of previous siquence
	for pointer := len(input); pointer < size; pointer = pointer*2 + 1 {
		res[pointer] = 0
		for i := 1; i <= pointer; i++ {
			if pointer+i >= size {
				return res
			}
			res[pointer+i] = 1 - res[pointer-i] // reverse and inverse bits
		}
	}
	return res
}

func checksum(input string, strLen int) string {
	inputBits := extend(input, strLen)
	oddLen := strLen
	for ; oddLen%2 == 0; oddLen /= 2 {
	}
	partLen := strLen / oddLen

	var sb strings.Builder
	zero := byte('0')
	for partStart := 0; partStart < strLen; partStart += partLen {
		r := byte(1)
		for i := partStart; i < partStart+partLen; i += 2 {
			r ^= xnor(inputBits[i], inputBits[i+1])
		}
		sb.WriteByte(zero + r)
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

	length = 35651584
	ts = time.Now()
	fmt.Printf("PartTwo(%s, %d): %s [%s] \n", input, length, checksum(input, length), time.Since(ts))
}
