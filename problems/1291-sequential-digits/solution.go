package sequentialdigits

func sequentialDigits(low int, high int) []int {
	result := make([]int, 0, 36)

	first := 12
	step := 11

	for digits := 2; digits <= 9 && first <= high; digits++ {
		current := first
		count := 10 - digits

		for i := 0; i < count; i++ {
			if current > high {
				return result
			}

			if current >= low {
				result = append(result, current)
			}

			current += step
		}

		first = first*10 + digits + 1
		step = step*10 + 1
	}

	return result
}
