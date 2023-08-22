export function exponent(a: number, n: number): number {
  if (n === 0) return 1
  if (n === 1) return a
  if (n % 2 === 0)
    return exponent(a, Math.floor(n / 2)) * exponent(a, Math.floor(n / 2))
  return exponent(a, Math.floor(n / 2)) * exponent(a, Math.floor(n / 2)) * a
}
