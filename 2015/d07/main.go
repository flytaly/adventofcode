package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parse(lines []string) [][]string {
	res := make([][]string, len(lines))
	for i, line := range lines {
		line = strings.Replace(line, " -> ", " ", -1)
		res[i] = strings.Split(line, " ")
	}
	return res
}

func PartOne(lines []string) int {
	wires := map[string]uint16{}

	convert := func(ww ...string) []uint16 {
		res := []uint16{}
		for _, w := range ww {
			if i, err := strconv.Atoi(w); err == nil {
				res = append(res, uint16(i))
				continue
			}
			if _, ok := wires[w]; !ok {
				return nil
			}
			res = append(res, wires[w])
		}
		return res
	}

	ops := func(a, op, b string) (uint16, bool) {
		nums := convert(a, b)
		if nums == nil {
			return 0, false
		}
		n1, n2 := nums[0], nums[1]

		switch op {
		case "AND":
			return n1 & n2, true
		case "OR":
			return n1 | n2, true
		case "LSHIFT":
			return n1 << n2, true
		case "RSHIFT":
			return n1 >> n2, true
		}
		return 0, false
	}

	q := parse(lines)

	for _, has := wires["a"]; len(q) > 0 && !has; {
		w := q[0]
		q = q[1:]
		switch len(w) {
		case 2:
			if v := convert(w[0]); v != nil {
				wires[w[1]] = v[0]
				continue
			}
		case 3:
			if v := convert(w[1]); v != nil {
				wires[w[2]] = ^v[0]
				continue
			}
		case 4:
			if res, ok := ops(w[0], w[1], w[2]); ok {
				wires[w[3]] = res
				continue
			}
		}
		q = append(q, w)
	}
	return int(wires["a"])
}

func PartTwo(lines []string) int {
	re := regexp.MustCompile(`^(\w+) -> b$`)
	a := PartOne(lines)
	for i, line := range lines {
		if re.MatchString(line) {
			lines[i] = fmt.Sprintf("%d -> b", a)
			break
		}
	}

	return PartOne(lines)
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
