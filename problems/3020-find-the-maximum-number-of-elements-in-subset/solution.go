package findthemaximumnumberofelementsinsubset

func maximumLength(nums []int) int {
	freq := make(map[int64]int, len(nums))

	for _, num := range nums {
		freq[int64(num)]++
	}

	answer := 1

	if ones := freq[1]; ones > 0 {
		answer = ones - 1 + ones%2
	}

	for start, count := range freq {
		if start == 1 || count < 2 {
			continue
		}

		length := 1
		current := start

		for count >= 2 {
			next := current * current
			nextCount := freq[next]

			if nextCount == 0 {
				break
			}

			length += 2
			current = next
			count = nextCount
		}

		if length > answer {
			answer = length
		}
	}

	return answer
}
