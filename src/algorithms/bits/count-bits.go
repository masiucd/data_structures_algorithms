package bits

import "strconv"

func countBits(n int) []int {
	var result []int
	for i := 0; i <= n; i++ {
		result = append(result, getAmountOfBits(i))
	}
	return result
}
func countBitsV2(n int) []int {
	var result []int
	for i := 0; i <= n; i++ {
		result = append(result, getAmountOfBitsV2(i))
	}
	return result
}

// getAmountOfBits returns the number of bits in a number
func getAmountOfBits(n int) int {
	var res int
	if n == 0 {
		return res
	}
	binary := strconv.FormatInt(int64(n), 2)

	for _, v := range binary {
		if string(v) == "1" {
			res++
		}
	}
	return res
}

// getAmountOfBitsV2 uses bitwise operators to count the number of bits
func getAmountOfBitsV2(n int) int {
	var res int
	if n == 0 {
		return res
	}
	for n > 0 {
		if n&1 == 1 {
			res++
		}
		n >>= 1
	}
	return res
}
