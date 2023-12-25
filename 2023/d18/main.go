package main

import (
	"fmt"
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

func parse(lines []string) [][]string {
	res := [][]string{}
	for _, l := range lines {
		r := strings.Split(l, " ")
		res = append(res, r)
	}
	return res
}

func parseHex(hx string) (string, int) {
	dir, _ := strconv.Atoi(hx[len(hx)-2 : len(hx)-1])
	m := string("RDLU"[dir])
	n, _ := strconv.ParseInt(hx[2:len(hx)-2], 16, 0)
	return m, int(n)
}

type P struct {
	x, y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func solve(lines []string, p2 bool) int {
	point := P{0, 0}
	points := []P{point}
	moves := map[string]P{"R": {1, 0}, "L": {-1, 0}, "U": {0, 1}, "D": {0, -1}}
	b := 0 // number of boundary points

	for _, step := range parse(lines) {
		m, hx := step[0], step[2]
		n, _ := strconv.Atoi(step[1])
		if p2 {
			m, n = parseHex(hx)
		}

		b += n
		move := moves[m]
		point.x += n * move.x
		point.y += n * move.y
		points = append(points, point)
	}

	// Shoelace formula
	// https://en.wikipedia.org/wiki/Shoelace_formula
	A := 0
	for cur := 0; cur < len(points); cur++ {
		p, ln := points, len(points)
		prev, next := (ln+cur-1)%ln, (cur+1)%ln
		A += p[cur].x * (p[prev].y - p[next].y)
	}
	A = abs(A / 2)

	// This area (A) is smaller than what we need because
	// shoelace formula doesn't take into account that in our case
	// the boundary points fill whole cells.
	// But with it we can calculate interior area using Pick's theorem
	// https://en.wikipedia.org/wiki/Pick%27s_theorem
	i := A - b/2 + 1

	return i + b
}

func main() {
	var inputFile = "input.txt"
	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	}
	lines := readLines(inputFile)
	fmt.Println("Part 1:", solve(lines, false))
	fmt.Println("Part 2:", solve(lines, true))
}
