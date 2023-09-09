package mcd_math

import "math"

func isPrime(n int) bool {
	if n == 1 {
		return false
	}
	if n == 2 {
		return true
	}
	for i := 2; i <= int(math.Floor(float64(n)/2)); i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}
func isPrimeTwo(n int) bool {
	if n == 1 {
		return false
	}
	if n == 2 {
		return true
	}
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(n)))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}
