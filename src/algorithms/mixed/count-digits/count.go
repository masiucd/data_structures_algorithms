package countdigits

func countAmountOfDigits(n int) int {
	length, temp := 0, n
	for temp > 0 {
		length++
		temp /= 10
	}
	return length
}
