package insert

import "go-ds/src/data-structures/trees"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *trees.TreeNode, val int) *trees.TreeNode {
	return insert(root, val)
}

func insert(root *trees.TreeNode, val int) *trees.TreeNode {
	if root == nil {
		return &trees.TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = insert(root.Left, val)
	} else if val > root.Val {
		root.Right = insert(root.Right, val)
	}
	return root
}
