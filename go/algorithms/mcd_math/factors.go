package mcd_math

func factors(n int) []int {
	var xs []int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			xs = append(xs, i)
		}
	}
	return xs
}
