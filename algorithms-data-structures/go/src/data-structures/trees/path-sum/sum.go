package pathsum

import "go-ds/src/data-structures/trees"

func hasPathSum(root *trees.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	// If we are at a leaf node and the sum is equal to the target sum, return true
	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		return true
	}
	// Recursively call the function on the left and right subtrees
	if hasPathSum(root.Left, targetSum-root.Val) {
		return true
	}
	// If the left subtree does not have the target sum, check the right subtree
	if hasPathSum(root.Right, targetSum-root.Val) {
		return true
	}
	return false
}
