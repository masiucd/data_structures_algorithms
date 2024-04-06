# Kth Smallest Element in a BST

Given a binary search tree, write a function `kthSmallest` to find the kth smallest element in it.

## Example 1:

```
Input: root = [3,1,4,null,2], k = 1
Output: 1
```

## Example 2:

```
Input: root = [5,3,6,2,4,null,null,1], k = 3
Output: 3
```

## solution

```go
func kthSmallest(root *trees.TreeNode, k int) int {
	var result []int
	dfsInOrder(root, &result)

	if k == 0 {
		return result[k]
	}
	return result[k-1]
}

func dfsInOrder(root *trees.TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	dfsInOrder(root.Left, nums)
	*nums = append(*nums, root.Val)
	dfsInOrder(root.Right, nums)
}
```

The main function here is `kthSmallest`, which takes two parameters: root, a pointer to the root node of the tree, and k, the position of the smallest element you're looking for.

The function `kthSmallest` first declares an empty slice result. It then calls the helper function dfsInOrder with the root of the tree and a pointer to the result slice.

The `dfsInOrder` function performs an _in-order traversal_ of the tree. In an in-order traversal, the left subtree is visited first, then the root node, and finally the right subtree. This order of traversal in a BST results in visiting the nodes in ascending order.

If the root node is nil, the function returns immediately, as there's nothing to traverse. If the root node is not nil, the function first recursively visits the left subtree, then appends the value of the root node to the nums slice, and finally recursively visits the right subtree.

After the `dfsInOrder` function has completed, the result slice contains the values of all nodes in the tree in ascending order. The `kthSmallest` function then returns the kth smallest element. If k is 0, it returns the first element in the slice. Otherwise, it returns the k-1th element.
