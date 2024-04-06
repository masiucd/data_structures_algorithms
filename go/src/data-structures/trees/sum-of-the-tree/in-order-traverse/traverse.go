package inordertraverse

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	recursiveHelper(&result, root)
	return result
}

func recursiveHelper(result *[]int, root *TreeNode) {
	if root == nil {
		return
	}
	recursiveHelper(result, root.Left)
	*result = append(*result, root.Val)
	recursiveHelper(result, root.Right)
}
