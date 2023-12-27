package main

import (
	"aoc/utils"
	"testing"
)

func TestD1(t *testing.T) {
	t.Run("p1", func(t *testing.T) {
		tests := utils.TestSuite[string, int]{{"(())", 0}, {"))(((((", 3}, {"())", -1}, {")())())", -3}}
		tests.Run(t, PartOne)
	})

	t.Run("p2", func(t *testing.T) {
		tests := utils.TestSuite[string, int]{{")", 1}, {"()())", 5}}
		tests.Run(t, PartTwo)
	})
}
