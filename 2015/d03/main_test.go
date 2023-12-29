package main

import (
	"aoc/utils"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("p1", func(t *testing.T) {
		tests := utils.TestSuite[[]string, int]{
			{[]string{">"}, 2},
			{[]string{"^>v<"}, 4},
			{[]string{"^v^v^v^v^"}, 2},
		}
		tests.Run(t, PartOne)
	})

	t.Run("p2", func(t *testing.T) {
		tests := utils.TestSuite[[]string, int]{
			{[]string{"^v"}, 3},
			{[]string{"^>v<"}, 3},
			{[]string{"^v^v^v^v^v"}, 11},
		}
		tests.Run(t, PartTwo)
	})
}
