package searchbst

import "go-ds/src/data-structures/trees"

func searchBST(root *trees.TreeNode, val int) *trees.TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if l := searchBST(root.Left, val); l != nil {
		return l
	}
	if r := searchBST(root.Right, val); r != nil {
		return r
	}
	return nil
}
