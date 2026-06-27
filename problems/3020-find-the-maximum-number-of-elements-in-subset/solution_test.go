package findthemaximumnumberofelementsinsubset

import "testing"

func TestMaximumLength(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{5, 4, 1, 2, 2},
			want: 3,
		},
		{
			name: "example 2",
			nums: []int{1, 3, 2, 4},
			want: 1,
		},
		{
			name: "three ones",
			nums: []int{1, 1, 1},
			want: 3,
		},
		{
			name: "power chain 2 4",
			nums: []int{2, 2, 4},
			want: 3,
		},
		{
			name: "two equal non-one elements",
			nums: []int{3, 3, 7},
			want: 1,
		},
		{
			name: "longer palindrome chain",
			nums: []int{2, 2, 4, 4, 16, 16},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumLength(tt.nums)
			if got != tt.want {
				t.Fatalf("maximumLength(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
