package numbers

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}

	if n%2 == 0 {
		return isUgly(n / 2)
	}
	if n%3 == 0 {
		return isUgly(n / 3)
	}
	if n%5 == 0 {
		return isUgly(n / 5)
	}
	return false
}

func isUglyV2(num int) bool {
	if num <= 0 {
		return false
	}
	for _, i := range []int{2, 3, 5} {
		for num%i == 0 {
			num /= i
		}
	}
	return num == 1
}
