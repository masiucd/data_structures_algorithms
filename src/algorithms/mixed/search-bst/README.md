# Search in a Binary Search Tree

## Description

Given the root of a binary search tree (BST) and an integer val, find the node in the BST that the node's value equals val and return the subtree rooted with that node. If such a node does not exist, return null.

## Examples

### Example 1

Input: root = [4,2,7,1,3], val = 2
Output: [2,1,3]

### Example 2

Input: root = [4,2,7,1,3], val = 5
Output: []

## Constraints

- The number of nodes in the tree is in the range [1, 5000].
- 1 <= Node.val <= 10^7
- root is a binary search tree.
- 1 <= val <= 10^7

## Approach

The problem can be solved using a recursive approach. The key idea is to compare the value of the current node with the target value. If the current node's value is greater than the target value, we search in the left subtree. If the current node's value is less than the target value, we search in the right subtree. If the current node's value is equal to the target value, we return the current node.

Here is the step-by-step approach:

1. If the current node is null, return null.
2. If the current node's value is equal to the target value, return the current node.
3. If the current node's value is greater than the target value, search in the left subtree.
4. If the current node's value is less than the target value, search in the right subtree.
5. Return null if the target value is not found in the tree.

This approach ensures that we efficiently search the binary search tree and return the correct subtree rooted at the node with the value equal to the target value.
