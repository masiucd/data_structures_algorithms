function count(n: number, result: number[] = []): number[] {
  if (n <= 0) return result
  result.push(n)
  return count(n - 1, result)
}

console.log(count(10))
