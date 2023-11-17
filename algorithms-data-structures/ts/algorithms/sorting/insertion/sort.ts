export function insertionSort(nums: number[]): number[] {
  for (let i = 1; i < nums.length; i++) {
    let j = i;
    while (j >= 0 && nums[j - 1] > nums[j]) {
      // swap
      let temp = nums[j - 1];
      nums[j - 1] = nums[j];
      nums[j] = temp;
      j--;
    }
  }

  return nums;
}
