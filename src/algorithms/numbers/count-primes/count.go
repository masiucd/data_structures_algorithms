package numbers

func countPrimes(n int) int {
	count := 0
	for i := 2; i < n; i++ {
		if isPrime(i) {
			count++
		}
	}

	return count
}

func isPrime(n int) bool {
	if n == 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
