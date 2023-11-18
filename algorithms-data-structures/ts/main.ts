const numbers = [100, 5, 1, 0, 21, 0, -4];

function insertionSort(nums: number[]): number[] {
  for (let i = 1; i < nums.length; i++) {
    let j = i;
    while (j > 0 && nums[j - 1] > nums[j]) {
      swap(nums, j - 1, j);
      j--;
    }
  }

  return nums;
}

console.log(insertionSort(numbers));

function swap(nums: number[], i: number, j: number) {
  // swap using array destructuring
  [nums[i], nums[j]] = [nums[j], nums[i]];
}
