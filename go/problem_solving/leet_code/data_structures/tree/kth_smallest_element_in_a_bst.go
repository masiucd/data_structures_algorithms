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
func kthSmallest(root *tree.TreeNode, k int) int {
	var stack []*tree.TreeNode

	stack = append(stack, root)
	var curr = root
	for len(stack) != 0 {
		// Go to the left leaf
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		// Process the node (ie: decrement the k) and check if condition is met
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k -= 1
		if k == 0 {
			return curr.Val
		}

		// Go to right
		curr = curr.Right
	}

	return -1
}
