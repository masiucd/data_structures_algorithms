package sumofthetree

import "go-ds/src/data-structures/trees"

func treeSum(root *trees.TreeNode) int {
	if root == nil {
		return 0
	}
	left := treeSum(root.Left)
	right := treeSum(root.Right)
	return root.Val + left + right
}
