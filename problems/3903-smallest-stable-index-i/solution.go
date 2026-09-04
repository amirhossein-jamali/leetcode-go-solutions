package smalleststableindexi

func firstStableIndex(nums []int, k int) int {
	n := len(nums)
	candidate := 0
	prefixMax := nums[0]
	candidateMax := prefixMax

	for i := 1; i < n; i++ {
		x := nums[i]

		if x > prefixMax {
			prefixMax = x
		}

		if i == candidate {
			candidateMax = prefixMax
		}

		if x+k < candidateMax {
			candidate = i + 1
		}
	}

	if candidate == n {
		return -1
	}
	return candidate
}
