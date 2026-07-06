package removecoveredintervals

import "sort"

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	remaining := 0
	maxEnd := -1

	for _, interval := range intervals {
		if interval[1] > maxEnd {
			remaining++
			maxEnd = interval[1]
		}
	}

	return remaining
}
