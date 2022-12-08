package can_reach_leaf

import "data_structures_algos_go/common/tree"

//Determine if a path exists from the root of the tree to a lead node. It may NOT contain any zeroes.

func canReachLeaf(root *tree.TreeNode) bool {
	if root == nil || root.Val == 0 {
		return false
	}
	//Is it a leafNode?
	if root.Left == nil && root.Right == nil {
		return true
	}
	if canReachLeaf(root.Left) {
		return true
	}
	if canReachLeaf(root.Right) {
		return true
	}
	return false
}

func canReachLeafWithPath(root *tree.TreeNode, path *[]int) bool {
	if root == nil || root.Val == 0 {
		return false
	}
	*path = append(*path, root.Val)
	if root.Left == nil && root.Right == nil {
		return true
	}
	if canReachLeafWithPath(root.Left, path) {
		return true
	}
	if canReachLeafWithPath(root.Right, path) {
		return true
	}
	*path = (*path)[:len(*path)-1]
	return false
}
