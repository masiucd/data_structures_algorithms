package breathfirstsearch

import (
	"go-ds/src/data-structures/trees/bst"
	"slices"
)

func breathFirstSearch(root *bst.Node) []int {
	var result []int
	var queue []*bst.Node = []*bst.Node{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = slices.Delete(queue, 0, 1)
		result = append(result, node.Value)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return result
}
