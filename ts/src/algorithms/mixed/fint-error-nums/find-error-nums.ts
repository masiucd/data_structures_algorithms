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
