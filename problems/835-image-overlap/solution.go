package imageoverlap

import "math/bits"

func largestOverlap(img1 [][]int, img2 [][]int) int {
	n := len(img1)

	var a, b [30]uint32
	var bp [31]int

	onesA, onesB := 0, 0

	for i := 0; i < n; i++ {
		var x, y uint32

		for j := 0; j < n; j++ {
			x |= uint32(img1[i][j]) << j
			y |= uint32(img2[i][j]) << j
		}

		a[i] = x
		b[i] = y

		onesA += bits.OnesCount32(x)
		onesB += bits.OnesCount32(y)
		bp[i+1] = onesB
	}

	if onesA == 0 || onesB == 0 {
		return 0
	}

	target := onesA
	if onesB < target {
		target = onesB
	}

	mask := uint32(1<<n) - 1

	var shifted [30]uint32
	var ap [31]int

	ans := 0

	for tx := 0; tx < 2*n-1; tx++ {
		dx := 0

		if tx > 0 {
			dx = (tx + 1) >> 1
			if tx&1 == 0 {
				dx = -dx
			}
		}

		ap[0] = 0

		if dx >= 0 {
			s := uint(dx)

			for i := 0; i < n; i++ {
				shifted[i] = (a[i] << s) & mask
				ap[i+1] = ap[i] + bits.OnesCount32(shifted[i])
			}
		} else {
			s := uint(-dx)

			for i := 0; i < n; i++ {
				shifted[i] = a[i] >> s
				ap[i+1] = ap[i] + bits.OnesCount32(shifted[i])
			}
		}

		if ap[n] <= ans {
			continue
		}

		for ty := 0; ty < 2*n-1; ty++ {
			dy := 0

			if ty > 0 {
				dy = (ty + 1) >> 1
				if ty&1 == 0 {
					dy = -dy
				}
			}

			ai, bi, rows := 0, 0, n

			if dy > 0 {
				bi = dy
				rows = n - dy
			} else if dy < 0 {
				ai = -dy
				rows = n + dy
			}

			ca := ap[ai+rows] - ap[ai]
			cb := bp[bi+rows] - bp[bi]

			if ca <= ans || cb <= ans {
				continue
			}

			cur := 0

			for k := 0; k < rows; k++ {
				cur += bits.OnesCount32(
					shifted[ai+k] & b[bi+k],
				)
			}

			if cur > ans {
				ans = cur

				if ans == target {
					return ans
				}
			}
		}
	}

	return ans
}
