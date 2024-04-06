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

export function insertionSortUsingSwapHelper(nums: number[]): number[] {
  for (let i = 1; i < nums.length; i++) {
    let j = i;
    while (j > 0 && nums[j - 1] > nums[j]) {
      swap(nums, j - 1, j);
      j--;
    }
  }

  return nums;
}

function swap(nums: number[], i: number, j: number) {
  // swap using array destructuring
  [nums[i], nums[j]] = [nums[j], nums[i]];
}
