package main

import (
	"fmt"
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

type Dir int

const (
	U Dir = iota
	R
	D
	L
)

func (d Dir) String() string {
	switch d {
	case U:
		return "U"
	case R:
		return "R"
	case D:
		return "D"
	case L:
		return "L"
	}
	return ""
}

type Beam struct {
	row int
	col int
	dir Dir
}

func (b Beam) String() string {
	return fmt.Sprintf("{%d, %d, %s}", b.row, b.col, b.dir)
}

func (b *Beam) move() {
	switch b.dir {
	case R:
		b.col++
	case L:
		b.col--
	case U:
		b.row--
	case D:
		b.row++
	}
}

func (b Beam) isOutside(w, h int) bool {
	return b.col >= w || b.col < 0 || b.row >= h || b.row < 0
}

func (b *Beam) splitter(sp byte) *Beam {
	if (b.dir == R || b.dir == L) && sp == '-' {
		return nil
	}
	if (b.dir == U || b.dir == D) && sp == '|' {
		return nil
	}
	if sp == '|' {
		b.dir = U
		return &Beam{row: b.row, col: b.col, dir: D}
	}
	// -
	b.dir = L
	return &Beam{row: b.row, col: b.col, dir: R}
}

func (b *Beam) mirror(sp byte) {
	if (b.dir == L && sp == '/') || (b.dir == R && sp == '\\') {
		b.dir = D
		return
	}
	if (b.dir == R && sp == '/') || (b.dir == L && sp == '\\') {
		b.dir = U
		return
	}
	if (b.dir == U && sp == '/') || (b.dir == D && sp == '\\') {
		b.dir = R
		return
	}
	if (b.dir == D && sp == '/') || (b.dir == U && sp == '\\') {
		b.dir = L
		return
	}
}

func filterBeams(beams *[]*Beam, removeQueue []*Beam) {
	for _, b := range removeQueue {
		for i, b2 := range *beams {
			if b == b2 {
				*beams = append((*beams)[:i], (*beams)[i+1:]...)
				break
			}
		}
	}
}

func countEnergy(grid []string, beams []*Beam) int {
	energy := make([][][]Dir, len(grid))
	for i, v := range grid {
		energy[i] = make([][]Dir, len(v))
	}

	seen := func(b Beam) bool {
		cell := energy[b.row][b.col]
		for _, d := range cell {
			if d == b.dir {
				return true
			}
		}
		return false
	}

	w, h := len(grid[0]), len(grid)
	for i := 0; len(beams) != 0; i++ {
		deleteQueue := []*Beam{}
		for _, b := range beams {
			b.move()
			if b.isOutside(w, h) || seen(*b) {
				deleteQueue = append(deleteQueue, b)
				continue
			}

			energy[b.row][b.col] = append(energy[b.row][b.col], b.dir)

			switch cell := grid[b.row][b.col]; cell {
			case '.':
				continue
			case '|':
				fallthrough
			case '-':
				b2 := b.splitter(cell)
				if b2 != nil {
					beams = append(beams, b2)
				}
			case '/':
				fallthrough
			case '\\':
				b.mirror(cell)
			}
		}

		filterBeams(&beams, deleteQueue)

		// fmt.Printf("-> %d %s\n", i, beams)
	}

	count := 0
	for _, c := range energy {
		for _, r := range c {
			if len(r) > 0 {
				count++
			}
		}
	}

	return count

}

func PartOne(grid []string) {
	beams := []*Beam{{0, -1, R}}

	fmt.Println("Part 1:", countEnergy(grid, beams))
}

func PartTwo(grid []string) {
	res := 0

	for i := 0; i < len(grid); i++ {
		res = max(res, countEnergy(grid, []*Beam{{i, -1, R}}))
		res = max(res, countEnergy(grid, []*Beam{{i, len(grid), L}}))
	}

	for j := 0; j < len(grid[0]); j++ {
		res = max(res, countEnergy(grid, []*Beam{{-1, j, D}}))
		res = max(res, countEnergy(grid, []*Beam{{len(grid[0]), j, U}}))
	}

	fmt.Println("Part 2:", res)
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
