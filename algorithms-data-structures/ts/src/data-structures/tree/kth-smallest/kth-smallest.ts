import {TreeNode} from "../TreeNode";

export function kthSmallest(root: TreeNode | null, k: number): number {
  let nums: number[] = [];
  dfsInOrder(root, nums);
  return k === 0 ? nums[0] : nums[k - 1];
}

function dfsInOrder(root: TreeNode | null, nums: number[]) {
  if (!root) return;
  dfsInOrder(root.left, nums);
  nums.push(root.val);
  dfsInOrder(root.right, nums);
}
