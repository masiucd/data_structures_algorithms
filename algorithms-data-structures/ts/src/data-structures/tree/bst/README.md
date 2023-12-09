# Binary search Tree

## Description

A binary search tree is a binary tree in which every node fits a specific ordering property:

- All left descendants <= n < all right descendants
- This must be true for each node n.

## Operations

- `insert(value)`: inserts a new value in the tree, recursively
- `insertIterative(value)`: inserts a new value in the tree, iteratively
- `find(value)`: returns the node with the given value
- `remove(value)`: removes the node with the given value
- `contains(value)`: returns true if the value exists in the tree, otherwise false
- `size()`: returns the number of nodes in the tree
- `depth()`: returns the maximum depth of the tree
- `min()`: returns the minimum value in the tree
- `max()`: returns the maximum value in the tree
- `isEmpty()`: returns true if the tree is empty, otherwise false
- `toArray()`: returns an array with all the values in the tree, in order
- `toString()`: returns a string representation of the tree
- `bfs()`: returns an array with all the values in the tree, in breadth-first order
- `dfs()`: returns an array with all the values in the tree, in depth-first order

![Binary Search Tree](https://upload.wikimedia.org/wikipedia/commons/thumb/d/da/Binary_search_tree.svg/1200px-Binary_search_tree.svg.png)
