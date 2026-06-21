package maximumicecreambars

import "testing"

func TestMaxIceCream(t *testing.T) {
	tests := []struct {
		name   string
		costs  []int
		coins  int
		want   int
	}{
		{
			name:  "example 1",
			costs: []int{1, 3, 2, 4, 1},
			coins: 7,
			want:  4,
		},
		{
			name:  "example 2",
			costs: []int{10, 6, 8, 7, 7, 8},
			coins: 5,
			want:  0,
		},
		{
			name:  "example 3",
			costs: []int{1, 6, 3, 1, 2, 5},
			coins: 20,
			want:  6,
		},
		{
			name:  "single affordable bar",
			costs: []int{3},
			coins: 3,
			want:  1,
		},
		{
			name:  "single unaffordable bar",
			costs: []int{5},
			coins: 4,
			want:  0,
		},
		{
			name:  "all same price partial buy",
			costs: []int{2, 2, 2, 2},
			coins: 5,
			want:  2,
		},
		{
			name:  "exact coins for repeated cheapest",
			costs: []int{1, 1, 1, 5},
			coins: 3,
			want:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxIceCream(tt.costs, tt.coins)
			if got != tt.want {
				t.Fatalf("maxIceCream(%v, %d) = %d, want %d", tt.costs, tt.coins, got, tt.want)
			}
		})
	}
}
