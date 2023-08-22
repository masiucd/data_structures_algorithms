export function solution(a: string, b: string): number {
  let distance = 0
  for (let i = 0; i < a.length; i++) {
    const charFromA = a[i]
    const charFromB = b[i]
    if (charFromA !== charFromB) {
      distance++
    }
  }
  return distance
}
