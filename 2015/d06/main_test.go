package main

import (
	"aoc/utils"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("p1", func(t *testing.T) {
		tests := utils.TestSuite[[]string, int]{
			{[]string{
				"turn on 0,0 through 999,999",
				"toggle 0,0 through 999,0",
				"turn off 499,499 through 500,500",
			}, 1000*1000 - 1000 - 4},
		}
		tests.Run(t, PartOne)
	})

	t.Run("p2", func(t *testing.T) {
		tests := utils.TestSuite[[]string, int]{
			{[]string{
				"turn on 0,0 through 0,0",
				"toggle 0,0 through 999,999",
			}, 2000000 + 1},
		}
		tests.Run(t, PartTwo)
	})
}
