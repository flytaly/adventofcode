package main

import (
	"aoc/utils"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("p1", func(t *testing.T) {
		tests := utils.TestSuite[[]string, int]{
			{[]string{"2x3x4"}, 58},
			{[]string{"1x1x10"}, 43}}
		tests.Run(t, PartOne)
	})

	t.Run("p2", func(t *testing.T) {
		tests := utils.TestSuite[[]string, int]{
			{[]string{"2x4x3"}, 34},
			{[]string{"10x1x1"}, 14}}
		tests.Run(t, PartTwo)
	})
}
