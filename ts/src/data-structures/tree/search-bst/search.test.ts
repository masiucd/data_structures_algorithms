import {test, expect} from "bun:test";

import {searchBST, searchBST2, searchBSTV3} from "./search";
import {TreeNode} from "../TreeNode";

function nodesToArray(root: TreeNode | null): number[] {
  if (root === null) return [];
  return [root.val, ...nodesToArray(root.left), ...nodesToArray(root.right)];
}

test("searchBST", () => {
  let root = new TreeNode(4);
  root.left = new TreeNode(2);
  root.right = new TreeNode(7);
  root.left.left = new TreeNode(1);
  root.left.right = new TreeNode(3);

  expect(nodesToArray(searchBST(root, 2))).toEqual([2, 1, 3]);
  expect(nodesToArray(searchBST2(root, 2))).toEqual([2, 1, 3]);
  expect(nodesToArray(searchBSTV3(root, 2))).toEqual([2, 1, 3]);
});
