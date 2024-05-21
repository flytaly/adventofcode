package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

var SwapPosRe = regexp.MustCompile(`swap position (\d+) with position (\d+)`)
var SwapLetterRe = regexp.MustCompile(`swap letter (\w) with letter (\w)`)
var RotateRe = regexp.MustCompile(`rotate (left|right) (\d+) step`)
var RotateBasedRe = regexp.MustCompile(`rotate based on position of letter (\w)`)
var ReverseRe = regexp.MustCompile(`reverse positions (\d+) through (\d+)`)
var MoveToRe = regexp.MustCompile(`move position (\d+) to position (\d+)`)

type Op interface {
	process(password []byte)
}

type SwapPos struct {
	a, b int
}

func (s SwapPos) process(p []byte) {
	utils.Swap(p, s.a, s.b)
}

type SwapLetter struct {
	a, b byte
}

func (s SwapLetter) process(p []byte) {
	utils.Swap(p, slices.Index(p, s.a), slices.Index(p, s.b))
}

type Reverse struct {
	a, b int
}

func (r Reverse) process(p []byte) {
	slices.Reverse(p[r.a : r.b+1])
}

type Rotate struct {
	n int
}

// positive modulo
func Mod(a, b int) int {
	return (a%b + b) % b
}

func (r Rotate) process(p []byte) {
	temp := make([]byte, len(p))
	copy(temp, p)
	for i := 0; i < len(p); i++ {
		p[i] = temp[Mod(i-r.n, len(p))]
	}
}

type RotateBased struct {
	letter byte
}

func (r RotateBased) process(p []byte) {
	times := slices.Index(p, r.letter)
	if times >= 4 {
		times++
	}
	Rotate{times + 1}.process(p)
}

type MoveTo struct {
	start, end int
}

func (m MoveTo) process(p []byte) {
	char, chIndex := p[m.end], m.end
	p[m.end] = p[m.start]
	shift := -1
	if m.start > m.end {
		shift = 1
	}
	for chIndex != m.start {
		chIndex += shift
		p[chIndex], char = char, p[chIndex]
	}
}

func parseOps(input []string) []Op {
	ops := []Op{}

	for _, op := range input {
		if SwapPosRe.MatchString(op) {
			nums := utils.ToInts(SwapPosRe.FindStringSubmatch(op)[1:])
			ops = append(ops, SwapPos{nums[0], nums[1]})
			continue
		}
		if SwapLetterRe.MatchString(op) {
			l := SwapLetterRe.FindStringSubmatch(op)[1:]
			ops = append(ops, SwapLetter{l[0][0], l[1][0]})
			continue
		}
		if RotateRe.MatchString(op) {
			match := RotateRe.FindStringSubmatch(op)[1:]
			n, _ := strconv.Atoi(match[1])
			if match[0] == "left" {
				n = -n
			}
			ops = append(ops, Rotate{n})
			continue
		}
		if RotateBasedRe.MatchString(op) {
			match := RotateBasedRe.FindStringSubmatch(op)[1:]
			ops = append(ops, RotateBased{match[0][0]})
			continue
		}
		if ReverseRe.MatchString(op) {
			nums := utils.ToInts(ReverseRe.FindStringSubmatch(op)[1:])
			ops = append(ops, Reverse{nums[0], nums[1]})
			continue
		}
		if MoveToRe.MatchString(op) {
			nums := utils.ToInts(MoveToRe.FindStringSubmatch(op)[1:])
			ops = append(ops, MoveTo{nums[0], nums[1]})
			continue
		}
	}
	return ops
}

func scramble(p []byte, ops []Op) string {
	pass := make([]byte, len(p))
	copy(pass, p)
	for _, op := range ops {
		op.process(pass)
	}
	return string(pass)
}

func P1(input []string, start string) string {
	ops := parseOps(input)
	password := []byte(start)
	return scramble(password, ops)
}

func nextPerm(p []int) {
	for i := len(p) - 1; i >= 0; i-- {
		if i == 0 || p[i] < len(p)-i-1 {
			p[i]++
			return
		}
		p[i] = 0
	}
}

func getPerm[T any](orig []T, p []int) []T {
	result := append([]T{}, orig...)
	for i, v := range p {
		result[i], result[i+v] = result[i+v], result[i]
	}
	return result
}

// just brute-force
// TODO: make proper un-scramble
func P2(input []string, start string) string {
	ops := parseOps(input)
	scrambled := []byte(start)
	for p := make([]int, len(scrambled)); p[0] < len(p); nextPerm(p) {
		testPass := getPerm(scrambled, p)
		if scramble(testPass, ops) == start {
			return string(testPass)
		}

	}
	return ""
}

func main() {
	lines := []string{
		"swap position 4 with position 0",
		"swap letter d with letter b",
		"reverse positions 0 through 4",
		"rotate left 1 step",
		"move position 1 to position 4",
		"move position 3 to position 0",
		"rotate based on position of letter b",
		"rotate based on position of letter d",
	}
	start := "abcde"
	if len(os.Args) > 1 {
		inputFile := os.Args[1]
		lines = utils.ReadLines(inputFile)
		start = "abcdefgh"
	}
	fmt.Println("PartOne: ", P1(lines, start))
	fmt.Println("PartTwo: ", P2(lines, "fbgdceah"))
}
