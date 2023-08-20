function fib(n: number): number {
  if (n <= 2) return 1
  return fib(n - 1) + fib(n - 2)
}

function fibMemo(n: number, memo: Record<number, number>): number {
  if (n <= 2) return 1
  if (memo[n]) {
    return memo[n]
  }
  const result = fibMemo(n - 1, memo) + fibMemo(n - 2, memo)
  memo[n] = result
  return result
}
