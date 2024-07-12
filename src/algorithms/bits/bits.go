package bits

import "fmt"

// BitwiseAnd checks if both bits are 1 if so it returns 1 else 0
func flip(n int) int {
	n ^= 0xFFFFFFFF
	return n
}

// If the first bit is 1 then the number is odd else it is even
// The bits that ends in 1 are odd numbers and the bits that ends in 0 are even numbers
func isEven(n int) bool {
	return n&1 == 0
}

func intToBinary(num int) string {
	return fmt.Sprintf("\n %d = %b \n", num, num)
}

func isPowerofTwo(n int) bool {
	if n&(n-1) == 0 {
		return true
	}
	return false
}
