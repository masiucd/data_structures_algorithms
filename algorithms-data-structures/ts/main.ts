import {BinarySearchTree} from "@/data-structures/tree/bst/tree";

let bst = new BinarySearchTree();
bst.insert(10);
bst.insert(8);
bst.insert(16);
bst.insert(5);
bst.insert(7);
bst.insert(32);
bst.insert(12);
bst.insert(20);

console.log(bst.bfs());
console.log(bst.min());
console.log(bst.max());
console.log(bst.dfs("POST"));
console.log(bst.find(5));
