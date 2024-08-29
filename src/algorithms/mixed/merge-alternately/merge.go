package mixed

func mergeAlternately(word1, word2 string) string {
	var result string
	word, longestSize, shortestSize := word1, 0, 0
	if len(word2) > len(word1) {
		word = word2
		longestSize = len(word2)
		shortestSize = len(word1)
	} else {
		longestSize = len(word1)
		shortestSize = len(word2)
	}

	for i := 0; i < shortestSize; i++ {
		result += string(word1[i]) + string(word2[i])
	}

	if longestSize > shortestSize {
		result += word[shortestSize:]
	}

	return result

}
