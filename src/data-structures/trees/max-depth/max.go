package maxdepth

import (
	"go-ds/src/data-structures/trees"
)

// MaxDepth returns the maximum depth of a binary tree
func maxDepth(root *trees.TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)

	return max(l, r) + 1
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
