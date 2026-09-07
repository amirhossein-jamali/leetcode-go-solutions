package distinctsubsequencesii

func distinctSubseqII(s string) int {
	const mod int64 = 1_000_000_007

	var end [26]int64
	var total int64

	for i := 0; i < len(s); i++ {
		c := s[i] - 'a'
		old := end[c]

		newEnd := total + 1
		if newEnd == mod {
			newEnd = 0
		}

		total = total*2 + 1 - old

		if total < 0 {
			total += mod
		} else if total >= mod {
			total -= mod
		}

		end[c] = newEnd
	}

	return int(total)
}
