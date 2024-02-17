package mixed

// if we store the count of each character in a map
// we can then check if the count of each character is even
// if there is more than one character with an odd count
// then it is not possible to rearrange the string to form a palindrome
// if there is only one character with an odd count
// then it is possible to rearrange the string to form a palindrome
// if there are no characters with an odd count
// then the string is already a palindrome
// so we can return true
func palindromeRearranging(inputString string) bool {
	charCount := make(map[rune]int)
	for _, c := range inputString {
		charCount[c]++
	}

	var oddCount = 0
	for _, v := range charCount {
		if v%2 != 0 {
			oddCount++
		}
	}

	return oddCount <= 1
}
