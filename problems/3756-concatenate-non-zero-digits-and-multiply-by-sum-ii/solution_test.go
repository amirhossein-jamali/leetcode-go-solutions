package concatenatenonzerodigitsandmultiplybysumii

import (
	"math/rand"
	"reflect"
	"testing"
)

func referenceSumAndMultiply(s string, queries [][]int) []int {
	const MOD = 1_000_000_007

	ans := make([]int, len(queries))

	for i, query := range queries {
		sub := s[query[0] : query[1]+1]

		x := 0
		sum := 0

		for j := 0; j < len(sub); j++ {
			d := int(sub[j] - '0')
			if d == 0 {
				continue
			}

			x = (x*10 + d) % MOD
			sum += d
		}

		ans[i] = (x * sum) % MOD
	}

	return ans
}

func TestSumAndMultiply(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		queries [][]int
		want    []int
	}{
		{
			name:    "leetcode example 1",
			s:       "10203004",
			queries: [][]int{{0, 7}, {1, 3}, {4, 6}},
			want:    []int{12340, 4, 9},
		},
		{
			name:    "leetcode example 2",
			s:       "1000",
			queries: [][]int{{0, 3}, {1, 1}},
			want:    []int{1, 0},
		},
		{
			name:    "leetcode example 3",
			s:       "9876543210",
			queries: [][]int{{0, 9}},
			want:    []int{444444137},
		},
		{
			name:    "single zero query",
			s:       "0",
			queries: [][]int{{0, 0}},
			want:    []int{0},
		},
		{
			name:    "single non-zero digit query",
			s:       "5",
			queries: [][]int{{0, 0}},
			want:    []int{25},
		},
		{
			name:    "query with only zeros in range",
			s:       "10001",
			queries: [][]int{{1, 3}},
			want:    []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sumAndMultiply(tt.s, tt.queries)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf(
					"sumAndMultiply(%q, %v) = %v, want %v",
					tt.s,
					tt.queries,
					got,
					tt.want,
				)
			}

			reference := referenceSumAndMultiply(tt.s, tt.queries)
			if !reflect.DeepEqual(got, reference) {
				t.Fatalf(
					"%s: reference = %v, optimized = %v",
					tt.name,
					reference,
					got,
				)
			}
		})
	}
}

func TestSumAndMultiplyRandomized(t *testing.T) {
	rng := rand.New(rand.NewSource(42))

	for trial := 0; trial < 200; trial++ {
		length := 1 + rng.Intn(20)
		digits := make([]byte, length)

		for i := range digits {
			digits[i] = byte('0' + rng.Intn(10))
		}

		s := string(digits)
		queryCount := 1 + rng.Intn(8)
		queries := make([][]int, queryCount)

		for i := range queries {
			left := rng.Intn(length)
			right := left + rng.Intn(length-left)
			queries[i] = []int{left, right}
		}

		got := sumAndMultiply(s, queries)
		want := referenceSumAndMultiply(s, queries)

		if !reflect.DeepEqual(got, want) {
			t.Fatalf(
				"trial %d: sumAndMultiply(%q, %v) = %v, want %v",
				trial,
				s,
				queries,
				got,
				want,
			)
		}
	}
}
