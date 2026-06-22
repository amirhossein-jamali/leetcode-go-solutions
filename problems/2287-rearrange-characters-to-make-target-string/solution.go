package rearrangecharacterstomaketargetstring

func rearrangeCharacters(s string, target string) int {
	var sourceCount [26]int
	var targetCount [26]int

	for i := 0; i < len(s); i++ {
		sourceCount[s[i]-'a']++
	}

	for i := 0; i < len(target); i++ {
		targetCount[target[i]-'a']++
	}

	result := len(s) / len(target)

	for i := 0; i < 26; i++ {
		if targetCount[i] == 0 {
			continue
		}

		copies := sourceCount[i] / targetCount[i]

		if copies < result {
			result = copies
		}
	}

	return result
}
