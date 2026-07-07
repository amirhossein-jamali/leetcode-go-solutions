package concatenatenonzerodigitsandmultiplybysumi

func sumAndMultiply(n int) int64 {
	var x, sum, place int64 = 0, 0, 1

	for n > 0 {
		if digit := n % 10; digit != 0 {
			d := int64(digit)
			x += d * place
			sum += d
			place *= 10
		}

		n /= 10
	}

	return x * sum
}
