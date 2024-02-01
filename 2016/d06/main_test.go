package main

import (
	"aoc/utils"
	"testing"
)

var input = []string{
	"eedadn",
	"drvtee",
	"eandsr",
	"raavrd",
	"atevrs",
	"tsrnev",
	"sdttsa",
	"rasrtv",
	"nssdts",
	"ntnada",
	"svetve",
	"tesnvt",
	"vntsnd",
	"vrdear",
	"dvrsen",
	"enarar",
}

func TestMain(t *testing.T) {
	t.Run("p1", func(t *testing.T) {
		tests := utils.TestSuite[[]string, string]{{input, "easter"}}
		tests.Run(t, PartOne)
	})
	t.Run("p2", func(t *testing.T) {
		tests := utils.TestSuite[[]string, string]{{input, "advent"}}
		tests.Run(t, PartTwo)
	})
}
