export function removeElement(nums: number[], val: number): number {
  const result = nums.filter((num) => num !== val);
  nums.length = result.length;
  for (let i = 0; i < result.length; i++) {
    nums[i] = result[i];
  }
  return result.length;
}
