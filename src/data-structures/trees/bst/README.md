# Binary Search Tree

## Definition

A binary search tree is a binary tree where each node has values greater than all nodes in its left subtree and less than all nodes in its right subtree.

## Where are Binary search trees used in real life?

- Binary search trees are used in many search applications where data is constantly entering/leaving, such as the map and set objects in many languages' libraries.
- Binary search trees provide O(log n) time complexity for insertion, deletion, and search operations. This is because the height of the tree is always log n for a balanced tree.

## Mental model

A binary search tree is a tree where each node has a value greater than all nodes in its left subtree and less than all nodes in its right subtree.
Imagine a tree where each node has a value and two children. The left child is always less than the parent and the right child is always greater than the parent.
This is a binary search tree.

## Methods

- Insert(value int)
- Search(value int) bool
- Remove(value int)
- Size() int
- Traverse()
- Find(value int) \*Node
- Contains(value int) bool
- BFS() []int
- DFS(order string) []int // in, pre, post
- Min() int
- Max() int

## Time Complexity

|  Access   |  Search   | Insertion | Deletion  |
| :-------: | :-------: | :-------: | :-------: |
| O(log(n)) | O(log(n)) | O(log(n)) | O(log(n)) |

## Space Complexity

O(n) where n is the number of nodes in the tree.

## Resources

[Wikipedia](https://en.wikipedia.org/wiki/Binary_search_tree)
