package main

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		name         string
		span         Span
		convMap      ConvMap
		wantMapped   []Span
		wantUnMapped []Span
	}{
		{"left_not_intersect", Span{10, 20}, ConvMap{1, 30, 10}, nil, []Span{{10, 20}}},
		{"right_not_intersect", Span{30, 40}, ConvMap{1, 10, 10}, nil, []Span{{30, 40}}},
		{"left_intersect", Span{10, 20}, ConvMap{1, 15, 15}, []Span{{1, 6}}, []Span{{10, 15}}},
		{"left_overlap", Span{10, 40}, ConvMap{1, 15, 15}, []Span{{1, 16}}, []Span{{10, 15}, {30, 40}}},
		{"overlap_exact", Span{15, 30}, ConvMap{1, 15, 15}, []Span{{1, 16}}, nil},
		{"overlap_nested", Span{16, 20}, ConvMap{1, 15, 15}, []Span{{2, 6}}, nil},
		{"right_overlap", Span{15, 40}, ConvMap{1, 15, 15}, []Span{{1, 16}}, []Span{{30, 40}}},
		{"right_intersect", Span{20, 40}, ConvMap{1, 15, 15}, []Span{{6, 16}}, []Span{{30, 40}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := mapSpan(tt.span, tt.convMap)
			if !reflect.DeepEqual(got1, tt.wantMapped) || !reflect.DeepEqual(got2, tt.wantUnMapped) {
				t.Errorf("convertSpan() = %v, %v;  want %v, %v", got1, got2, tt.wantMapped, tt.wantUnMapped)
			}
		})
	}

}
