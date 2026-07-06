package removecoveredintervals

import "testing"

func cloneIntervals(intervals [][]int) [][]int {
	cloned := make([][]int, len(intervals))

	for i, interval := range intervals {
		cloned[i] = append([]int(nil), interval...)
	}

	return cloned
}

func TestRemoveCoveredIntervals(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      int
	}{
		{
			name:      "leetcode example 1",
			intervals: [][]int{{1, 4}, {3, 6}, {2, 8}},
			want:      2,
		},
		{
			name:      "leetcode example 2",
			intervals: [][]int{{1, 4}, {2, 3}},
			want:      1,
		},
		{
			name:      "equal start case",
			intervals: [][]int{{1, 4}, {1, 5}},
			want:      1,
		},
		{
			name:      "no covered intervals",
			intervals: [][]int{{1, 2}, {2, 3}, {3, 4}},
			want:      3,
		},
		{
			name:      "multiple covered intervals",
			intervals: [][]int{{1, 10}, {2, 3}, {4, 8}, {10, 11}},
			want:      2,
		},
		{
			name:      "same end different starts",
			intervals: [][]int{{1, 4}, {2, 4}, {3, 4}},
			want:      1,
		},
		{
			name:      "single interval",
			intervals: [][]int{{0, 100000}},
			want:      1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			intervals := cloneIntervals(tt.intervals)
			got := removeCoveredIntervals(intervals)
			if got != tt.want {
				t.Fatalf(
					"removeCoveredIntervals(%v) = %d, want %d",
					tt.intervals,
					got,
					tt.want,
				)
			}
		})
	}
}
