package countsubarrayswithmajorityelementi

func countMajoritySubarrays(nums []int, target int) int {
	n := len(nums)
	otherCount := 0

	for i := 0; i < n; i++ {
		if nums[i] == target {
			nums[i] = 1
		} else {
			nums[i] = 0
			otherCount++
		}
	}

	base := n + 1
	zeroFrequency := 0

	if otherCount == 0 {
		zeroFrequency = 1
	} else {
		nums[otherCount-1] += base
	}

	prefix := 0
	smaller := 0
	answer := 0

	for i := 0; i < n; i++ {
		isTarget := nums[i]%base == 1

		if isTarget {
			index := prefix + otherCount

			if index == 0 {
				smaller += zeroFrequency
			} else {
				smaller += nums[index-1] / base
			}

			prefix++
		} else {
			prefix--

			index := prefix + otherCount

			if index == 0 {
				smaller -= zeroFrequency
			} else {
				smaller -= nums[index-1] / base
			}
		}

		answer += smaller

		index := prefix + otherCount

		if index == 0 {
			zeroFrequency++
		} else {
			nums[index-1] += base
		}
	}

	return answer
}
