package minimumscoreofapathbetweentwocities

import (
	"math/rand"
	"testing"
)

func referenceMinScore(n int, roads [][]int) int {
	adjacency := make([][]int, n)

	for _, road := range roads {
		from := road[0] - 1
		to := road[1] - 1

		adjacency[from] = append(adjacency[from], to)
		adjacency[to] = append(adjacency[to], from)
	}

	inComponent := make([]bool, n)
	queue := []int{0}
	inComponent[0] = true

	for head := 0; head < len(queue); head++ {
		city := queue[head]

		for _, neighbor := range adjacency[city] {
			if !inComponent[neighbor] {
				inComponent[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	minimum := 1_000_001

	for _, road := range roads {
		from := road[0] - 1
		to := road[1] - 1
		weight := road[2]

		if inComponent[from] && inComponent[to] && weight < minimum {
			minimum = weight
		}
	}

	return minimum
}

func TestMinScore(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		roads [][]int
		want  int
	}{
		{
			name: "leetcode example 1",
			n:    4,
			roads: [][]int{
				{1, 2, 9},
				{2, 3, 6},
				{2, 4, 5},
				{1, 4, 7},
			},
			want: 5,
		},
		{
			name: "leetcode example 2 revisit allowed",
			n:    4,
			roads: [][]int{
				{1, 2, 2},
				{1, 3, 4},
				{3, 4, 7},
			},
			want: 2,
		},
		{
			name: "direct connection minimum n",
			n:    2,
			roads: [][]int{
				{1, 2, 10000},
			},
			want: 10000,
		},
		{
			name: "irrelevant disconnected component",
			n:    6,
			roads: [][]int{
				{1, 2, 8},
				{2, 6, 7},
				{3, 4, 1},
				{4, 5, 2},
			},
			want: 7,
		},
		{
			name: "branch minimum requires revisiting cities",
			n:    4,
			roads: [][]int{
				{1, 2, 10},
				{2, 4, 8},
				{2, 3, 1},
			},
			want: 1,
		},
		{
			name: "internal cycle edge updates component minimum",
			n:    4,
			roads: [][]int{
				{1, 2, 9},
				{2, 3, 8},
				{1, 3, 2},
				{3, 4, 7},
			},
			want: 2,
		},
		{
			name: "minimum edge before later component merges",
			n:    5,
			roads: [][]int{
				{1, 2, 3},
				{3, 4, 9},
				{4, 5, 8},
				{2, 3, 7},
			},
			want: 3,
		},
		{
			name: "minimum edge from second component during merge",
			n:    5,
			roads: [][]int{
				{1, 2, 9},
				{3, 4, 2},
				{4, 5, 8},
				{2, 3, 7},
			},
			want: 2,
		},
		{
			name: "several edges with maximum allowed weight",
			n:    4,
			roads: [][]int{
				{1, 2, 10000},
				{2, 3, 10000},
				{3, 4, 10000},
			},
			want: 10000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minScore(tt.n, tt.roads)
			if got != tt.want {
				t.Fatalf(
					"minScore(%d, %v) = %d, want %d",
					tt.n,
					tt.roads,
					got,
					tt.want,
				)
			}

			reference := referenceMinScore(tt.n, tt.roads)
			if got != reference {
				t.Fatalf("reference = %d, optimized = %d", reference, got)
			}
		})
	}
}

func TestMinScoreRandomized(t *testing.T) {
	rng := rand.New(rand.NewSource(42))

	for trial := 0; trial < 200; trial++ {
		n := 2 + rng.Intn(8)
		roads := make([][]int, 0, n+rng.Intn(10))

		for city := 1; city < n; city++ {
			weight := 1 + rng.Intn(10000)
			roads = append(roads, []int{city, city + 1, weight})
		}

		extraEdges := rng.Intn(n * 2)
		for edge := 0; edge < extraEdges; edge++ {
			from := 1 + rng.Intn(n-1)
			to := from + 1 + rng.Intn(n-from)
			weight := 1 + rng.Intn(10000)
			roads = append(roads, []int{from, to, weight})
		}

		got := minScore(n, roads)
		want := referenceMinScore(n, roads)

		if got != want {
			t.Fatalf(
				"trial %d: n=%d roads=%v got=%d want=%d",
				trial,
				n,
				roads,
				got,
				want,
			)
		}
	}
}
