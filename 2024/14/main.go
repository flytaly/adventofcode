package main

import (
	"aoc/utils"
	"fmt"
	"image"
	"os"
	"regexp"
)

type Robot struct {
	p, v image.Point
}

func parse(lines []string) (robots []Robot) {
	re := regexp.MustCompile(`p=(\d+),(\d+)\sv=(-?\d+),(-?\d+)`)
	for _, line := range lines {
		n := utils.ToInts(re.FindAllStringSubmatch(line, 1)[0][1:])
		robots = append(robots, Robot{p: image.Pt(n[0], n[1]), v: image.Pt(n[2], n[3])})
	}
	return robots
}

func PartOne(lines []string, time, cols, rows int) int {
	robots := parse(lines)

	rect := image.Rect(0, 0, cols, rows)
	q := image.Rect(0, 0, cols/2, rows/2)
	quads := []image.Rectangle{
		q,
		q.Add(image.Pt(cols/2+1, 0)),
		q.Add(image.Pt(0, rows/2+1)),
		q.Add(image.Pt(cols/2+1, rows/2+1)),
	}
	count := make([]int, len(quads))
	for _, robo := range robots {
		robo.p = robo.p.Add(robo.v.Mul(time)).Mod(rect)
		for i, quad := range quads {
			if robo.p.In(quad) {
				count[i]++
			}
		}
	}

	prod := 1
	for _, c := range count {
		prod *= c
	}
	return prod
}

func draw(robots []Robot, cols, rows int) {
	grid := map[image.Point]bool{}
	for _, r := range robots {
		grid[r.p] = true
	}
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			c := "."
			if grid[image.Pt(i, j)] {
				c = "#"
			}
			fmt.Print(c)
		}
		fmt.Println()
	}
}

func PartTwo(lines []string) int {
	robots := parse(lines)
	cols, rows := 101, 103
	rect := image.Rect(0, 0, cols, rows)

	for time := 1; ; time++ {
		unique := make(map[image.Point]struct{})
		for i, robo := range robots {
			robots[i].p = robo.p.Add(robo.v).Mod(rect)
			unique[robots[i].p] = struct{}{}
		}
		if len(unique) == len(robots) {
			draw(robots, cols, rows)
			return time
		}
	}
}

func main() {
	lines := []string{
		"p=0,4 v=3,-3",
		"p=6,3 v=-1,-3",
		"p=10,3 v=-1,2",
		"p=2,0 v=2,-1",
		"p=0,0 v=1,3",
		"p=3,0 v=-2,-2",
		"p=7,6 v=-1,-3",
		"p=3,0 v=-1,-2",
		"p=9,3 v=2,3",
		"p=7,3 v=-1,2",
		"p=2,4 v=2,-3",
		"p=9,5 v=-3,-3",
	}

	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
	}

	// fmt.Println("Part 1: ", PartOne(lines, 100, 11, 7))
	fmt.Println("Part 1: ", PartOne(lines, 100, 101, 103))
	fmt.Println("Part 2: ", PartTwo(lines))
}
