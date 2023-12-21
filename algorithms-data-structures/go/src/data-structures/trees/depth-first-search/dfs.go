package depthfirstsearch

import "go-ds/src/data-structures/trees/bst"

func depthFirstSearchIter(root *bst.Node) []int {
	var result []int
	var stack []*bst.Node = []*bst.Node{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Value)
		// Starting from the right since we are using a stack
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return result
}

func depthFirstSearchRec(root *bst.Node) []int {
	var result []int
	var dfs func(node *bst.Node)
	dfs = func(node *bst.Node) {
		if node == nil {
			return
		}
		result = append(result, node.Value)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return result
}
