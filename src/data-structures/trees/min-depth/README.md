# Min depth of a binary tree

Given a binary tree, find its minimum depth.
The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

## Example

```
Input: [3,9,20,null,null,15,7]
Output: 2
```

## Constraints

- The number of nodes in the tree is in the range [0, 10^5].
- -1000 <= Node.val <= 1000
- The depth of the tree will not exceed 1000.
- The tree is guaranteed to be a binary tree.

## Solution

The solution is to use a breadth-first search algorithm to traverse the tree and return the depth of the first leaf node found.

The algorithm is as follows:

1. If the root is null, return 0.
2. Initialize a queue with the root node and a depth of 1.
3. While the queue is not empty:
   1. Dequeue a node and its depth.
   2. If the node is a leaf node, return the depth.
   3. If the node has a left child, enqueue the left child with a depth of the current depth + 1.
   4. If the node has a right child, enqueue the right child with a depth of the current depth + 1.
   5. Repeat until the queue is empty.
   6. Return 0 if no leaf node is found.
   7. The time complexity of this algorithm is O(n), where n is the number of nodes in the tree.
   8. The space complexity of this algorithm is O(n), where n is the number of nodes in the tree.

## References

- [LeetCode](https://leetcode.com/problems/minimum-depth-of-binary-tree/)
- [GeeksforGeeks](https://www.geeksforgeeks.org/find-minimum-depth-of-a-binary-tree/)
