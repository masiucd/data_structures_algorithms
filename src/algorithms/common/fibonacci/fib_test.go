package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacciWithMemo(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
	}

	memo := make(map[int]int)
	for _, test := range tests {
		assert.Equal(t, test.expected, fibonacciWithMemo(test.n, memo))
	}
}

func TestFibonacciIterative(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, fibonacciIterative(test.n))
	}
}

func TestFibonacciRecursive(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, fibRecursive(test.n))
	}
}

// BenchmarkFibonacciWithMemo benchmarks the performance of the fibonacciWithMemo function.
// It calculates the 10th Fibonacci number repeatedly using memoization.
//
// The benchmark:
// - Creates a memoization map to store previously calculated Fibonacci numbers.
// - Runs the fibonacciWithMemo function with n=10 for b.N iterations.
// - Measures the time taken for these iterations to assess the function's performance.
//
// This benchmark helps evaluate the efficiency of the memoized Fibonacci implementation,
// particularly for repeated calculations of the same Fibonacci number.
func BenchmarkFibonacciWithMemo(b *testing.B) {
	memo := make(map[int]int)
	for i := 0; i < b.N; i++ {
		fibonacciWithMemo(10, memo)
	}
}

func BenchmarkFibonacciIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciIterative(10)
	}
}

func BenchmarkFibonacciRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibRecursive(10)
	}
}
