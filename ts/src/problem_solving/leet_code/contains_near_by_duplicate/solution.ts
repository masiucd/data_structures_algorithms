export function containsNearbyDuplicate(nums: number[], k: number): boolean {
  for (let i = 0; i < nums.length; i++) {
    for (let j = i + 1; j < nums.length; j++) {
      if (nums[i] === nums[j] && Math.abs(i - j) <= k) {
        return true
      }
    }
  }
  return false
}

export function containsNearbyDuplicate2(nums: number[], k: number): boolean {
  const store = new Map<number, number>()
  for (const [i, num] of nums.entries()) {
    if (store.has(num)) {
      if (Math.abs(store.get(num)! - i) <= k) return true
    }
    store.set(num, i)
  }
  return false
}
