package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("p1", func(t *testing.T) {
		input := []string{
			"Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.",
			"Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.",
		}
		want := 1120
		if got := PartOne(input, 1000); got != want {
			t.Errorf("PartOne() = %v, want %v", got, want)
		}
	})

	t.Run("p2", func(t *testing.T) {
		input := []string{
			"Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.",
			"Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.",
		}
		want := 689
		if got := PartTwo(input, 1000); got != want {
			t.Errorf("PartTwo() = %v, want %v", got, want)
		}
	})
}
