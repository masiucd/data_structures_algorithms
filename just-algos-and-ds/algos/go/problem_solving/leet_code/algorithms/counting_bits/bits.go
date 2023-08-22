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

func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res = res<<1 + num&1
		num >>= 1
	}
	return res
}
