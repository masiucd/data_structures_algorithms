package mcd_math

import "math"

func sumPowersOf2(n int) (result int) {
	for i := 0; i < n; i++ {
		result += int(math.Pow(2, float64(i+1)))
	}
	return result
}
