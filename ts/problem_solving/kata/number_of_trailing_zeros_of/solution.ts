// https://mathworld.wolfram.com/Factorial.html
export function zeros(n: number): number {
  let cnt = 0
  while (n > 0) {
    n = Math.floor(n / 5)
    cnt += n
  }
  return cnt
}
