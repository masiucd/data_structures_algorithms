/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     val: number
 *     left: TreeNode | null
 *     right: TreeNode | null
 *     constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.left = (left===undefined ? null : left)
 *         this.right = (right===undefined ? null : right)
 *     }
 * }
 */

import {TreeNode} from "../TreeNode";

export function hasPathSum(root: TreeNode | null, targetSum: number): boolean {
  if (root === null) return false;
  // If we are at a leaf node and the target sum is equal to the value of the leaf node, return true
  if (root.val === targetSum && root.left === null && root.right === null)
    return true;
  // Recursively call the function on the left and right nodes, subtracting the value of the current node from the target sum
  if (hasPathSum(root.left, targetSum - root.val)) return true;
  // Recursively call the function on the right node, subtracting the value of the current node from the target sum
  if (hasPathSum(root.right, targetSum - root.val)) return true;
  // If we reach this point, there is no path that sums to the target sum
  return false;
}
