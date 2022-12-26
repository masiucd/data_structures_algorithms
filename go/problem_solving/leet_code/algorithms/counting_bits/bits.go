package countingbits

func countBits(n int) []int {
	var bits []int
	for i := 0; i <= n; i++ {
		bits = append(bits, amountOfBits(uint(i)))
	}
	return bits
}

func amountOfBits(n uint) int {
	var count int
	for n > 0 {
		if n&1 == 1 {
			count++
		}
		n = n >> 1
	}
	return count
}
