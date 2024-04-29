package main

import (
	"aoc/utils"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("p1", func(t *testing.T) {
		utils.TestSuite[string, int]{
			{"ADVENT", 6},
			{"A(1x5)BC", 7},
			{"(3x3)XYZ", 9},
			{"A(2x2)BCD(2x2)EFG", 11},
			{"(6x1)(1x3)A", 6},
			{"X(8x2)(3x3)ABCY", 18},
		}.Run(t, decompressV1)
	})
}
