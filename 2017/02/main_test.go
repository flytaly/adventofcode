package main

import (
	"aoc/utils"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("p1", func(t *testing.T) {
		tests := utils.TestSuite[[]string, int]{
			{[]string{
				"5 1 9 5",
				"7 5 3",
				"2 4 6 8",
			}, 18},
		}
		tests.Run(t, P1)
	})
	t.Run("p2", func(t *testing.T) {
		tests := utils.TestSuite[[]string, int]{
			{[]string{
				"5 9 2 8",
				"9 4 7 3",
				"3 8 6 5",
			}, 9},
		}
		tests.Run(t, P2)
	})
}
