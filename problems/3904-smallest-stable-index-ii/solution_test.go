package smalleststableindexii

import "testing"

func referenceFirstStableIndex(nums []int, k int) int {
	n := len(nums)
	prefixMax := make([]int, n)
	suffixMin := make([]int, n)

	prefixMax[0] = nums[0]
	for i := 1; i < n; i++ {
		prefixMax[i] = prefixMax[i-1]
		if nums[i] > prefixMax[i] {
			prefixMax[i] = nums[i]
		}
	}

	suffixMin[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		suffixMin[i] = suffixMin[i+1]
		if nums[i] < suffixMin[i] {
			suffixMin[i] = nums[i]
		}
	}

	for i := 0; i < n; i++ {
		if prefixMax[i]-suffixMin[i] <= k {
			return i
		}
	}

	return -1
}

func TestFirstStableIndex(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "leetcode example 1",
			nums: []int{5, 0, 1, 4},
			k:    3,
			want: 3,
		},
		{
			name: "leetcode example 2",
			nums: []int{3, 2, 1},
			k:    1,
			want: -1,
		},
		{
			name: "leetcode example 3",
			nums: []int{0},
			k:    0,
			want: 0,
		},
		{
			name: "first index is stable",
			nums: []int{1, 2, 3},
			k:    10,
			want: 0,
		},
		{
			name: "stable at last index only",
			nums: []int{10, 1, 9},
			k:    1,
			want: 2,
		},
		{
			name: "all equal values",
			nums: []int{7, 7, 7, 7},
			k:    0,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := firstStableIndex(tt.nums, tt.k)
			if got != tt.want {
				t.Fatalf(
					"firstStableIndex(%v, %d) = %d, want %d",
					tt.nums,
					tt.k,
					got,
					tt.want,
				)
			}

			reference := referenceFirstStableIndex(tt.nums, tt.k)
			if got != reference {
				t.Fatalf(
					"%s: reference = %d, optimized = %d",
					tt.name,
					reference,
					got,
				)
			}
		})
	}
}
