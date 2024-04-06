export function twoSum(numbers: number[], target: number): number[] {
  let left = 0;
  let right = numbers.length - 1;
  while (left <= right) {
    let tempTarget = numbers[left] + numbers[right];
    if (tempTarget === target) return [left + 1, right + 1];
    if (tempTarget < target) {
      left++;
    } else {
      right--;
    }
  }
  return [];
}
