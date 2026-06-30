package numberofsubstringscontainingallthreecharacters

func numberOfSubstrings(s string) int {
	count := [3]int{}
	left := 0
	result := 0
	n := len(s)

	for right := 0; right < n; right++ {
		count[s[right]-'a']++

		for count[0] > 0 && count[1] > 0 && count[2] > 0 {
			result += n - right

			count[s[left]-'a']--
			left++
		}
	}

	return result
}
