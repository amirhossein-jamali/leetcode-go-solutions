package distinctsubsequences

func numDistinct(s string, t string) int {
	m := len(t)
	dp := make([]int, m+1)

	dp[0] = 1

	for i := 0; i < len(s); i++ {
		for j := m - 1; j >= 0; j-- {
			if s[i] == t[j] {
				dp[j+1] += dp[j]
			}
		}
	}

	return dp[m]
}
