export function twoSum(nums: number[], target: number): number[] {
  const store: Map<number, number> = new Map()
  for (const [i, n] of nums.entries()) {
    if (store.has(target - n)) {
      const storedIndex = store.get(target - n) as number
      return [storedIndex, i]
    }
    store.set(n, i)
  }

  return [0, 0]
}
