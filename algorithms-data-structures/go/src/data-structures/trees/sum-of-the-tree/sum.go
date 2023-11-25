package sumofthetree

type treeNode struct {
	value int
	left  *treeNode
	right *treeNode
}

func treeSum(root *treeNode) int {
	if root == nil {
		return 0
	}
	left := treeSum(root.left)
	right := treeSum(root.right)
	return root.value + left + right
}
