package mcd_math

func exponentByRecursion(a, n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return a
	}
	if n%2 == 0 {
		return exponentByRecursion(a, n/2) * exponentByRecursion(a, n/2)
	}
	return exponentByRecursion(a, n/2) * exponentByRecursion(a, n/2) * a
}
