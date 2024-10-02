package fibonacci

// fibRecursive calculates the nth Fibonacci number using a recursive approach.
// It has a time complexity of O(2^n) and a space complexity of O(n) due to the call stack.
//
// Parameters:
//   - n: The position of the Fibonacci number to calculate (0-indexed).
//
// Returns:
//   - The nth Fibonacci number.
//
// Example:
//
//	fibRecursive(5) returns 5 (0, 1, 1, 2, 3, 5)
func fibRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return fibRecursive(n-1) + fibRecursive(n-2)
}

// fibonacciWithMemo calculates the nth Fibonacci number using recursion with memoization.
// It has a time complexity of O(n) and a space complexity of O(n) due to the memoization map.
//
// Parameters:
//   - n: The position of the Fibonacci number to calculate (0-indexed).
//   - memo: A map to store previously calculated Fibonacci numbers.
//
// Returns:
//   - The nth Fibonacci number.
//
// Example:
//
//	memo := make(map[int]int)
//	fibonacciWithMemo(5, memo) returns 5 (0, 1, 1, 2, 3, 5)
func fibonacciWithMemo(n int, memo map[int]int) int {
	// If we have already calculated the result, return it
	if res, ok := memo[n]; ok {
		return res
	}
	// Base case
	if n <= 1 {
		return n
	}

	// Calculate and store the result in the memoization map
	memo[n] = fibonacciWithMemo(n-1, memo) + fibonacciWithMemo(n-2, memo)
	return memo[n]
}

// fibonacciIterative calculates the nth Fibonacci number using an iterative approach.
// It has a time complexity of O(n) and a space complexity of O(1).
//
// Parameters:
//   - n: The position of the Fibonacci number to calculate (0-indexed).
//
// Returns:
//   - The nth Fibonacci number.
//
// Example:
//
//	fibonacciIterative(5) returns 5 (0, 1, 1, 2, 3, 5)
func fibonacciIterative(n int) int {
	if n <= 1 {
		return n
	}
	prev, curr := 0, 1
	for i := 2; i <= n; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
}
