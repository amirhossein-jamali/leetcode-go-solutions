package concatenatenonzerodigitsandmultiplybysumii

func sumAndMultiply(s string, queries [][]int) []int {
	const MOD = 1_000_000_007
	const SHIFT = 20
	const MASK uint64 = (1 << SHIFT) - 1

	m := len(s)

	meta := make([]uint64, m+1)

	val := make([]int, m+1)

	pow10 := make([]int, m+1)
	pow10[0] = 1

	cnt := 0
	sum := 0
	cur := 0

	for i := 0; i < m; i++ {
		d := int(s[i] - '0')

		if d != 0 {
			cnt++
			sum += d
			cur = (cur*10 + d) % MOD
			pow10[cnt] = (pow10[cnt-1] * 10) % MOD
		}

		meta[i+1] = (uint64(cnt) << SHIFT) | uint64(sum)
		val[i+1] = cur
	}

	ans := make([]int, len(queries))

	for i := 0; i < len(queries); i++ {
		l := queries[i][0]
		r1 := queries[i][1] + 1

		leftMeta := meta[l]
		rightMeta := meta[r1]

		length := int(rightMeta>>SHIFT) - int(leftMeta>>SHIFT)
		if length == 0 {
			continue
		}

		x := val[r1] - (val[l]*pow10[length])%MOD
		if x < 0 {
			x += MOD
		}

		digitSum := int((rightMeta & MASK) - (leftMeta & MASK))

		ans[i] = (x * digitSum) % MOD
	}

	return ans
}
