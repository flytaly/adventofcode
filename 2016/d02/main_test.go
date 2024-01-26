package main

import (
	"aoc/utils"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("p1", func(t *testing.T) {
		tests := utils.TestSuite[[]string, string]{
			{[]string{
				"ULL",
				"RRDDD",
				"LURDL",
				"UUUUD",
			}, "1985"}}
		tests.Run(t, PartOne)
	})
	t.Run("p2", func(t *testing.T) {
		tests := utils.TestSuite[[]string, string]{
			{[]string{
				"ULL",
				"RRDDD",
				"LURDL",
				"UUUUD",
			}, "5DB3"}}
		tests.Run(t, PartTwo)
	})

}
