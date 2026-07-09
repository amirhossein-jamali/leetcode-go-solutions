package pathexistencequeriesinagraphi

func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {
	component := make([]uint32, n)

	var current uint32

	for i := 1; i < n; i++ {
		if nums[i]-nums[i-1] > maxDiff {
			current++
		}

		component[i] = current
	}

	answer := make([]bool, len(queries))

	for i, query := range queries {
		answer[i] = component[query[0]] == component[query[1]]
	}

	return answer
}
