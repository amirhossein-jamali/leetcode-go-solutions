package stonegamev

func stoneGameV(stoneValue []int) int {
	n := len(stoneValue)
	if n < 2 {
		return 0
	}

	prefix := make([]int, n+1)
	for i, value := range stoneValue {
		prefix[i+1] = prefix[i] + value
	}

	rightBest := make([]int, n*n)
	leftBest := make([]int, n)

	answer := 0

	for i := n - 1; i >= 0; i-- {
		leftBest[i] = stoneValue[i]
		rightBest[i*n+i] = stoneValue[i]

		mid := i

		for j := i + 1; j < n; j++ {
			total := prefix[j+1] - prefix[i]

			for mid < j &&
				2*(prefix[mid+1]-prefix[i]) < total {
				mid++
			}

			score := 0

			if mid > i {
				score = leftBest[mid-1]
			}

			if mid < j {
				rightScore := rightBest[(mid+1)*n+j]
				if rightScore > score {
					score = rightScore
				}

				leftSum := prefix[mid+1] - prefix[i]
				if 2*leftSum == total && leftBest[mid] > score {
					score = leftBest[mid]
				}
			}

			value := total + score

			if leftBest[j-1] > value {
				leftBest[j] = leftBest[j-1]
			} else {
				leftBest[j] = value
			}

			previousRight := rightBest[(i+1)*n+j]
			if previousRight > value {
				rightBest[i*n+j] = previousRight
			} else {
				rightBest[i*n+j] = value
			}

			if i == 0 && j == n-1 {
				answer = score
			}
		}
	}

	return answer
}
