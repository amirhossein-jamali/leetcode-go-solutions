package networkrecoverypathways

import "testing"

func buildBenchmarkInput() ([][]int, []bool, int64) {
	n := 5000
	edges := make([][]int, 0, n-1)

	for node := 0; node < n-1; node++ {
		edges = append(edges, []int{node, node + 1, node%50 + 1})
	}

	online := make([]bool, n)
	for i := range online {
		online[i] = true
	}

	return edges, online, 250_000
}

func BenchmarkFindMaxPathScore(b *testing.B) {
	edges, online, k := buildBenchmarkInput()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		findMaxPathScore(edges, online, k)
	}
}
