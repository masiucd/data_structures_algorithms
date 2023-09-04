export function sumRecursive(nums: number[]): number {
  if (nums.length === 0) return 0;
  let head = nums[0];
  let tail = nums.slice(1);
  return head + sumRecursive(tail);
}

export function sum(nums: number[]): number {
  let total = 0;
  for (let i = 0; i < nums.length; i++) {
    total += nums[i];
  }
  return total;
}

export function sumReduce(nums: number[]): number {
  return nums.reduce((total, num) => total + num, 0);
}
