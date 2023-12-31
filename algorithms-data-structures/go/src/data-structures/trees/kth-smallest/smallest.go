package kthsmallest

import "go-ds/src/data-structures/trees"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
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
