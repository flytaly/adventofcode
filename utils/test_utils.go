package utils

import (
	"fmt"
	"reflect"
	"testing"
)

type TestSuite[I any, W any] []struct {
	Input I
	Want  W
}

func (tests TestSuite[I, W]) Run(t *testing.T, fn func(I) W) {
	t.Helper()
	for i, v := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if got := fn(v.Input); !reflect.DeepEqual(got, v.Want) {
				t.Errorf("fn() = %v, want %v", got, v.Want)
			}
		})
	}
}
