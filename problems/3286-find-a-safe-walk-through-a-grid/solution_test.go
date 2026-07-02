package findasafewalkthroughagrid

import "testing"

func cloneGrid(grid [][]int) [][]int {
	cloned := make([][]int, len(grid))

	for i, row := range grid {
		cloned[i] = append([]int(nil), row...)
	}

	return cloned
}

func TestFindSafeWalk(t *testing.T) {
	tests := []struct {
		name   string
		grid   [][]int
		health int
		want   bool
	}{
		{
			name: "example 1",
			grid: [][]int{
				{0, 1, 0, 0, 0},
				{0, 1, 0, 1, 0},
				{0, 0, 0, 1, 0},
			},
			health: 1,
			want:   true,
		},
		{
			name: "example 2",
			grid: [][]int{
				{0, 1, 1, 0, 0, 0},
				{1, 0, 1, 0, 0, 0},
				{0, 1, 1, 1, 0, 1},
				{0, 0, 1, 0, 1, 0},
			},
			health: 3,
			want:   false,
		},
		{
			name: "example 3",
			grid: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			health: 5,
			want:   true,
		},
		{
			name: "cannot leave start",
			grid: [][]int{
				{1, 0},
				{0, 0},
			},
			health: 1,
			want:   false,
		},
		{
			name: "all safe cells",
			grid: [][]int{
				{0, 0},
				{0, 0},
			},
			health: 1,
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid := cloneGrid(tt.grid)
			got := findSafeWalk(grid, tt.health)
			if got != tt.want {
				t.Fatalf(
					"findSafeWalk(%v, %d) = %v, want %v",
					tt.grid,
					tt.health,
					got,
					tt.want,
				)
			}
		})
	}
}
