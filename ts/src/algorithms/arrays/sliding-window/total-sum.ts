export function totalSum(nums: number[], size: number): number {
  let tempSum = 0;
  let maxSum = 0;
  for (let i = 0; i < size; i++) {
    tempSum += nums[i];
  }
  maxSum = tempSum;

  for (let i = size; i < nums.length; i++) {
    tempSum = tempSum - nums[i - size] + nums[i];
    maxSum = Math.max(maxSum, tempSum);
  }
  return maxSum;
}
