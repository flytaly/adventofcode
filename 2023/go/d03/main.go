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

type coord struct {
	n     int
	start int
	end   int
}

func extractNumbers(s string) []coord {
	var re = regexp.MustCompile(`(\d+)`)
	var numbers []coord
	for _, match := range re.FindAllStringSubmatchIndex(s, -1) {
		n, _ := strconv.Atoi(s[match[2]:match[3]])
		numbers = append(numbers, coord{n: n, start: match[2], end: match[3] - 1})
	}
	return numbers
}

func PartOne(lines []string) int {
	numbers := [][]coord{}
	re := regexp.MustCompile(`\d`)
	for _, line := range lines {
		numbers = append(numbers, extractNumbers(line))
	}

	var sumNumsInRange = func(lineIdx, colIdx1, colIdx2 int) (sum int) {
		if (lineIdx < 0) || (lineIdx >= len(numbers)) {
			return sum
		}
		for _, number := range numbers[lineIdx] {
			if (number.start <= colIdx2) && (number.end >= colIdx1) {
				sum += number.n
			}
		}
		return sum

	}

	var total int = 0
	for i, line := range lines {
		for j := 0; j < len(line); j++ {
			char := line[j]
			if char == '.' || re.MatchString(string(char)) {
				continue
			}
			total += sumNumsInRange(i-1, j-1, j+1)
			total += sumNumsInRange(i, j-1, j+1)
			total += sumNumsInRange(i+1, j-1, j+1)
		}

	}

	return total
}

func main() {
	lines := readLines()
	fmt.Println(PartOne(lines))
}
