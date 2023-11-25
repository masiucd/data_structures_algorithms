package climbstairs

// bottom up approach
func ClimbStairs(n int) int {
	if n == 1 {
		return 1
	}
	first := 1
	second := 2
	for i := 3; i <= n; i++ {
		temp := first + second
		first = second
		second = temp
	}
	return second
}

func ClimbingStairsWithFib(n int) int {
	cache := make(map[int]int)
	var fn func(int) int
	fn = func(n int) int {
		if n <= 1 {
			return 1
		}
		// if not in cache, add it
		if _, ok := cache[n]; !ok {
			cache[n] = fn(n-1) + fn(n-2)
		}
		// return from cache
		return cache[n]
	}
	return fn(n)
}
