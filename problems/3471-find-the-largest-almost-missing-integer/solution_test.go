package findthelargestalmostmissinginteger

import "testing"

func TestLargestInteger(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "leetcode example 1",
			nums: []int{3, 9, 2, 1, 7},
			k:    3,
			want: 7,
		},
		{
			name: "leetcode example 2",
			nums: []int{3, 9, 7, 2, 1, 7},
			k:    4,
			want: 3,
		},
		{
			name: "leetcode example 3",
			nums: []int{0, 0},
			k:    1,
			want: -1,
		},
		{
			name: "single element",
			nums: []int{5},
			k:    1,
			want: 5,
		},
		{
			name: "whole array with duplicates",
			nums: []int{5, 5, 2},
			k:    3,
			want: 5,
		},
		{
			name: "unit windows unique maximum",
			nums: []int{1, 2, 3, 2},
			k:    1,
			want: 3,
		},
		{
			name: "endpoints duplicated",
			nums: []int{5, 1, 2, 5},
			k:    2,
			want: -1,
		},
		{
			name: "unique endpoints last is smaller",
			nums: []int{8, 1, 2, 7},
			k:    2,
			want: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := largestInteger(tt.nums, tt.k)
			if got != tt.want {
				t.Fatalf("largestInteger(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
