package numberofzigzagarraysi

const mod uint32 = 1_000_000_007

func zigZagArrays(n int, l int, r int) int {
	m := r - l + 1

	dp := make([]uint32, m)

	for i := range dp {
		dp[i] = 1
	}

	for length := 2; length <= n; length += 2 {
		var prefix uint32

		for value := 0; value < m; value++ {
			old := dp[value]
			dp[value] = prefix

			prefix += old
			if prefix >= mod {
				prefix -= mod
			}
		}

		if length == n {
			break
		}

		var suffix uint32

		for value := m - 1; value >= 0; value-- {
			old := dp[value]
			dp[value] = suffix

			suffix += old
			if suffix >= mod {
				suffix -= mod
			}
		}
	}

	var onePattern uint32

	for _, count := range dp {
		onePattern += count

		if onePattern >= mod {
			onePattern -= mod
		}
	}

	return int(uint64(onePattern) * 2 % uint64(mod))
}
