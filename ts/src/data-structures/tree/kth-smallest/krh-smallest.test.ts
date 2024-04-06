import {test, expect} from "bun:test";
import {kthSmallest} from "./kth-smallest";
import {TreeNode} from "../TreeNode";

test("kth-smallest", () => {
  let root = new TreeNode(3);
  root.left = new TreeNode(1);
  root.right = new TreeNode(4);
  root.left.right = new TreeNode(2);
  expect(kthSmallest(root, 1)).toBe(1);

  root = new TreeNode(5);
  root.left = new TreeNode(3);
  root.right = new TreeNode(6);
  root.left.left = new TreeNode(2);
  root.left.right = new TreeNode(4);
  root.left.left.left = new TreeNode(1);
  expect(kthSmallest(root, 3)).toBe(3);
});
