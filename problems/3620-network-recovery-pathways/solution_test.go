package networkrecoverypathways

import (
	"math"
	"math/rand"
	"testing"
)

func bruteForceMaxPathScore(edges [][]int, online []bool, k int64) int {
	n := len(online)
	target := n - 1

	adj := make([][]struct {
		to   int
		cost int
	}, n)

	for _, edge := range edges {
		from, to, cost := edge[0], edge[1], edge[2]
		if !online[from] || !online[to] {
			continue
		}
		adj[from] = append(adj[from], struct {
			to   int
			cost int
		}{to: to, cost: cost})
	}

	best := -1

	var dfs func(node int, total int64, bottleneck int, visited []bool)
	dfs = func(node int, total int64, bottleneck int, visited []bool) {
		if total > k {
			return
		}

		if node == target {
			if bottleneck > best {
				best = bottleneck
			}
			return
		}

		for _, edge := range adj[node] {
			if visited[edge.to] {
				continue
			}

			nextBottleneck := bottleneck
			if edge.cost < nextBottleneck {
				nextBottleneck = edge.cost
			}

			visited[edge.to] = true
			dfs(edge.to, total+int64(edge.cost), nextBottleneck, visited)
			visited[edge.to] = false
		}
	}

	visited := make([]bool, n)
	visited[0] = true
	dfs(0, 0, math.MaxInt, visited)

	return best
}

func TestFindMaxPathScore(t *testing.T) {
	tests := []struct {
		name   string
		edges  [][]int
		online []bool
		k      int64
		want   int
	}{
		{
			name:   "example 1",
			edges:  [][]int{{0, 1, 5}, {1, 3, 10}, {0, 2, 3}, {2, 3, 4}},
			online: []bool{true, true, true, true},
			k:      10,
			want:   3,
		},
		{
			name:   "example 2",
			edges:  [][]int{{0, 1, 7}, {1, 4, 5}, {0, 2, 6}, {2, 3, 6}, {3, 4, 2}, {2, 4, 6}},
			online: []bool{true, true, true, false, true},
			k:      12,
			want:   6,
		},
		{
			name:   "minimum graph with one direct valid edge",
			edges:  [][]int{{0, 1, 4}},
			online: []bool{true, true},
			k:      4,
			want:   4,
		},
		{
			name:   "minimum graph with no edge",
			edges:  [][]int{},
			online: []bool{true, true},
			k:      10,
			want:   -1,
		},
		{
			name:   "direct edge cost exceeds k",
			edges:  [][]int{{0, 1, 11}},
			online: []bool{true, true},
			k:      10,
			want:   -1,
		},
		{
			name:   "intermediate node offline blocks path",
			edges:  [][]int{{0, 1, 3}, {1, 2, 3}},
			online: []bool{true, false, true},
			k:      10,
			want:   -1,
		},
		{
			name:   "zero cost path with k zero",
			edges:  [][]int{{0, 1, 0}, {1, 2, 0}},
			online: []bool{true, true, true},
			k:      0,
			want:   0,
		},
		{
			name: "cheaper partial path has lower bottleneck",
			edges: [][]int{
				{0, 1, 10},
				{0, 2, 1},
				{2, 1, 1},
				{1, 3, 10},
			},
			online: []bool{true, true, true, true},
			k:      100,
			want:   10,
		},
		{
			name: "cheapest prefix alone is insufficient",
			edges: [][]int{
				{0, 1, 10},
				{0, 2, 1},
				{2, 1, 1},
				{1, 3, 10},
			},
			online: []bool{true, true, true, true},
			k:      25,
			want:   10,
		},
		{
			name: "later predecessor enables shortest route to target",
			edges: [][]int{
				{0, 1, 100},
				{0, 2, 1},
				{2, 1, 1},
				{1, 3, 8},
			},
			online: []bool{true, true, true, true},
			k:      10,
			want:   1,
		},
		{
			name: "duplicate edge costs",
			edges: [][]int{
				{0, 1, 5},
				{0, 1, 5},
				{1, 2, 5},
			},
			online: []bool{true, true, true},
			k:      10,
			want:   5,
		},
		{
			name: "multiple valid paths different bottlenecks",
			edges: [][]int{
				{0, 1, 3},
				{0, 2, 7},
				{1, 3, 3},
				{2, 3, 7},
			},
			online: []bool{true, true, true, true},
			k:      14,
			want:   7,
		},
		{
			name:   "accumulated cost exactly k",
			edges:  [][]int{{0, 1, 4}, {1, 2, 6}},
			online: []bool{true, true, true},
			k:      10,
			want:   4,
		},
		{
			name:   "all paths exceed k",
			edges:  [][]int{{0, 1, 6}, {1, 2, 6}},
			online: []bool{true, true, true},
			k:      10,
			want:   -1,
		},
		{
			name: "overflow sensitive path total",
			edges: [][]int{
				{0, 1, 2_000_000_000},
				{1, 2, 2_000_000_000},
				{0, 2, 1},
			},
			online: []bool{true, true, true},
			k:      4_000_000_000,
			want:   2_000_000_000,
		},
		{
			name: "disconnected unrelated online component",
			edges: [][]int{
				{0, 1, 5},
				{2, 3, 9},
			},
			online: []bool{true, true, true, true},
			k:      10,
			want:   -1,
		},
		{
			name: "manual nontrivial dag",
			edges: [][]int{
				{0, 1, 7},
				{0, 2, 100},
				{1, 3, 7},
				{2, 3, 1},
				{1, 4, 100},
				{3, 4, 7},
			},
			online: []bool{true, true, true, true, true},
			k:      21,
			want:   7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMaxPathScore(tt.edges, tt.online, tt.k)
			if got != tt.want {
				t.Fatalf(
					"findMaxPathScore(%v, %v, %d) = %d, want %d",
					tt.edges,
					tt.online,
					tt.k,
					got,
					tt.want,
				)
			}

			if len(tt.edges) <= 8 && len(tt.online) <= 8 {
				brute := bruteForceMaxPathScore(tt.edges, tt.online, tt.k)
				if got != brute {
					t.Fatalf("brute force = %d, optimized = %d", brute, got)
				}
			}
		})
	}
}

func TestFindMaxPathScoreRandomized(t *testing.T) {
	rng := rand.New(rand.NewSource(42))

	for trial := 0; trial < 200; trial++ {
		n := 2 + rng.Intn(6)
		m := rng.Intn(n*(n-1)/2 + 1)

		online := make([]bool, n)
		for i := 0; i < n; i++ {
			online[i] = rng.Intn(2) == 0
		}
		online[0] = true
		online[n-1] = true

		edges := make([][]int, 0, m)
		for u := 0; u < n; u++ {
			for v := u + 1; v < n; v++ {
				if rng.Intn(3) == 0 {
					cost := rng.Intn(21)
					edges = append(edges, []int{u, v, cost})
				}
			}
		}

		k := int64(rng.Intn(60))

		got := findMaxPathScore(edges, online, k)
		want := bruteForceMaxPathScore(edges, online, k)

		if got != want {
			t.Fatalf(
				"trial %d: n=%d edges=%v online=%v k=%d got=%d want=%d",
				trial,
				n,
				edges,
				online,
				k,
				got,
				want,
			)
		}
	}
}
