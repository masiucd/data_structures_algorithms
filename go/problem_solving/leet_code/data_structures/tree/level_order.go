package tree

import "data_structures_algos_go/common/tree"

// https://leetcode.com/problems/binary-tree-level-order-traversal/
func levelOrder(root *tree.TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := []*tree.TreeNode{root}
	for len(queue) > 0 {
		var level []int
		for range queue {
			node := queue[0]
			level = append(level, node.Val)
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, level)
	}
	return res
}
