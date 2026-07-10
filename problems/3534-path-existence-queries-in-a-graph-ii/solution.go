package pathexistencequeriesinagraphii

func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []int {
	answer := make([]int, len(queries))

	minValue := nums[0]
	maxValue := nums[0]

	for i := 1; i < n; i++ {
		if nums[i] < minValue {
			minValue = nums[i]
		}
		if nums[i] > maxValue {
			maxValue = nums[i]
		}
	}

	if maxValue-minValue <= maxDiff {
		for i, query := range queries {
			if query[0] == query[1] {
				answer[i] = 0
			} else {
				answer[i] = 1
			}
		}

		return answer
	}

	if maxDiff == 0 {
		for i, query := range queries {
			u := query[0]
			v := query[1]

			if u == v {
				answer[i] = 0
			} else if nums[u] == nums[v] {
				answer[i] = 1
			} else {
				answer[i] = -1
			}
		}

		return answer
	}

	valueRange := maxValue - minValue + 1
	rankByValue := make([]int32, valueRange)

	for i := range rankByValue {
		rankByValue[i] = -1
	}

	for _, value := range nums {
		rankByValue[value-minValue] = -2
	}

	values := make([]int32, 0, n)

	for offset := 0; offset < valueRange; offset++ {
		if rankByValue[offset] == -2 {
			rankByValue[offset] = int32(len(values))
			values = append(values, int32(offset+minValue))
		}
	}

	distinctCount := len(values)

	levels := 0
	for 1<<levels < distinctCount {
		levels++
	}

	jump := make([]uint32, levels*distinctCount)

	right := 0
	limit := int32(maxDiff)

	for left := 0; left < distinctCount; left++ {
		if right < left {
			right = left
		}

		for right+1 < distinctCount &&
			values[right+1]-values[left] <= limit {
			right++
		}

		jump[left] = uint32(right)
	}

	for level := 1; level < levels; level++ {
		previousBase := (level - 1) * distinctCount
		currentBase := level * distinctCount

		for i := 0; i < distinctCount; i++ {
			middle := int(jump[previousBase+i])
			jump[currentBase+i] = jump[previousBase+middle]
		}
	}

	for queryIndex, query := range queries {
		u := query[0]
		v := query[1]

		if u == v {
			answer[queryIndex] = 0
			continue
		}

		leftValue := nums[u]
		rightValue := nums[v]

		if leftValue == rightValue {
			answer[queryIndex] = 1
			continue
		}

		if leftValue > rightValue {
			leftValue, rightValue = rightValue, leftValue
		}

		if rightValue-leftValue <= maxDiff {
			answer[queryIndex] = 1
			continue
		}

		current := int(rankByValue[leftValue-minValue])
		target := int(rankByValue[rightValue-minValue])
		distance := 0

		for level := levels - 1; level >= 0; level-- {
			next := int(jump[level*distinctCount+current])

			if next > current && next < target {
				current = next
				distance += 1 << level
			}
		}

		if int(jump[current]) >= target {
			answer[queryIndex] = distance + 1
		} else {
			answer[queryIndex] = -1
		}
	}

	return answer
}
