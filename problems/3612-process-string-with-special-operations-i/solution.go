package processstringwithspecialoperationsi

func processStr(s string) string {
	result := []byte{}

	for i := 0; i < len(s); i++ {
		ch := s[i]

		switch ch {
		case '*':
			if len(result) > 0 {
				result = result[:len(result)-1]
			}

		case '#':
			result = append(result, result...)

		case '%':
			for left, right := 0, len(result)-1; left < right; left, right = left+1, right-1 {
				result[left], result[right] = result[right], result[left]
			}

		default:
			result = append(result, ch)
		}
	}

	return string(result)
}
