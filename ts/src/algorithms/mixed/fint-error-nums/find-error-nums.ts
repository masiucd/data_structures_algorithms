export function findErrorNums(nums: number[]): number[] {
  let counter = new Map<number, number>();
  for (let n of nums) {
    if (!counter.has(n)) {
      counter.set(n, 1);
    } else {
      counter.set(n, counter.get(n)! + 1);
    }
  }
  let duplicate = -1;
  let missing = -1;
  for (let i = 1; i <= nums.length; i++) {
    let n = counter.get(i);
    if (n) {
      if (n === 2) {
        duplicate = i;
      }
    } else {
      missing = i;
    }
  }
  return [duplicate, missing];
}

export function findErrorNumsV2(nums: number[]): number[] {
  let xs = new Array(nums.length + 1).fill(0);
  for (let i = 0; i < nums.length; i++) {
    xs[nums[i]]++;
  }

  let duplicate = -1;
  let missing = -1;

  for (let i = 1; i < xs.length; i++) {
    //
    if (xs[i] === 2) {
      duplicate = i;
    } else if (xs[i] === 0) {
      missing = i;
    }
  }
  return [duplicate, missing];
}
