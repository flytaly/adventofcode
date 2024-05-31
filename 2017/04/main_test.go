package main

import (
	"aoc/utils"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("p1", func(t *testing.T) {
		tests := utils.TestSuite[string, bool]{
			{"aa bb cc dd ee", true},
			{"aa bb cc dd aa", false},
			{"aa bb cc dd aaa", true},
		}
		tests.Run(t, isValid1)
	})
}
