package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

const inputFile = "input.txt"

func readLines() []string {
	_, filename, _, _ := runtime.Caller(0)
	file := filepath.Join(path.Dir(filename), inputFile)
	f, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	input := strings.Trim(string(f), "\n")
	return strings.Split(input, "\n")
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverseStrings(ss []string) []string {
	r := make([]string, len(ss))
	for i, s := range ss {
		r[i] = reverse(s)
	}
	return r
}

var names = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var namesR = reverseStrings(names)

var reForward = regexp.MustCompile(strings.Join(names, "|") + "|[0-9]")
var reBack = regexp.MustCompile(strings.Join(namesR, "|") + "|[0-9]")

func convert(s string) string {
	if len(s) == 1 {
		return s
	}
	for i, name := range names {
		if s == name {
			return fmt.Sprintf("%d", i+1)
		}
	}
	return "0"
}

func calibrationValue(s string) int {
	a := reForward.FindString(s)
	b := reverse(reBack.FindString(reverse(s)))
	n, _ := strconv.Atoi(convert(a) + convert(b))
	return n
}

func main() {
	lines := readLines()
	sum := 0
	for _, l := range lines {
		sum += calibrationValue(l)
	}
	fmt.Println(sum)
}
