package smalleststableindexii

func firstStableIndex(nums []int, k int) int {
	ans, mx, pm := 0, nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > mx {
			mx = nums[i]
		}
		if i == ans {
			pm = mx
		}
		if pm-nums[i] > k {
			ans = i + 1
		}
	}

	if ans == len(nums) {
		return -1
	}
	return ans
}
