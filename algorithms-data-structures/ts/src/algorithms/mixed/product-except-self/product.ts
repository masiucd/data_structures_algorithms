export function productExceptSelf(nums: number[]): number[] {
  let result: number[] = [];
  let prefix = 1;
  let postfix = 1;
  for (let [i, n] of nums.entries()) {
    result[i] = prefix;
    prefix *= n;
  }

  for (let i = nums.length - 1; i >= 0; i--) {
    result[i] *= postfix;
    if (result[i] === -0) result[i] = 0;
    postfix *= nums[i];
  }

  return result;
}
