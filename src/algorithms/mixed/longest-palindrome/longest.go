package mixed

func longestPalindrome(input string) int {
	count := make(map[rune]int)

	for _, v := range input {
		count[v]++
	}

	size := 0
	hasOdd := false
	for _, v := range count {
		if isEven(v) {
			size += v
		} else {
			hasOdd = true
			size += v - 1
		}
	}

	if hasOdd {
		size++
	}

	return size
}

func isEven(n int) bool {
	return n%2 == 0
}
