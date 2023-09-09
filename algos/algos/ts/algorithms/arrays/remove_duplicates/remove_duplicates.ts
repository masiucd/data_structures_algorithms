export function removeDuplicates(nums: number[]): number {
  return [...new Set(nums)].length;
}
export function removeDuplicatesRec(
  nums: number[],
  store: Set<number>
): number {
  if (nums.length === 0) return store.size;
  let head = nums[0];
  let tail = nums.slice(1);
  store.add(head);
  return removeDuplicatesRec(tail, store);
}
