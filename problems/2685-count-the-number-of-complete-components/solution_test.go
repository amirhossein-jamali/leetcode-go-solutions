package countthenumberofcompletecomponents

import (
	"math/rand"
	"testing"
)

func referenceCountCompleteComponents(n int, edges [][]int) int {
	adjacency := make([][]int, n)

	for _, edge := range edges {
		from := edge[0]
		to := edge[1]

		adjacency[from] = append(adjacency[from], to)
		adjacency[to] = append(adjacency[to], from)
	}

	visited := make([]bool, n)
	answer := 0

	for start := 0; start < n; start++ {
		if visited[start] {
			continue
		}

		queue := []int{start}
		visited[start] = true
		component := []int{start}

		for head := 0; head < len(queue); head++ {
			node := queue[head]

			for _, neighbor := range adjacency[node] {
				if visited[neighbor] {
					continue
				}

				visited[neighbor] = true
				queue = append(queue, neighbor)
				component = append(component, neighbor)
			}
		}

		size := len(component)
		edgeCount := 0

		for i := 0; i < size; i++ {
			for j := i + 1; j < size; j++ {
				left := component[i]
				right := component[j]

				for _, neighbor := range adjacency[left] {
					if neighbor == right {
						edgeCount++
						break
					}
				}
			}
		}

		if edgeCount == size*(size-1)/2 {
			answer++
		}
	}

	return answer
}

func TestCountCompleteComponents(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		edges [][]int
		want  int
	}{
		{
			name:  "leetcode example 1",
			n:     6,
			edges: [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}},
			want:  3,
		},
		{
			name:  "leetcode example 2",
			n:     6,
			edges: [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}, {3, 5}},
			want:  1,
		},
		{
			name:  "single isolated vertex",
			n:     1,
			edges: nil,
			want:  1,
		},
		{
			name:  "two isolated vertices",
			n:     2,
			edges: nil,
			want:  2,
		},
		{
			name:  "complete graph on three vertices",
			n:     3,
			edges: [][]int{{0, 1}, {0, 2}, {1, 2}},
			want:  1,
		},
		{
			name:  "path of three vertices is not complete",
			n:     3,
			edges: [][]int{{0, 1}, {1, 2}},
			want:  0,
		},
		{
			name:  "complete pair and isolated vertex",
			n:     3,
			edges: [][]int{{0, 1}},
			want:  2,
		},
		{
			name:  "complete graph on four vertices",
			n:     4,
			edges: [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}, {2, 3}},
			want:  1,
		},
		{
			name:  "two complete pairs",
			n:     4,
			edges: [][]int{{0, 1}, {2, 3}},
			want:  2,
		},
		{
			name:  "star graph is not complete",
			n:     4,
			edges: [][]int{{0, 1}, {0, 2}, {0, 3}},
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countCompleteComponents(tt.n, tt.edges)
			if got != tt.want {
				t.Fatalf(
					"countCompleteComponents(%d, %v) = %d, want %d",
					tt.n,
					tt.edges,
					got,
					tt.want,
				)
			}

			reference := referenceCountCompleteComponents(tt.n, tt.edges)
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

func TestCountCompleteComponentsRandomized(t *testing.T) {
	rng := rand.New(rand.NewSource(42))

	for trial := 0; trial < 200; trial++ {
		n := 1 + rng.Intn(12)
		edgeLimit := n * (n - 1) / 2
		edgeCount := rng.Intn(edgeLimit + 1)

		edges := make([][]int, 0, edgeCount)
		seen := make(map[[2]int]struct{}, edgeCount)

		for len(edges) < edgeCount {
			left := rng.Intn(n)
			right := rng.Intn(n)

			if left == right {
				continue
			}

			if left > right {
				left, right = right, left
			}

			key := [2]int{left, right}
			if _, exists := seen[key]; exists {
				continue
			}

			seen[key] = struct{}{}
			edges = append(edges, []int{left, right})
		}

		got := countCompleteComponents(n, edges)
		want := referenceCountCompleteComponents(n, edges)

		if got != want {
			t.Fatalf(
				"trial %d: countCompleteComponents(%d, %v) = %d, want %d",
				trial,
				n,
				edges,
				got,
				want,
			)
		}
	}
}
