package maximumelementafterdecreasingandrearranging

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	n := len(arr)

	for i := range arr {
		if arr[i] > n {
			arr[i] = n
		}
	}

	for i := 0; i < n; i++ {
		index := (arr[i] - 1) % n
		arr[index] += n
	}

	maxValue := 0

	for value := 1; value <= n; value++ {
		frequency := (arr[value-1] - 1) / n

		maxValue += frequency
		if maxValue > value {
			maxValue = value
		}
	}

	return maxValue
}
