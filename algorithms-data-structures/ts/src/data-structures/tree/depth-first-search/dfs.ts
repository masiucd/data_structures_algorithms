import {TreeNode} from "../TreeNode";

export function dfsUsingStack(root: TreeNode) {
  let stack = [root];
  let res = [];
  while (stack.length > 0) {
    let node = stack.pop();
    if (node) {
      res.push(node.val);
      if (node.right) stack.push(node.right);
      if (node.left) stack.push(node.left);
    }
  }
  return res;
}

export function dfsUsingRecursion(
  root: TreeNode | null,
  result: number[] = []
) {
  if (root === null) return;
  result.push(root.val);
  dfsUsingRecursion(root.left, result);
  dfsUsingRecursion(root.right, result);
}
