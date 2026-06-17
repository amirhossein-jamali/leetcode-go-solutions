package processstringwithspecialoperationsii

func processStr(s string, k int64) byte {
	var length int64

	for i := 0; i < len(s); i++ {
		ch := s[i]

		switch {
		case ch >= 'a' && ch <= 'z':
			length++
		case ch == '*':
			if length > 0 {
				length--
			}
		case ch == '#':
			length *= 2
		case ch == '%':
		}
	}

	if k >= length {
		return '.'
	}

	for i := len(s) - 1; i >= 0; i-- {
		ch := s[i]

		switch {
		case ch >= 'a' && ch <= 'z':
			if k == length-1 {
				return ch
			}
			length--

		case ch == '*':
			length++

		case ch == '#':
			half := length / 2
			k %= half
			length = half

		case ch == '%':
			k = length - 1 - k
		}
	}

	return '.'
}
