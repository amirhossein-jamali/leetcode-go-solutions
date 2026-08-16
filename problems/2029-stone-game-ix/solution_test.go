package stonegameix

import "testing"

func TestStoneGameIX(t *testing.T) {
	tests := []struct {
		name   string
		stones []int
		want   bool
	}{
		{
			name:   "leetcode example 1",
			stones: []int{2, 1},
			want:   true,
		},
		{
			name:   "leetcode example 2",
			stones: []int{2},
			want:   false,
		},
		{
			name:   "leetcode example 3",
			stones: []int{5, 1, 2, 4, 3},
			want:   false,
		},
		{
			name:   "single multiple of three",
			stones: []int{3},
			want:   false,
		},
		{
			name:   "only remainder one stones",
			stones: []int{1, 4},
			want:   false,
		},
		{
			name:   "even zeros with both remainders",
			stones: []int{3, 6, 1, 2},
			want:   true,
		},
		{
			name:   "odd zeros with remainder difference two",
			stones: []int{1, 1, 3},
			want:   false,
		},
		{
			name:   "odd zeros with remainder difference three",
			stones: []int{1, 1, 1, 3},
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stoneGameIX(tt.stones)
			if got != tt.want {
				t.Fatalf("stoneGameIX(%v) = %v, want %v", tt.stones, got, tt.want)
			}
		})
	}
}
