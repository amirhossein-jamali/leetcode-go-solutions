package maximumnumberofballoons

func maxNumberOfBalloons(text string) int {
	var b, a, l, o, n int

	for i := 0; i < len(text); i++ {
		switch text[i] {
		case 'b':
			b++
		case 'a':
			a++
		case 'l':
			l++
		case 'o':
			o++
		case 'n':
			n++
		}
	}

	l /= 2
	o /= 2

	result := b

	if a < result {
		result = a
	}
	if l < result {
		result = l
	}
	if o < result {
		result = o
	}
	if n < result {
		result = n
	}

	return result
}
