export function solution<T>(list: T[], k: number): T[][] {
  const result: T[][] = []
  for (let i = 0; i < list.length; i += k) {
    const xs = list.slice(i, i + k)
    result.push(xs)
  }
  return result
}
