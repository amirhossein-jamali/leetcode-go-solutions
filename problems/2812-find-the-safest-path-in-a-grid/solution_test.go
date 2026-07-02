package findthesafestpathinagrid

import "testing"

func cloneGrid(grid [][]int) [][]int {
	cloned := make([][]int, len(grid))

	for i, row := range grid {
		cloned[i] = append([]int(nil), row...)
	}

	return cloned
}

func TestMaximumSafenessFactor(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "example 1",
			grid: [][]int{
				{1, 0, 0},
				{0, 0, 0},
				{0, 0, 1},
			},
			want: 0,
		},
		{
			name: "example 2",
			grid: [][]int{
				{0, 0, 1},
				{0, 0, 0},
				{0, 0, 0},
			},
			want: 2,
		},
		{
			name: "example 3",
			grid: [][]int{
				{0, 0, 0, 1},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{1, 0, 0, 0},
			},
			want: 2,
		},
		{
			name: "start or end on thief",
			grid: [][]int{
				{1, 0},
				{0, 0},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid := cloneGrid(tt.grid)
			got := maximumSafenessFactor(grid)
			if got != tt.want {
				t.Fatalf(
					"maximumSafenessFactor(%v) = %d, want %d",
					tt.grid,
					got,
					tt.want,
				)
			}
		})
	}
}
