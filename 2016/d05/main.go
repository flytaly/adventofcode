package main

import (
	"aoc/utils"
	"crypto/md5"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAX_VALUE = 100000000

func hashWithPrefix(input string, zeroes int, start int) (int, [16]byte) {
	prefix := strings.Repeat("0", zeroes)
	for i := start; i < MAX_VALUE; i++ {
		inp := input + strconv.Itoa(i)
		hash := md5.Sum([]byte(inp))
		if strings.HasPrefix(fmt.Sprintf("%x", hash), prefix) {
			return i, hash
		}
	}
	panic("no hash found")
}

func PartOne(id string) (password string) {
	var hash [16]byte
	index := -1
	for i := 0; i < 8; i++ {
		index, hash = hashWithPrefix(id, 5, index+1)
		h16 := fmt.Sprintf("%x", hash)
		password += string(h16[5])
	}

	return password
}

func main() {
	lines := []string{""}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("PartOne: ", PartOne(lines[0]))
}
