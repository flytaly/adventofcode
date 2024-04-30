package main

import (
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("p1", func(t *testing.T) {
		lines := []string{
			"value 5 goes to bot 2",
			"bot 2 gives low to bot 1 and high to bot 0",
			"value 3 goes to bot 1",
			"bot 1 gives low to output 1 and high to bot 0",
			"bot 0 gives low to output 2 and high to output 0",
			"value 2 goes to bot 2",
		}
		want := 2
		if got := Process(lines, 2, 5, true); !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
