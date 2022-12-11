package common

func fibonacciSequence(n int) []int {
	var result []int
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
		result = append(result, a)
	}
	return result
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func fibonacciMemo(n int, cache map[int]int) int {
	if n < 2 {
		return n
	}
	if _, ok := cache[n]; !ok {
		cache[n] = fibonacciMemo(n-1, cache) + fibonacciMemo(n-2, cache)
	}
	return cache[n]
}

func fibIter(n int) int {
	a, b := 1, 1
	var next int
	for i := 1; i < n; i++ {
		next = a + b
		a = b
		b = next
	}
	return a
}
