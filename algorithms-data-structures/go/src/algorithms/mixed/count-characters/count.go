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

func countCharactersTwo(words []string, chars string) int {
	var result int
	var count [26]int
	for _, ch := range chars {
		count[ch-'a']++
	}
	for _, word := range words {
		// make a copy of count
		currentWord := count
		// check if word is ok
		isOk := true
		for _, ch := range word {
			// if currentWord[ch-'a'] > 0, then decrement it
			if currentWord[ch-'a'] > 0 {
				currentWord[ch-'a']--
			} else {
				// if currentWord[ch-'a'] == 0, then break
				isOk = false
				break
			}
		}
		if isOk {
			// if word is ok, then add its length to result
			result += len(word)
		}
	}
	return result
}
