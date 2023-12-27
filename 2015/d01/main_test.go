package main

import (
	"fmt"
	"testing"
)

func TestD1(t *testing.T) {
	t.Run("p1", func(t *testing.T) {
		tests := []struct {
			input string
			want  int
		}{{"(())", 0}, {"))(((((", 3}, {"())", -1}, {")())())", -3}}

		for i, v := range tests {
			t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
				if got := PartOne(v.input); got != v.want {
					t.Errorf("PartOne() = %v, want %v", got, v.want)
				}
			})
		}
	})

	t.Run("p2", func(t *testing.T) {
		tests := []struct {
			input string
			want  int
		}{{")", 1}, {"()())", 5}}

		for i, v := range tests {
			t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
				if got := PartTwo(v.input); got != v.want {
					t.Errorf("PartTwo() = %v, want %v", got, v.want)
				}
			})
		}
	})
}
