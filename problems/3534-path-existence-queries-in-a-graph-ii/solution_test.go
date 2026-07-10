package pathexistencequeriesinagraphii

import (
	"math/rand"
	"reflect"
	"testing"
)

func referencePathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []int {
	adjacency := make([][]int, n)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			diff := nums[i] - nums[j]
			if diff < 0 {
				diff = -diff
			}

			if diff <= maxDiff {
				adjacency[i] = append(adjacency[i], j)
				adjacency[j] = append(adjacency[j], i)
			}
		}
	}

	shortestDistance := func(start, end int) int {
		if start == end {
			return 0
		}

		distances := make([]int, n)
		for i := range distances {
			distances[i] = -1
		}

		queue := []int{start}
		distances[start] = 0

		for head := 0; head < len(queue); head++ {
			node := queue[head]

			for _, neighbor := range adjacency[node] {
				if distances[neighbor] != -1 {
					continue
				}

				distances[neighbor] = distances[node] + 1
				if neighbor == end {
					return distances[neighbor]
				}

				queue = append(queue, neighbor)
			}
		}

		return -1
	}

	answer := make([]int, len(queries))

	for i, query := range queries {
		answer[i] = shortestDistance(query[0], query[1])
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
		want    []int
	}{
		{
			name:    "leetcode example 1",
			n:       5,
			nums:    []int{1, 8, 3, 4, 2},
			maxDiff: 3,
			queries: [][]int{{0, 3}, {2, 4}},
			want:    []int{1, 1},
		},
		{
			name:    "leetcode example 2",
			n:       5,
			nums:    []int{5, 3, 1, 9, 10},
			maxDiff: 2,
			queries: [][]int{{0, 1}, {0, 2}, {2, 3}, {4, 3}},
			want:    []int{1, 2, -1, 1},
		},
		{
			name:    "leetcode example 3",
			n:       3,
			nums:    []int{3, 6, 1},
			maxDiff: 1,
			queries: [][]int{{0, 0}, {0, 1}, {1, 2}},
			want:    []int{0, -1, -1},
		},
		{
			name:    "fully connected value range",
			n:       4,
			nums:    []int{1, 2, 3, 4},
			maxDiff: 3,
			queries: [][]int{{0, 3}, {1, 2}},
			want:    []int{1, 1},
		},
		{
			name:    "max diff zero with equal values",
			n:       3,
			nums:    []int{5, 5, 7},
			maxDiff: 0,
			queries: [][]int{{0, 1}, {0, 2}},
			want:    []int{1, -1},
		},
		{
			name:    "single node",
			n:       1,
			nums:    []int{42},
			maxDiff: 0,
			queries: [][]int{{0, 0}},
			want:    []int{0},
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

	for trial := 0; trial < 100; trial++ {
		n := 1 + rng.Intn(10)
		nums := make([]int, n)

		for i := range nums {
			nums[i] = rng.Intn(20)
		}

		maxDiff := rng.Intn(8)
		queryCount := 1 + rng.Intn(8)
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
