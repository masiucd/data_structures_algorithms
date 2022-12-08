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
func postorderTraversal(root *tree.TreeNode) []int {
	var result []int
	dfs(root, &result)
	return result
}

func dfs(node *tree.TreeNode, result *[]int) {
	if node == nil {
		return
	}
	dfs(node.Left, result)
	dfs(node.Right, result)
	*result = append(*result, node.Val)

}
