package weightedwordmapping

func mapWordWeights(words []string, weights []int) string {
	res := make([]byte, len(words))

	for i := 0; i < len(words); i++ {
		word := words[i]
		sum := 0

		for j := 0; j < len(word); j++ {
			idx := int(word[j] - 'a')
			sum += weights[idx]
		}

		res[i] = byte(int('z') - (sum % 26))
	}

	return string(res)
}
