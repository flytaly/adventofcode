package main

import (
	"aoc/utils"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("p1", func(t *testing.T) {
		tests := utils.TestSuite[[]string, int]{
			{[]string{
				`""`,
				`"abc"`,
				`"aaa\"aaa"`,
				`"\x27"`,
			}, 12},
		}
		tests.Run(t, PartOne)
	})

	t.Run("p2", func(t *testing.T) {
		tests := utils.TestSuite[[]string, int]{
			{[]string{
				`""`,
				`"abc"`,
				`"aaa\"aaa"`,
				`"\x27"`,
			}, 19},
		}
		tests.Run(t, PartTwo)
	})
}
