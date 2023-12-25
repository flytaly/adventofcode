package main

import (
	"fmt"
	"math"
	"os"
	"path"
	"path/filepath"
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

func toFloat(s []string) []float64 {
	res := make([]float64, len(s))
	for i, v := range s {
		res[i], _ = strconv.ParseFloat(strings.Trim(v, " "), 64)
	}
	return res
}

type hailstone struct {
	px, py, pz, vx, vy, vz float64
}

func parse(lines []string) []hailstone {
	res := []hailstone{}
	for _, line := range lines {
		line = strings.Replace(line, " @", ",", -1)
		ints := toFloat(strings.Split(line, ", "))
		h := hailstone{}
		h.px, h.py, h.pz = ints[0], ints[1], ints[2]
		h.vx, h.vy, h.vz = ints[3], ints[4], ints[5]
		res = append(res, h)
	}
	return res
}

func isInThePast(h hailstone, x, y float64) bool {
	S := math.Signbit
	if (S(h.vx) && x > h.px) || (!S(h.vx) && x < h.px) {
		return true
	}
	if (S(h.vy) && y > h.py) || (!S(h.vy) && y < h.py) {
		return true
	}
	return false
}

func isInside(x, y float64, area []float64) bool {
	return x >= area[0] && x <= area[1] && y >= area[0] && y <= area[1]
}

func PartOne(lines []string, test bool) {
	hailstones := parse(lines)
	area := []float64{7, 27}
	if !test {
		area = []float64{200000000000000, 400000000000000}
	}
	count := 0

	// y - y01 = m1(x - x01)
	// y - y02 = m2(x - x02)
	//
	// x -> -((-m1*x01+m2*x02+y01-y02)/(m1-m2))
	// y -> -((-m1*m2*x01 + m1*m2*x02 + m2*y01 - m1*y02)/(m1 - m2))
	for i := 0; i < len(hailstones)-1; i++ {
		for j := i + 1; j < len(hailstones); j++ {
			h1, h2 := hailstones[i], hailstones[j]

			m1 := h1.vy / h1.vx
			m2 := h2.vy / h2.vx

			x := -((-m1*h1.px + m2*h2.px + h1.py - h2.py) / (m1 - m2))
			y := -((-m1*m2*h1.px + m1*m2*h2.px + m2*h1.py - m1*h2.py) / (m1 - m2))

			if isInside(x, y, area) && !isInThePast(h1, x, y) && !isInThePast(h2, x, y) {
				count++
			}
		}
	}

	fmt.Println("Part 2:", count)
}

func main() {
	inputFile := "input.txt"
	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	}
	lines := readLines(inputFile)
	PartOne(lines, inputFile == "input.txt")
}
