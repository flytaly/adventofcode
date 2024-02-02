package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"regexp"
)

func containABBA(s string) bool {
	if len(s) < 4 {
		return false
	}
	for i := 0; i < len(s)-3; i += 1 {
		if s[i] != s[i+1] && s[i] == s[i+3] && s[i+1] == s[i+2] {
			return true
		}
	}
	return false
}

func supportTLS(s string) bool {
	getEnd := func(s string, inBrackets bool) (end int) {
		re := regexp.MustCompile(`\[|\]`)
		loc := re.FindStringIndex(s)
		if loc == nil {
			return len(s)
		}
		return loc[0]
	}

	hasAbba := false
	inBrackets := false
	for end := 0; ; s = s[end+1:] {
		end = getEnd(s, inBrackets)
		part := s[0:end]
		if !inBrackets && containABBA(part) {
			hasAbba = true
		}
		if inBrackets && containABBA(part) {
			return false
		}
		if end >= len(s) {
			return hasAbba
		}
		inBrackets = !inBrackets
	}
}

func PartOne(lines []string) (count int) {
	for _, v := range lines {
		if supportTLS(v) {
			count++
		}
	}
	return count
}

func main() {
	lines := []string{}
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}
	fmt.Println("PartOne: ", PartOne(lines))
}
