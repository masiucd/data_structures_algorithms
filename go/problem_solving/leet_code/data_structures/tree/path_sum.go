package tree

import "data_structures_algos_go/common/tree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Use inorder DFS
func hasPathSum(root *tree.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	//Check for leaf node
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}
