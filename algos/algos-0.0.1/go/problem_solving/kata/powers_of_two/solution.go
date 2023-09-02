package powers_of_two

import "math"

func solution(n int) (powers []uint64) {
	for i := 0; i <= n; i++ {
		powers = append(powers, uint64(math.Pow(2, float64(i))))
	}
	return
}
