# Count nodes

Given a binary tree, count the number of nodes in the tree.

## Example

```go
tree := &Node{
    Value: 1,
    Left: &Node{
        Value: 2,
        Left: &Node{
            Value: 4,
        },
        Right: &Node{
            Value: 5,
        },
    },
    Right: &Node{
        Value: 3,
        Left: &Node{
            Value: 6,
        },
    },
}

fmt.Println(CountNodes(tree)) // Output: 6
```

### Time Complexity

The time complexity is O(n) because we visit all the nodes in the tree.

### Space Complexity

The space complexity is O(h) where h is the height of the tree.

### Resources

- [geeksforgeeks.org](https://www.geeksforgeeks.org/write-a-c-program-to-calculate-size-of-a-tree/)
- [wikipedia.org](https://en.wikipedia.org/wiki/Tree_traversal)
- [LeetCode](https://leetcode.com/problems/count-complete-tree-nodes/)
