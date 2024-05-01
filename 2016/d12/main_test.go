package main

import (
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("p1", func(t *testing.T) {
		lines := []string{
			"cpy 41 a",
			"inc a",
			"inc a",
			"dec a",
			"jnz a 2",
			"dec a",
		}
		want := 42
		if got := execute(lines, map[string]int{}); !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
