function missingNumber(nums: number[]): number {
  let missing = -1;
  let numsSet = new Set([...nums]);
  for (let i = 0; i <= nums.length; i++) {
    if (!numsSet.has(i)) {
      missing = i;
    }
  }
  return missing;
}

let nums = [3, 0, 1];
let res = missingNumber(nums);
console.log("res", res); // 2

nums = [0, 1];
res = missingNumber(nums);
console.log("res", res); // 2

nums = [9, 6, 4, 2, 3, 5, 7, 0, 1];
res = missingNumber(nums);
console.log("res", res); // 8
