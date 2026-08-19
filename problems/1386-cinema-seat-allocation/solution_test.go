package cinemaseatallocation

import "testing"

func TestMaxNumberOfFamilies(t *testing.T) {
	tests := []struct {
		name          string
		n             int
		reservedSeats [][]int
		want          int
	}{
		{
			name: "leetcode example 1",
			n:    3,
			reservedSeats: [][]int{
				{1, 2},
				{1, 3},
				{1, 8},
				{2, 6},
				{3, 1},
				{3, 10},
			},
			want: 4,
		},
		{
			name: "leetcode example 2",
			n:    2,
			reservedSeats: [][]int{
				{2, 1},
				{1, 8},
				{2, 6},
			},
			want: 2,
		},
		{
			name: "leetcode example 3",
			n:    4,
			reservedSeats: [][]int{
				{4, 3},
				{1, 4},
				{4, 6},
				{1, 7},
			},
			want: 4,
		},
		{
			name: "aisle seats only",
			n:    1,
			reservedSeats: [][]int{
				{1, 1},
				{1, 10},
			},
			want: 2,
		},
		{
			name: "single middle-left seat",
			n:    1,
			reservedSeats: [][]int{
				{1, 4},
			},
			want: 1,
		},
		{
			name: "left and right blocked",
			n:    1,
			reservedSeats: [][]int{
				{1, 2},
				{1, 8},
			},
			want: 1,
		},
		{
			name: "all three blocks blocked",
			n:    1,
			reservedSeats: [][]int{
				{1, 4},
				{1, 6},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxNumberOfFamilies(tt.n, tt.reservedSeats)
			if got != tt.want {
				t.Fatalf("maxNumberOfFamilies(%d, %v) = %d, want %d", tt.n, tt.reservedSeats, got, tt.want)
			}
		})
	}
}
