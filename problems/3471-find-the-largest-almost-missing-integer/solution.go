package findthelargestalmostmissinginteger

func largestInteger(nums []int, k int) int {
	n := len(nums)

	if k == n {
		ans := nums[0]
		for i := 1; i < n; i++ {
			if nums[i] > ans {
				ans = nums[i]
			}
		}
		return ans
	}

	var freq [51]int
	for _, num := range nums {
		freq[num]++
	}

	if k == 1 {
		ans := -1
		for num, count := range freq {
			if count == 1 {
				ans = num
			}
		}
		return ans
	}

	ans := -1

	if freq[nums[0]] == 1 {
		ans = nums[0]
	}

	if freq[nums[n-1]] == 1 && nums[n-1] > ans {
		ans = nums[n-1]
	}

	return ans
}
