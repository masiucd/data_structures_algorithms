export function zipWith(
  fn: (n: number, b: number) => number,
  a0: number[],
  a1: number[]
): number[] {
  const result: number[] = []
  const shortestLength = Math.min(a0.length, a1.length)
  for (let i = 0; i < shortestLength; i++) {
    result.push(fn(a0[i], a1[i]))
  }
  return result
}
