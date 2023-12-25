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

func readLines(inputFile string) []string {
	_, filename, _, _ := runtime.Caller(0)
	file := filepath.Join(path.Dir(filename), inputFile)
	f, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	input := strings.Trim(string(f), "\n")
	return strings.Split(input, "\n")
}

func hash(str string) (res int) {
	for _, v := range str {
		res += int(v)
		res *= 17
		res %= 256
	}

	return res
}

func PartOne(lines []string) {
	total := 0
	for _, step := range strings.Split(lines[0], ",") {
		total += hash(step)
	}
	fmt.Println("Part 1:", total)
}

type Lens struct {
	label string
	focal int
}

func remove(slice *[]Lens, lb string) {
	for i, v := range *slice {
		if v.label == lb {
			*slice = append((*slice)[:i], (*slice)[i+1:]...)
		}
	}
}

func add(slice *[]Lens, ln Lens) {
	for i, v := range *slice {
		if v.label == ln.label {
			(*slice)[i].focal = ln.focal
			return
		}
	}
	*slice = append(*slice, ln)
}

func PartTwo(lines []string) {
	re := regexp.MustCompile(`(\w+)(-|=)(\d+)?`)

	boxes := make([][]Lens, 256)
	for _, step := range strings.Split(lines[0], ",") {
		matches := re.FindStringSubmatch(step)
		label, operation, focal := matches[1], matches[2], 0
		if len(matches) == 4 {
			focal, _ = strconv.Atoi(matches[3])
		}
		boxId := hash(label)

		switch operation {
		case "=":
			add(&boxes[boxId], Lens{label, focal})
		case "-":
			remove(&boxes[boxId], label)
		}
	}

	total := 0
	for i, lens := range boxes {
		for j, l := range lens {
			total += (1 + i) * (j + 1) * (l.focal)
		}
	}
	fmt.Println("Part 2:", total)
}

func main() {
	var inputFile = "input.txt"
	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	}
	lines := readLines(inputFile)
	PartOne(lines)
	PartTwo(lines)
}
