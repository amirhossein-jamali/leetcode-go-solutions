package maximumscoreofnonoverlappingintervals

import "slices"

func insert3414(path uint64, x uint16, k int) uint64 {
	v := uint64(x)

	switch k {
	case 1:
		return v << 48

	case 2:
		a := uint16(path >> 48)

		if x < a {
			return v<<48 | uint64(a)<<32
		}

		return uint64(a)<<48 | v<<32

	case 3:
		a := uint16(path >> 48)
		b := uint16(path >> 32)

		if x < a {
			return v<<48 |
				uint64(a)<<32 |
				uint64(b)<<16
		}

		if x < b {
			return uint64(a)<<48 |
				v<<32 |
				uint64(b)<<16
		}

		return uint64(a)<<48 |
			uint64(b)<<32 |
			v<<16

	default:
		a := uint16(path >> 48)
		b := uint16(path >> 32)
		c := uint16(path >> 16)

		if x < a {
			return v<<48 |
				uint64(a)<<32 |
				uint64(b)<<16 |
				uint64(c)
		}

		if x < b {
			return uint64(a)<<48 |
				v<<32 |
				uint64(b)<<16 |
				uint64(c)
		}

		if x < c {
			return uint64(a)<<48 |
				uint64(b)<<32 |
				v<<16 |
				uint64(c)
		}

		return uint64(a)<<48 |
			uint64(b)<<32 |
			uint64(c)<<16 |
			v
	}
}

func maximumWeight(intervals [][]int) []int {
	n := len(intervals)

	data := make([]uint64, n)
	order := make([]uint64, n)

	for i, v := range intervals {
		data[i] =
			uint64(uint32(v[0]))<<32 |
				uint64(uint32(v[2]))

		order[i] =
			uint64(uint32(v[1]))<<32 |
				uint64(uint16(i))
	}

	slices.Sort(order)

	for i := 0; i < n; i++ {
		id := uint16(order[i])
		left := uint32(data[id] >> 32)

		lo, hi := 0, i

		for lo < hi {
			m := int(uint(lo+hi) >> 1)

			if uint32(order[m]>>32) < left {
				lo = m + 1
			} else {
				hi = m
			}
		}

		order[i] |= uint64(uint16(lo)) << 16
	}

	for i, x := range order {
		id := uint16(x)
		w := uint32(data[id])

		order[i] =
			uint64(w)<<32 |
				uint64(uint32(x))
	}

	m := n + 1

	scores := make([]uint32, m<<1)
	paths := make([]uint64, m<<1)

	prevS := scores[:m]
	currS := scores[m:]

	prevP := paths[:m]
	currP := paths[m:]

	var ansScore uint32
	var ansPath uint64
	ansK := 0

	for k := 1; k <= 4; k++ {
		currS[0] = 0
		currP[0] = 0

		for i := 1; i <= n; i++ {
			x := order[i-1]

			w := uint32(x >> 32)
			p := int(uint16(x >> 16))
			id := uint16(x)

			bestS := currS[i-1]
			bestP := currP[i-1]

			if k == 1 || prevS[p] != 0 {
				s := prevS[p] + w
				q := insert3414(
					prevP[p],
					id+1,
					k,
				)

				if s > bestS ||
					s == bestS && q < bestP {

					bestS = s
					bestP = q
				}
			}

			currS[i] = bestS
			currP[i] = bestP
		}

		s := currS[n]
		p := currP[n]

		if s > ansScore ||
			s == ansScore &&
				(ansK == 0 || p < ansPath) {

			ansScore = s
			ansPath = p
			ansK = k
		}

		prevS, currS = currS, prevS
		prevP, currP = currP, prevP
	}

	ans := make([]int, ansK)

	for i := 0; i < ansK; i++ {
		ans[i] =
			int(uint16(ansPath>>uint(48-16*i))) - 1
	}

	return ans
}
