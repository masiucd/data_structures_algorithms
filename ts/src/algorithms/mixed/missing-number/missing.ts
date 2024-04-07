export function missingNumber(nums: number[]): number {
  let missing = -1;
  let numsSet = new Set([...nums]);
  for (let i = 0; i <= nums.length; i++) {
    if (!numsSet.has(i)) {
      missing = i;
    }
  }
  return missing;
}
