package grid

import (
	"fmt"
	"image"
	. "image"
	"iter"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

type GridValue interface {
	string | byte | constraints.Integer | constraints.Float
}

type Grid[T GridValue] struct {
	Values                   map[Point]T
	Left, Top, Right, Bottom int
}

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
		return "", fmt.Errorf("unsupported type: %T", v)
	}
}

func (g Grid[T]) String() string {
	var sb strings.Builder

	for p := range g.PointsIter() {
		s, err := CellToString(g.Values[p])
		if err != nil {
			return "couldn't convert value to string"
		}
		sb.WriteString(s)
		end := " "
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

func (g Grid[T]) PointsIter() iter.Seq[Point] {
	return func(yield func(Point) bool) {
		for y := g.Top; y <= g.Bottom; y++ {
			for x := g.Left; x <= g.Right; x++ {
				if !yield(Pt(x, y)) {
					return
				}
			}
		}
	}
}
