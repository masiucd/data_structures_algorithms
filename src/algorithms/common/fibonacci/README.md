# Fibonacci

## Recursion with Memoization

### Explanation

The `fibonacciWithMemo` function calculates the nth Fibonacci number using recursion and memoization.

### Steps

1. **Base Cases**:

   - If `n` is 0, return 0.
   - If `n` is 1, return 1.

2. **Memoization Check**:

   - Before performing the recursive calculation, check if the result for `n` is already stored in the `memo` map.
   - If it is, return the stored result to avoid redundant calculations.
   - If it is not, proceed with the recursive calculation.

3. **Recursive Calculation**:
   - Calculate the nth Fibonacci number recursively by calling `fibonacciWithMemo(n-1, memo)` and `fibonacciWithMemo(n-2, memo)`.
   - Store the result of the recursive calculation in the `memo` map.
   - Return the result.

### Example

For `n = 5`, the function calculates:

```
fibonacciWithMemo(5, memo) = fibonacciWithMemo(4, memo) + fibonacciWithMemo(3, memo)
fibonacciWithMemo(4, memo) = fibonacciWithMemo(3, memo) + fibonacciWithMemo(2, memo)
fibonacciWithMemo(3, memo) = fibonacciWithMemo(2, memo) + fibonacciWithMemo(1, memo)
fibonacciWithMemo(2, memo) = fibonacciWithMemo(1, memo) + fibonacciWithMemo(0, memo)
fibonacciWithMemo(1, memo) = 1
fibonacciWithMemo(0, memo) = 0
```

### Performance

The `fibonacciWithMemo` function has a time complexity of O(n) and a space complexity of O(n).

### Benchmark

The `BenchmarkFibonacciWithMemo` function is a benchmark for the `fibonacciWithMemo` function. It calculates the 10th Fibonacci number repeatedly using memoization.

The benchmark:

- Creates a memoization map to store previously calculated Fibonacci numbers.
- Runs the `fibonacciWithMemo` function with `n=10` for `b.N` iterations.
- Measures the time taken for these iterations to assess the function's performance.

### Mental Model

The mental model for the `fibonacciWithMemo` function is to think of it as a tree of recursive calls, where each node represents a call to the function. The function starts at the root of the tree and calculates the Fibonacci number for the current node. It then recursively calculates the Fibonacci numbers for the children of the current node. The function continues this process until it reaches the base cases.

The function uses a memoization map to store the results of the recursive calculations. This map is passed between recursive calls to store the results of the calculations.

The function uses a recursive approach to calculate the Fibonacci number for the current node. It then stores the result of the calculation in the memoization map.

The function continues this process until it reaches the base cases.
