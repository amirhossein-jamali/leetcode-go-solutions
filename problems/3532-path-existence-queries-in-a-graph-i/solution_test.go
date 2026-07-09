package pathexistencequeriesinagraphi

import (
	"math/rand"
	"reflect"
	"testing"
)

func referencePathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {
	parent := make([]int, n)

	for i := range parent {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}

		return parent[x]
	}

	union := func(a, b int) {
		rootA := find(a)
		rootB := find(b)

		if rootA != rootB {
			parent[rootB] = rootA
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[j]-nums[i] <= maxDiff {
				union(i, j)
			}
		}
	}

	answer := make([]bool, len(queries))

	for i, query := range queries {
		answer[i] = find(query[0]) == find(query[1])
	}

	return answer
}

func TestPathExistenceQueries(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		nums    []int
		maxDiff int
		queries [][]int
		want    []bool
	}{
		{
			name:    "leetcode example 1",
			n:       2,
			nums:    []int{1, 3},
			maxDiff: 1,
			queries: [][]int{{0, 0}, {0, 1}},
			want:    []bool{true, false},
		},
		{
			name:    "leetcode example 2",
			n:       4,
			nums:    []int{2, 5, 6, 8},
			maxDiff: 2,
			queries: [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}},
			want:    []bool{false, false, true, true},
		},
		{
			name:    "single node",
			n:       1,
			nums:    []int{7},
			maxDiff: 0,
			queries: [][]int{{0, 0}},
			want:    []bool{true},
		},
		{
			name:    "fully connected chain",
			n:       4,
			nums:    []int{1, 2, 3, 4},
			maxDiff: 1,
			queries: [][]int{{0, 3}, {1, 2}},
			want:    []bool{true, true},
		},
		{
			name:    "disconnected components",
			n:       5,
			nums:    []int{1, 2, 10, 11, 20},
			maxDiff: 1,
			queries: [][]int{{0, 1}, {2, 3}, {1, 2}, {0, 4}},
			want:    []bool{true, true, false, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pathExistenceQueries(tt.n, tt.nums, tt.maxDiff, tt.queries)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf(
					"pathExistenceQueries(%d, %v, %d, %v) = %v, want %v",
					tt.n,
					tt.nums,
					tt.maxDiff,
					tt.queries,
					got,
					tt.want,
				)
			}

			reference := referencePathExistenceQueries(tt.n, tt.nums, tt.maxDiff, tt.queries)
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

func TestPathExistenceQueriesRandomized(t *testing.T) {
	rng := rand.New(rand.NewSource(42))

	for trial := 0; trial < 200; trial++ {
		n := 1 + rng.Intn(12)
		nums := make([]int, n)
		nums[0] = rng.Intn(20)

		for i := 1; i < n; i++ {
			nums[i] = nums[i-1] + rng.Intn(6)
		}

		maxDiff := rng.Intn(10)
		queryCount := 1 + rng.Intn(10)
		queries := make([][]int, queryCount)

		for i := range queries {
			left := rng.Intn(n)
			right := rng.Intn(n)
			queries[i] = []int{left, right}
		}

		got := pathExistenceQueries(n, nums, maxDiff, queries)
		want := referencePathExistenceQueries(n, nums, maxDiff, queries)

		if !reflect.DeepEqual(got, want) {
			t.Fatalf(
				"trial %d: pathExistenceQueries(%d, %v, %d, %v) = %v, want %v",
				trial,
				n,
				nums,
				maxDiff,
				queries,
				got,
				want,
			)
		}
	}
}
