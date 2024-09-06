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

func reverseVowelsV2(s string) string {
	var stack []rune
	result := []rune(s)
	vowels := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}

	for _, r := range s {
		if vowels[r] {
			stack = append(stack, r)
		}
	}

	for i, r := range s {
		if vowels[r] {
			// get the peek value
			popped := stack[len(stack)-1]
			// pop from the stack
			stack = stack[:len(stack)-1]
			// replace
			result[i] = popped
		}
	}
	return string(result)
}
