package main

import (
	"fmt"
	"math"
	"os"
	"path"
	"path/filepath"
	"runtime"
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

type C complex128

func (c C) visit(gridCell string, coords map[C]bool, point C) {
	if gridCell == "." {
		coords[point] = true
	}
}

func (c C) toXY() (int, int) {
	return int(real(c)), int(imag(c))
}

func toC(x, y int) C {
	return C(complex(float64(x), float64(y)))
}

type Grid[T any] [][]T

func (g Grid[T]) fromC(coords C) T {
	return g[int(real(coords))][int(imag(coords))]
}

func (g Grid[T]) isInBounds(row, col int) bool {
	return row >= 0 && row < len(g) && col >= 0 && col < len(g[0])
}

func (g Grid[T]) getRepeated(row, col int) T {
	n, m := len(g), len(g[0])
	row %= n
	col %= m
	if row < 0 {
		row += n
	}
	if col < 0 {
		col += m
	}
	return g[row][col]
}

func parse(lines []string) Grid[string] {
	grid := [][]string{}
	for _, row := range lines {
		grid = append(grid, strings.Split(row, ""))
	}
	return grid
}

func clone[M map[K]V, K comparable, V any](m M) M {
	n := make(M, len(m))
	for k, v := range m {
		n[k] = v
	}
	return n
}

func step(reached map[C]bool, grid Grid[string]) map[C]bool {
	next := clone(reached)
	for coord := range reached {
		if !reached[coord] {
			continue
		}
		next[coord] = false
		for _, c := range []C{-1, 1, -1i, 1i} {
			y, x := (coord - c).toXY()
			if !grid.isInBounds(y, x) {
				continue
			}
			coord.visit(grid[y][x], next, coord-c)
		}
	}
	return next
}

func removeStart(grid [][]string) C {
	for y, row := range grid {
		for x, v := range row {
			if v == "S" {
				grid[x][y] = "."
				return C(complex(float64(y), float64(x)))
			}
		}
	}
	return C(0)
}

func countTrue[C comparable](m map[C]bool) (count int) {
	for _, v := range m {
		if v {
			count++
		}
	}
	return count
}

func PartOne(lines []string) {
	grid := parse(lines)

	start := removeStart(grid)

	reached := map[C]bool{start: true}

	for i := 0; i < 64; i++ {
		reached = step(reached, grid)
	}

	fmt.Println("Part 1:", countTrue(reached))
}

func stepInf(reached map[C]bool, grid Grid[string]) map[C]bool {
	next := clone(reached)
	for coord := range reached {
		if !reached[coord] {
			continue
		}
		next[coord] = false
		for _, c := range []C{-1, 1, -1i, 1i} {
			cn := coord - c
			y, x := int(real(cn)), int(imag(cn))
			coord.visit(grid.getRepeated(y, x), next, cn)
		}
	}
	return next
}

func PrintGrid(visited map[C]bool, grid Grid[string]) string {
	x0, y0, x1, y1 := math.MaxInt, math.MaxInt, math.MinInt, math.MinInt
	for c := range visited {
		y, x := int(real(c)), int(imag(c))
		x0, y0 = min(x, x0), min(y, y0)
		x1, y1 = max(x, x1), max(y, y1)
	}
	sb := strings.Builder{}
	for i := y0; i <= y1; i++ {
		for j := x0; j <= x1; j++ {
			if visited[C(complex(float64(i), float64(j)))] {
				sb.WriteString("O")
				continue
			}
			sb.WriteString(grid.getRepeated(i, j))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func counter(grid Grid[string], start C, steps int, countTotal bool) int {
	reached := map[C]bool{start: true}
	for i := 0; i < steps; i++ {
		reached = stepInf(reached, grid)
	}

	// fmt.Println(PrintGrid(reached, grid))
	if countTotal {
		return countTrue(reached)
	}
	res := 0
	for i, row := range grid {
		for j := range row {
			coord := C(complex(float64(i), float64(j)))
			if reached[coord] {
				res++
			}
		}
	}
	return res
}

func Pow2(n int) int {
	return n * n
}

func PartTwo(lines []string) {
	grid := parse(lines)
	start := removeStart(grid)

	steps := 26501365
	n := len(grid) // 131
	half := float64(n / 2)
	hY, hX := C(complex(half, 0)), C(complex(0., half))

	gridWidth := steps/n - 1

	count := func(start C, steps int) int {
		return counter(grid, start, steps, false)
	}

	sum := 0
	// grids fully inside the "diamond"
	oddGrids := Pow2(gridWidth/2*2+1) * count(start, n+n%2+1)
	evenGrids := Pow2((gridWidth+1)/2*2) * count(start, n+n%2)
	sum += oddGrids + evenGrids

	// corners of the "diamond"
	top := count(start+hY, n-1)
	bottom := count(start-hY, n-1)
	right := count(start-hX, n-1)
	left := count(start+hX, n-1)
	sum += top + bottom + right + left

	// diagonals
	// top right corner, starts with bottom left e.t.c.
	for _, v := range [][2]int{{n - 1, 0}, {0, 0}, {0, n - 1}, {n - 1, n - 1}} {
		coords := toC(v[0], v[1])
		sum += (gridWidth + 1) * count(coords, n/2-1) // smaller grids
		sum += (gridWidth) * count(coords, n*3/2-1)   // bigger grids
	}

	fmt.Println("Part 2:", sum)
}

func main() {
	inputFile := "input.txt"
	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	}
	lines := readLines(inputFile)
	PartOne(lines)
	PartTwo(lines)
}
