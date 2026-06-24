package numberofzigzagarraysii

const mod int64 = 1_000_000_007

func zigZagArrays(n int, l int, r int) int {
	q := r - l

	if q == 1 {
		return 2
	}

	exponent := int64((n - 2) / 2)
	oddLength := n&1 == 1

	initial := buildInitialSequence(q, oddLength)

	if exponent < int64(q) {
		return int(2 * initial[int(exponent)] % mod)
	}

	recurrence := buildRecurrence(q)

	oneDirection := nthLinearRecurrence(
		initial,
		recurrence,
		exponent,
	)

	return int(2 * oneDirection % mod)
}

func buildInitialSequence(q int, oddLength bool) []int64 {
	current := make([]int64, q)
	next := make([]int64, q)
	sequence := make([]int64, q)

	for i := 0; i < q; i++ {
		current[i] = int64(i + 1)
	}

	for step := 0; step < q; step++ {
		var sum int64

		if oddLength {
			for i, value := range current {
				sum = (sum + int64(i+1)*value%mod) % mod
			}
		} else {
			for _, value := range current {
				sum += value

				if sum >= mod {
					sum -= mod
				}
			}
		}

		sequence[step] = sum

		if step+1 < q {
			multiplyMinMatrixVector(current, next)
			current, next = next, current
		}
	}

	return sequence
}

func multiplyMinMatrixVector(x []int64, y []int64) {
	var total int64

	for _, value := range x {
		total += value

		if total >= mod {
			total -= mod
		}
	}

	var prefixSum int64
	var prefixWeighted int64

	for i := 1; i <= len(x); i++ {
		value := x[i-1]

		prefixSum += value

		if prefixSum >= mod {
			prefixSum -= mod
		}

		prefixWeighted = (prefixWeighted + int64(i)*value%mod) % mod

		suffixSum := total - prefixSum

		if suffixSum < 0 {
			suffixSum += mod
		}

		y[i-1] = (prefixWeighted + int64(i)*suffixSum%mod) % mod
	}
}

func buildRecurrence(q int) []int64 {
	limit := 2 * q

	factorial := make([]int64, limit+1)
	inverseFactorial := make([]int64, limit+1)
	recurrence := make([]int64, q)

	factorial[0] = 1

	for i := 1; i <= limit; i++ {
		factorial[i] =
			factorial[i-1] * int64(i) % mod
	}

	inverseFactorial[limit] =
		modPow(factorial[limit], mod-2)

	for i := limit; i > 0; i-- {
		inverseFactorial[i-1] =
			inverseFactorial[i] * int64(i) % mod
	}

	for j := 0; j < q; j++ {
		n := q + 1 + j
		k := q - 1 - j

		coefficient :=
			factorial[n] *
				inverseFactorial[k] %
				mod

		coefficient =
			coefficient *
				inverseFactorial[n-k] %
				mod

		if j&1 == 1 && coefficient != 0 {
			coefficient = mod - coefficient
		}

		recurrence[j] = coefficient
	}

	return recurrence
}

func nthLinearRecurrence(
	initial []int64,
	recurrence []int64,
	n int64,
) int64 {
	q := len(initial)

	if n < int64(q) {
		return initial[int(n)]
	}

	result := make([]int64, q)
	base := make([]int64, q)

	result[0] = 1
	base[1] = 1

	for n > 0 {
		if n&1 == 1 {
			result = multiplyAndReduce(
				result,
				base,
				recurrence,
				false,
			)
		}

		n >>= 1

		if n > 0 {
			base = multiplyAndReduce(
				base,
				base,
				recurrence,
				true,
			)
		}
	}

	var answer int64

	for i := 0; i < q; i++ {
		answer = (answer + result[i]*initial[i]%mod) % mod
	}

	return answer
}

func multiplyAndReduce(
	a []int64,
	b []int64,
	recurrence []int64,
	square bool,
) []int64 {
	q := len(recurrence)
	temporary := make([]int64, 2*q-1)

	if square {
		for i := 0; i < q; i++ {
			if a[i] == 0 {
				continue
			}

			temporary[2*i] += a[i] * a[i] % mod

			if temporary[2*i] >= mod {
				temporary[2*i] -= mod
			}

			for j := i + 1; j < q; j++ {
				term :=
					(2 * a[i] % mod) *
						a[j] %
						mod

				temporary[i+j] += term

				if temporary[i+j] >= mod {
					temporary[i+j] -= mod
				}
			}
		}
	} else {
		for i := 0; i < q; i++ {
			if a[i] == 0 {
				continue
			}

			for j := 0; j < q; j++ {
				temporary[i+j] +=
					a[i] * b[j] % mod

				if temporary[i+j] >= mod {
					temporary[i+j] -= mod
				}
			}
		}
	}

	for degree := 2*q - 2; degree >= q; degree-- {
		coefficient := temporary[degree]

		if coefficient == 0 {
			continue
		}

		for j := 1; j <= q; j++ {
			index := degree - j

			temporary[index] +=
				coefficient *
					recurrence[j-1] %
					mod

			if temporary[index] >= mod {
				temporary[index] -= mod
			}
		}
	}

	result := make([]int64, q)
	copy(result, temporary[:q])

	return result
}

func modPow(base int64, exponent int64) int64 {
	result := int64(1)

	for exponent > 0 {
		if exponent&1 == 1 {
			result = result * base % mod
		}

		base = base * base % mod
		exponent >>= 1
	}

	return result
}
