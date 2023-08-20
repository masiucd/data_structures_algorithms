package common

func Factorial(n int) int {
	if n <= 2 {
		return n
	}
	return n * Factorial(n-1)
}

func FactorialIterator(n int) int {
	var num = n
	var res = 1
	for num > 0 {
		res *= num
		num--
	}
	return res
}
