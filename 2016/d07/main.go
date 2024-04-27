package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"regexp"
	"strings"
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
	getEnd := func(s string) (end int) {
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
		end = getEnd(s)
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

func splitString(s string) (outside []string, inside []string) {
	re := regexp.MustCompile(`\[|\]`)
	parts := re.Split(s, -1)

	for i, v := range parts {
		if i%2 == 0 { // outside
			outside = append(outside, v)
			continue
		}
		inside = append(inside, v)
	}

	return outside, inside
}

func getNextABA(s string, index int) (string, int) {
	for i := index; i < len(s)-2; i++ {
		if s[i] != s[i+1] && s[i] == s[i+2] {
			return s[i : i+3], i
		}
	}

	return "", -1
}

func hasSubstring(s []string, sub string) bool {
	for _, v := range s {
		if strings.Contains(v, sub) {
			return true
		}
	}
	return false
}

func supportSSL(s string) bool {
	outside, inside := splitString(s)

	for _, str := range outside {
		// search for all overlapping ABAs
		for aba, pos := "", 0; pos != -1; pos++ {
			aba, pos = getNextABA(str, pos)
			if pos == -1 {
				break
			}
			// reverse aba -> bab
			bab := string([]byte{aba[1], aba[0], aba[1]})
			if hasSubstring(inside, bab) {
				return true
			}
		}
	}

	return false
}

func PartTwo(lines []string) (count int) {
	for _, v := range lines {
		if supportSSL(v) {
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
	fmt.Println("PartTwo: ", PartTwo(lines))
}
