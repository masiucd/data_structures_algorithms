import {TreeNode} from "../TreeNode";

export function searchBST(root: TreeNode | null, val: number): TreeNode | null {
  if (root === null) return null;
  if (root.val === val) return root;
  let left = searchBST(root.left, val);
  if (left) return left;
  let right = searchBST(root.right, val);
  if (right) return right;
  return null;
}

export function searchBST2(
  root: TreeNode | null,
  val: number
): TreeNode | null {
  if (root === null || root.val === val) return root;
  return val < root.val
    ? searchBST(root.left, val)
    : searchBST(root.right, val);
}
