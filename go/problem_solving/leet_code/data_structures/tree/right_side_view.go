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

func rightSideView(root *tree.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var result []int

	var fn func(n *tree.TreeNode, result *[]int, depth int)
	fn = func(n *tree.TreeNode, result *[]int, depth int) {
		if n == nil {
			return
		}
		if depth == len(*result) {
			*result = append(*result, n.Val)
		}
		fn(n.Right, result, depth+1)
		fn(n.Left, result, depth+1)
	}

	fn(root, &result, 0)

	return result
}
