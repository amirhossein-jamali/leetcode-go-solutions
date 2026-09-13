package imageoverlap

import "testing"

func TestLargestOverlap(t *testing.T) {
	tests := []struct {
		name string
		img1 [][]int
		img2 [][]int
		want int
	}{
		{
			name: "leetcode example 1",
			img1: [][]int{
				{1, 1, 0},
				{0, 1, 0},
				{0, 1, 0},
			},
			img2: [][]int{
				{0, 0, 0},
				{0, 1, 1},
				{0, 0, 1},
			},
			want: 3,
		},
		{
			name: "leetcode example 2",
			img1: [][]int{{1}},
			img2: [][]int{{1}},
			want: 1,
		},
		{
			name: "leetcode example 3",
			img1: [][]int{{0}},
			img2: [][]int{{0}},
			want: 0,
		},
		{
			name: "identical 2x2 images",
			img1: [][]int{
				{1, 1},
				{1, 0},
			},
			img2: [][]int{
				{1, 1},
				{1, 0},
			},
			want: 3,
		},
		{
			name: "no ones in first image",
			img1: [][]int{
				{0, 0},
				{0, 0},
			},
			img2: [][]int{
				{1, 1},
				{1, 1},
			},
			want: 0,
		},
		{
			name: "single ones aligned by translation",
			img1: [][]int{
				{1, 0},
				{0, 0},
			},
			img2: [][]int{
				{0, 0},
				{0, 1},
			},
			want: 1,
		},
		{
			name: "disjoint patterns with partial overlap possible",
			img1: [][]int{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			img2: [][]int{
				{0, 0, 1},
				{0, 1, 0},
				{1, 0, 0},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := largestOverlap(tt.img1, tt.img2)
			if got != tt.want {
				t.Fatalf(
					"largestOverlap(%v, %v) = %d, want %d",
					tt.img1,
					tt.img2,
					got,
					tt.want,
				)
			}
		})
	}
}
