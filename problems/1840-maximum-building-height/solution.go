package maximumbuildingheight

import "slices"

type limit struct {
	position int
	height   int
}

func maxBuilding(n int, restrictions [][]int) int {
	limits := make([]limit, 0, len(restrictions)+2)
	limits = append(limits, limit{position: 1, height: 0})

	for _, restriction := range restrictions {
		limits = append(limits, limit{
			position: restriction[0],
			height:   restriction[1],
		})
	}

	limits = append(limits, limit{position: n, height: n - 1})

	slices.SortFunc(limits, func(a, b limit) int {
		return a.position - b.position
	})

	merged := limits[:0]
	for _, current := range limits {
		if len(merged) > 0 && merged[len(merged)-1].position == current.position {
			if current.height < merged[len(merged)-1].height {
				merged[len(merged)-1].height = current.height
			}
			continue
		}
		merged = append(merged, current)
	}
	limits = merged

	for i := 1; i < len(limits); i++ {
		distance := limits[i].position - limits[i-1].position
		maxReachable := limits[i-1].height + distance

		if limits[i].height > maxReachable {
			limits[i].height = maxReachable
		}
	}

	maxHeight := 0

	for i := len(limits) - 2; i >= 0; i-- {
		distance := limits[i+1].position - limits[i].position
		maxReachable := limits[i+1].height + distance

		if limits[i].height > maxReachable {
			limits[i].height = maxReachable
		}

		peak := (limits[i].height + limits[i+1].height + distance) / 2

		if peak > maxHeight {
			maxHeight = peak
		}
	}

	return maxHeight
}

func maxBuildingInPlace(n int, restrictions [][]int) int {
	if len(restrictions) == 0 {
		return n - 1
	}

	slices.SortFunc(restrictions, func(a, b []int) int {
		return a[0] - b[0]
	})

	previousPosition := 1
	previousHeight := 0

	for i := 0; i < len(restrictions); i++ {
		position := restrictions[i][0]
		maxReachable := previousHeight + position - previousPosition

		if restrictions[i][1] > maxReachable {
			restrictions[i][1] = maxReachable
		}

		previousPosition = position
		previousHeight = restrictions[i][1]
	}

	last := restrictions[len(restrictions)-1]
	maxHeight := last[1] + n - last[0]

	for i := len(restrictions) - 2; i >= 0; i-- {
		current := restrictions[i]
		next := restrictions[i+1]

		distance := next[0] - current[0]
		maxReachable := next[1] + distance

		if restrictions[i][1] > maxReachable {
			restrictions[i][1] = maxReachable
		}

		peak := (restrictions[i][1] + next[1] + distance) / 2

		if peak > maxHeight {
			maxHeight = peak
		}
	}

	firstPosition := restrictions[0][0]
	firstHeight := restrictions[0][1]
	firstPeak := (firstHeight + firstPosition - 1) / 2

	if firstPeak > maxHeight {
		maxHeight = firstPeak
	}

	return maxHeight
}
