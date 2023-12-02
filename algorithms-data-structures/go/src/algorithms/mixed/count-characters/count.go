package countcharacters

func countCharacters(words []string, chars string) int {
	var result int
	count := wordToCharMap(chars)
	for _, word := range words {
		currentWord := wordToCharMap(word)
		isOk := true
		for _, ch := range word {
			if _, ok := count[string(ch)]; !ok || currentWord[string(ch)] > count[string(ch)] {
				isOk = false
				break
			}
		}
		if isOk {
			result += len(word)
		}
	}
	return result
}

func wordToCharMap(word string) map[string]int {
	charMap := make(map[string]int)
	for _, ch := range word {
		if _, ok := charMap[string(ch)]; ok {
			charMap[string(ch)]++
		} else {
			charMap[string(ch)] = 1

		}
	}
	return charMap
}
