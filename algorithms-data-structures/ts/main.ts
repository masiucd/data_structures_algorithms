import {TreeNode} from "@/data-structures/tree/TreeNode";
import {
  dfsUsingRecursion,
  dfsUsingStack,
} from "@/data-structures/tree/depth-first-search/dfs";

let node = new TreeNode(1);
node.left = new TreeNode(2);
node.right = new TreeNode(3);
node.left.left = new TreeNode(4);
node.left.right = new TreeNode(5);
node.right.left = new TreeNode(6);

let numbers: Array<number> = [];
console.log(dfsUsingStack(node));
console.log("----");
dfsUsingRecursion(node, numbers);
console.log(numbers);
