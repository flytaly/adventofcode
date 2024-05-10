package main

import (
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("p1", func(t *testing.T) {
		lines := []string{
			"The first floor contains a hydrogen-compatible microchip and a lithium-compatible microchip.",
			"The second floor contains a hydrogen generator.",
			"The third floor contains a lithium generator.",
			"The fourth floor contains nothing relevant.",
		}
		want := 11
		if got := Count(lines, false); !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
