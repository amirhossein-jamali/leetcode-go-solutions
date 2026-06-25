package countsubarrayswithmajorityelementi

import "testing"

func TestCountMajoritySubarrays(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "example 1",
			nums:   []int{1, 2, 2, 3},
			target: 2,
			want:   5,
		},
		{
			name:   "example 2",
			nums:   []int{1, 1, 1, 1},
			target: 1,
			want:   10,
		},
		{
			name:   "example 3 target absent",
			nums:   []int{1, 2, 3},
			target: 4,
			want:   0,
		},
		{
			name:   "single matching element",
			nums:   []int{7},
			target: 7,
			want:   1,
		},
		{
			name:   "single non-matching element",
			nums:   []int{7},
			target: 8,
			want:   0,
		},
		{
			name:   "no majority in length two",
			nums:   []int{1, 2},
			target: 1,
			want:   1,
		},
		{
			name:   "all others no target",
			nums:   []int{3, 3, 3},
			target: 1,
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := append([]int(nil), tt.nums...)
			got := countMajoritySubarrays(nums, tt.target)
			if got != tt.want {
				t.Fatalf(
					"countMajoritySubarrays(%v, %d) = %d, want %d",
					tt.nums,
					tt.target,
					got,
					tt.want,
				)
			}
		})
	}
}
