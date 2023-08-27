package tree

import "data_structures_algos_go/algorithms/tree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//https://leetcode.com/problems/binary-tree-inorder-traversal/

func inorderTraversal(root *tree.TreeNode) []int {
	var ans []int
	var inorder func(*tree.TreeNode)

	inorder = func(root *tree.TreeNode) {
		if root == nil {
			return
		}

		inorder(root.Left)
		ans = append(ans, root.Val)
		inorder(root.Right)

	}
	inorder(root)
	return ans
}
