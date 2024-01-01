package main

import (
	"aoc/utils"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("p1", func(t *testing.T) {
		tests := utils.TestSuite[[]string, int]{
			{[]string{
				"NOT y -> a",
				"x AND y -> d",
				"x LSHIFT 2 -> f",
				"x OR y -> e",
				"y RSHIFT 2 -> g",
				"123 -> x",
				"456 -> y",
				"NOT x -> h",
			}, 65079},
		}
		tests.Run(t, PartOne)
	})
}
