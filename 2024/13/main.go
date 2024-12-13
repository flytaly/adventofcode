package main

import (
	"aoc/utils"
	"fmt"
	"image"
	"os"
	"regexp"
)

type Point = image.Point

type Machine struct {
	A, B, Prize Point
}

func parse(lines []string) (result []Machine) {
	numsRe := regexp.MustCompile(`.*[+=](\d+).*[+=](\d+)`)
	var conv = func(line string) Point {
		n := utils.ToInts(numsRe.FindAllStringSubmatch(line, -1)[0][1:])
		return image.Pt(n[0], n[1])
	}
	for i := 0; i < len(lines); i += 4 {
		result = append(result, Machine{
			A:     conv(lines[i]),
			B:     conv(lines[i+1]),
			Prize: conv(lines[i+2]),
		})
	}
	return result
}

// a*Ax + b*Bx = Cx
// a*Ay + b*By = Cy
// a = (Cx*By - Cy*Bx) / (Ax*By - Ay*Bx)
// b = (Cx*Ay - Cy*Ax) / (Bx*Ay - By*Ax)
var A = func(a, b, c Point) int {
	return (c.X*b.Y - c.Y*b.X) / (a.X*b.Y - a.Y*b.X)
}
var B = func(a, b, c Point) int {
	return (c.X*a.Y - c.Y*a.X) / (b.X*a.Y - b.Y*a.X)
}

func Solve(lines []string, add int) (result int) {
	machines := parse(lines)
	for _, m := range machines {
		prize := m.Prize.Add(image.Pt(add, add))
		b := B(m.A, m.B, prize)
		a := A(m.A, m.B, prize)
		if m.A.Mul(a).Add(m.B.Mul(b)).Eq(prize) {
			result += a*3 + b
		}
	}
	return result
}

func main() {
	lines := []string{
		"Button A: X+94, Y+34",
		"Button B: X+22, Y+67",
		"Prize: X=8400, Y=5400",
		"",
		"Button A: X+26, Y+66",
		"Button B: X+67, Y+21",
		"Prize: X=12748, Y=12176",
		"",
		"Button A: X+17, Y+86",
		"Button B: X+84, Y+37",
		"Prize: X=7870, Y=6450",
		"",
		"Button A: X+69, Y+23",
		"Button B: X+27, Y+71",
		"Prize: X=18641, Y=10279",
	}

	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}

	fmt.Println("Part 1: ", Solve(lines, 0))
	fmt.Println("Part 2: ", Solve(lines, 10000000000000))
}
