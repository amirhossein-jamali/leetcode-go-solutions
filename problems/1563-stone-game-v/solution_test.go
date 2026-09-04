package stonegamev

import (
	"math/rand"
	"testing"
)

func referenceStoneGameV(stoneValue []int) int {
	n := len(stoneValue)
	if n < 2 {
		return 0
	}

	prefix := make([]int, n+1)
	for i, value := range stoneValue {
		prefix[i+1] = prefix[i] + value
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			best := 0

			for k := i; k < j; k++ {
				left := prefix[k+1] - prefix[i]
				right := prefix[j+1] - prefix[k+1]

				score := 0
				switch {
				case left < right:
					score = left + dp[i][k]
				case left > right:
					score = right + dp[k+1][j]
				default:
					leftScore := left + dp[i][k]
					rightScore := right + dp[k+1][j]
					if leftScore > rightScore {
						score = leftScore
					} else {
						score = rightScore
					}
				}

				if score > best {
					best = score
				}
			}

			dp[i][j] = best
		}
	}

	return dp[0][n-1]
}

func TestStoneGameV(t *testing.T) {
	tests := []struct {
		name       string
		stoneValue []int
		want       int
	}{
		{
			name:       "leetcode example 1",
			stoneValue: []int{6, 2, 3, 4, 5, 5},
			want:       18,
		},
		{
			name:       "leetcode example 2",
			stoneValue: []int{7, 7, 7, 7, 7, 7, 7},
			want:       28,
		},
		{
			name:       "leetcode example 3",
			stoneValue: []int{4},
			want:       0,
		},
		{
			name:       "two equal stones",
			stoneValue: []int{1, 1},
			want:       1,
		},
		{
			name:       "two stones left is smaller",
			stoneValue: []int{1, 2},
			want:       1,
		},
		{
			name:       "two stones right is smaller",
			stoneValue: []int{2, 1},
			want:       1,
		},
		{
			name:       "three stones",
			stoneValue: []int{1, 2, 3},
			want:       4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stoneGameV(tt.stoneValue)
			if got != tt.want {
				t.Fatalf("stoneGameV(%v) = %d, want %d", tt.stoneValue, got, tt.want)
			}

			reference := referenceStoneGameV(tt.stoneValue)
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

func TestStoneGameVRandomized(t *testing.T) {
	rng := rand.New(rand.NewSource(42))

	for trial := 0; trial < 200; trial++ {
		n := 1 + rng.Intn(8)
		stoneValue := make([]int, n)

		for i := range stoneValue {
			stoneValue[i] = 1 + rng.Intn(20)
		}

		got := stoneGameV(stoneValue)
		want := referenceStoneGameV(stoneValue)

		if got != want {
			t.Fatalf(
				"trial %d: stoneGameV(%v) = %d, want %d",
				trial,
				stoneValue,
				got,
				want,
			)
		}
	}
}
