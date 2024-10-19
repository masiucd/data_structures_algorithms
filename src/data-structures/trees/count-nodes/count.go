package countnodes

import "go-ds/src/data-structures/trees"

// CountNodes returns the total number of nodes in a binary tree.
// It takes the root node of the tree as an argument and recursively
// counts the nodes in the left and right subtrees, adding 1 for the root node itself.
//
// Parameters:
// - root: A pointer to the root node of the binary tree.
//
// Returns:
// - int: The total number of nodes in the binary tree.
func CountNodes(root *trees.TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + CountNodes(root.Left) + CountNodes(root.Right)
}
