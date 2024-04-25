package mixed

func reverseVowels(s string) string {
	var result string
	var vowelsInInput []rune
	for _, char := range s {
		if isVowel(char) {
			vowelsInInput = append(vowelsInInput, char)
		}
	}

	for _, char := range s {
		if isVowel(char) {
			last := vowelsInInput[len(vowelsInInput)-1]
			result += string(last)
			vowelsInInput = vowelsInInput[:len(vowelsInInput)-1]
		} else {
			result += string(char)
		}
	}
	return result
}

func isVowel(r rune) bool {
	for _, v := range []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'} {
		if r == v {
			return true
		}
	}
	return false
}
