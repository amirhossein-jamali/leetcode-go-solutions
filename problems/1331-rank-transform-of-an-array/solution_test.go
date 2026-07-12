package ranktransformofanarray

import (
	"math/rand"
	"sort"
	"testing"
)

func referenceArrayRankTransform(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return []int{}
	}

	sorted := make([]int, n)
	copy(sorted, arr)
	sort.Ints(sorted)

	valueToRank := make(map[int]int)
	rank := 1

	for _, value := range sorted {
		if _, exists := valueToRank[value]; exists {
			continue
		}

		valueToRank[value] = rank
		rank++
	}

	answer := make([]int, n)

	for i, value := range arr {
		answer[i] = valueToRank[value]
	}

	return answer
}

func cloneInts(values []int) []int {
	if len(values) == 0 {
		return nil
	}

	cloned := make([]int, len(values))
	copy(cloned, values)
	return cloned
}

func TestArrayRankTransform(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "leetcode example 1",
			arr:  []int{40, 10, 20, 30},
			want: []int{4, 1, 2, 3},
		},
		{
			name: "leetcode example 2",
			arr:  []int{100, 100, 100},
			want: []int{1, 1, 1},
		},
		{
			name: "leetcode example 3",
			arr:  []int{37, 12, 28, 9, 100, 56, 80, 5, 12},
			want: []int{5, 3, 4, 2, 8, 6, 7, 1, 3},
		},
		{
			name: "empty array",
			arr:  []int{},
			want: []int{},
		},
		{
			name: "single element",
			arr:  []int{42},
			want: []int{1},
		},
		{
			name: "already sorted distinct values",
			arr:  []int{1, 2, 3, 4},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "reverse sorted distinct values",
			arr:  []int{4, 3, 2, 1},
			want: []int{4, 3, 2, 1},
		},
		{
			name: "pairs of equal values",
			arr:  []int{5, 5, 1, 1},
			want: []int{2, 2, 1, 1},
		},
		{
			name: "negative values",
			arr:  []int{-5, -1, -5, 0},
			want: []int{1, 2, 1, 3},
		},
		{
			name: "all distinct negative and positive",
			arr:  []int{-100, 0, 100},
			want: []int{1, 2, 3},
		},
		{
			name: "constraint minimum value",
			arr:  []int{-1_000_000_000, 0, 1_000_000_000},
			want: []int{1, 2, 3},
		},
		{
			name: "duplicate values at constraint extremes",
			arr:  []int{-1_000_000_000, -1_000_000_000, 1_000_000_000},
			want: []int{1, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := cloneInts(tt.arr)
			got := arrayRankTransform(input)

			if len(got) != len(tt.want) {
				t.Fatalf(
					"arrayRankTransform(%v) length = %d, want %d",
					tt.arr,
					len(got),
					len(tt.want),
				)
			}

			for i := range got {
				if got[i] != tt.want[i] {
					t.Fatalf(
						"arrayRankTransform(%v) = %v, want %v",
						tt.arr,
						got,
						tt.want,
					)
				}
			}

			reference := referenceArrayRankTransform(tt.arr)
			if len(got) != len(reference) {
				t.Fatalf(
					"%s: reference length = %d, optimized length = %d",
					tt.name,
					len(reference),
					len(got),
				)
			}

			for i := range reference {
				if got[i] != reference[i] {
					t.Fatalf(
						"%s: reference = %v, optimized = %v",
						tt.name,
						reference,
						got,
					)
				}
			}
		})
	}
}

func TestArrayRankTransformRandomized(t *testing.T) {
	rng := rand.New(rand.NewSource(42))

	for trial := 0; trial < 200; trial++ {
		n := rng.Intn(20)
		arr := make([]int, n)

		for i := range arr {
			arr[i] = rng.Intn(2_000_001) - 1_000_000
		}

		input := cloneInts(arr)
		got := arrayRankTransform(input)
		want := referenceArrayRankTransform(arr)

		if len(got) != len(want) {
			t.Fatalf(
				"trial %d: length got = %d, want = %d",
				trial,
				len(got),
				len(want),
			)
		}

		for i := range want {
			if got[i] != want[i] {
				t.Fatalf(
					"trial %d: arrayRankTransform(%v) = %v, want %v",
					trial,
					arr,
					got,
					want,
				)
			}
		}
	}
}
