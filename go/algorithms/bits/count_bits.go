package bits

//  AND
// n = 1 & 1

//  OR
// n = 1 | 0

//  XOR
// n = 0 ^ 1

//  NOT (negation)
// n = ~n

//  Bit shifting
// n = 1
// n = n << 1
// n = n >> 1

func countBits(n uint) int {
	var count int
	for n > 0 {
		if n&1 == 1 {
			count++
		}
		n = n >> 1
	}
	return count
}
