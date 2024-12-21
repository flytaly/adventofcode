package grid

import (
	"fmt"
	"image"
	. "image"
	"iter"
	"strconv"
	"strings"
)

type GridValue interface {
	// string | byte | constraints.Integer | constraints.Float
	any
}

type Grid[T GridValue] struct {
	Values                   map[Point]T
	Left, Top, Right, Bottom int
}

var ToTop = image.Point{0, -1}
var ToRight = image.Point{1, 0}
var ToBottom = image.Point{0, 1}
var ToLeft = image.Point{-1, 0}

var Dirs = []image.Point{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func NewGrid[T GridValue](cols, rows int) Grid[T] {
	var R, B int
	if cols > 0 {
		R = cols - 1
	}
	if rows > 0 {
		B = rows - 1
	}
	return Grid[T]{
		Values: map[Point]T{},
		Right:  R,
		Bottom: B,
	}
}

func (g Grid[T]) At(p Point) T {
	return g.Values[p]
}

func (g *Grid[T]) Set(p Point, v T) {
	g.Values[p] = v
}

func (g *Grid[t]) Span(p Point, dir image.Point, length int) []t {
	res := make([]t, length)
	for n := range length {
		res[n] = g.At(p.Add(dir.Mul(n)))
	}
	return res
}

// Neighbs returns a neighbor iterator.
// If no directions are given, it returns only direct neighbors
func (grid Grid[T]) Neighbs(p Point, dirs ...Point) iter.Seq2[Point, T] {
	if len(dirs) == 0 {
		dirs = []Point{ToTop, ToRight, ToBottom, ToLeft}
	}
	return func(yield func(Point, T) bool) {
		for _, d := range dirs {
			if neighb := p.Add(d); !yield(neighb, grid.At(neighb)) {
				return
			}
		}
	}
}

func (g *Grid[T]) Fill(value T) {
	for c := g.Left; c <= g.Right; c++ {
		for r := g.Top; r <= g.Bottom; r++ {
			g.Set(Point{c, r}, value)
		}
	}
}

func (g Grid[T]) IsInside(p Point) bool {
	return p.X >= g.Left && p.X <= g.Right && p.Y >= g.Top && p.Y <= g.Bottom
}

func CellToString(value interface{}) (string, error) {
	switch v := value.(type) {
	case int:
		return strconv.Itoa(v), nil
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 32), nil
	case string:
		return v, nil
	case byte:
		return string(v), nil
	case rune:
		return string(v), nil
	default:
		return fmt.Sprintf("%v", v), nil
	}
}

func (g Grid[T]) String() string {
	var sb strings.Builder

	for p := range g.PointsIter() {
		s, err := CellToString(g.Values[p])
		if err != nil {
			return err.Error()
		}
		sb.WriteString(s)
		end := ""
		if p.X == g.Right {
			end = "\n"
		}
		sb.WriteString(end)
	}

	return sb.String()
}

func NewStringGrid(lines []string) Grid[string] {
	grid := NewGrid[string](len(lines[0]), len(lines))
	for i, row := range lines {
		for j, col := range row {
			grid.Values[Point{j, i}] = string(col)
		}
	}
	return grid
}

func NewRuneGrid(lines []string) Grid[rune] {
	grid := NewGrid[rune](len(lines[0]), len(lines))
	for i, row := range lines {
		for j, col := range row {
			grid.Values[Point{j, i}] = col
		}
	}
	return grid
}

func (grid Grid[T]) PointsIter() iter.Seq2[Point, T] {
	return func(yield func(Point, T) bool) {
		for y := grid.Top; y <= grid.Bottom; y++ {
			for x := grid.Left; x <= grid.Right; x++ {
				point := Pt(x, y)
				if !yield(point, grid.At(point)) {
					return
				}
			}
		}
	}
}
