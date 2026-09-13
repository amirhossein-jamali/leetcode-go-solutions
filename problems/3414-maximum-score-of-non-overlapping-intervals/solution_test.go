package maximumscoreofnonoverlappingintervals

import (
	"reflect"
	"testing"
)

func TestMaximumWeight(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      []int
	}{
		{
			name: "leetcode example 1",
			intervals: [][]int{
				{1, 3, 2},
				{4, 5, 2},
				{1, 5, 5},
				{6, 9, 3},
				{6, 7, 1},
				{8, 9, 1},
			},
			want: []int{2, 3},
		},
		{
			name: "leetcode example 2",
			intervals: [][]int{
				{5, 8, 1},
				{6, 7, 7},
				{4, 7, 3},
				{9, 10, 6},
				{7, 8, 2},
				{11, 14, 3},
				{3, 5, 5},
			},
			want: []int{1, 3, 5, 6},
		},
		{
			name: "single interval",
			intervals: [][]int{
				{1, 2, 9},
			},
			want: []int{0},
		},
		{
			name: "choose heavier non-overlapping interval",
			intervals: [][]int{
				{1, 4, 3},
				{2, 5, 7},
			},
			want: []int{1},
		},
		{
			name: "touching boundaries overlap",
			intervals: [][]int{
				{1, 2, 4},
				{2, 3, 5},
			},
			want: []int{1},
		},
		{
			name: "three disjoint intervals all chosen",
			intervals: [][]int{
				{1, 2, 5},
				{3, 4, 5},
				{5, 6, 5},
			},
			want: []int{0, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumWeight(tt.intervals)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf(
					"maximumWeight(%v) = %v, want %v",
					tt.intervals,
					got,
					tt.want,
				)
			}
		})
	}
}
