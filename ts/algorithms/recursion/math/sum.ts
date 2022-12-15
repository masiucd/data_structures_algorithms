export function sum(xs: number[]): number {
  if (xs.length === 0) return 0
  const head = xs[0]
  const tail = xs.slice(1)
  return head + sum(tail)
}
