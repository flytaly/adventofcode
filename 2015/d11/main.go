package main

import (
	"aoc/utils"
	"fmt"
	"os"
)

func hasTwoPairs(pass string) bool {
	var pair byte = 0
	for i := 0; i < len(pass)-1; i++ {
		if pass[i] != pass[i+1] || pair == pass[i] {
			continue
		}
		if pair != 0 {
			return true
		}
		pair = pass[i]
	}
	return false
}

func hasIncreasing(pass string) bool {
	length := 1
	for i := 1; i < len(pass) && length < 3; i++ {
		if pass[i]-pass[i-1] == 1 {
			length++
			continue
		}
		length = 1
	}

	return length == 3
}

func isAcceptable(password string) bool {
	for _, ch := range password {
		if ch == 'i' || ch == 'o' || ch == 'l' {
			return false
		}
	}

	return hasIncreasing(password) && hasTwoPairs(password)
}

func nextPass(pass string) string {
	next := []rune(pass)

	for i := len(next) - 1; i >= 0; i-- {
		if next[i] == 'z' {
			next[i] = 'a'
			continue
		}
		next[i]++
		break
	}

	return string(next)
}

func nextAcceptable(input string) string {
	input = nextPass(input)
	for !isAcceptable(input) {
		input = nextPass(input)
	}
	return input
}

func main() {
	lines := []string{"ghijklmn"}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	result1 := nextAcceptable(lines[0])
	fmt.Println("PartOne: ", result1)
	fmt.Println("PartTwo: ", nextAcceptable(result1))
}
