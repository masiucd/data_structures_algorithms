export function solution(n: number): number {
  let count = 0
  while (n > 0) {
    if ((n & 1) === 1) {
      count++
    }
    n >>>= 1
  }
  return count
}

export function solutionTwo(n: number): number {
  return n
    .toString(2)
    .split("0")
    .filter(Boolean)
    .map(x => x.length)
    .reduce((a, b) => a + b, 0)
}

export function solutionThree(n: number): number {
  const bin = n.toString(2)
  let count = 0
  for (const b of bin) {
    if (b === "1") {
      count++
    }
  }
  return count
}
