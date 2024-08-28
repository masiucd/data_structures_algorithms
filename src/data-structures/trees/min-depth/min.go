package mindepth

import "go-ds/src/data-structures/trees"

type Node struct {
	value *trees.TreeNode
	level int
}

func minDepth(root *trees.TreeNode) int {
	if root == nil {
		return 0
	}
	var q []*Node
	var min int
	q = append(q, &Node{value: root, level: 1})
	for len(q) > 0 {
		popped := q[0]
		q = q[1:]
		if popped.value.Left == nil && popped.value.Right == nil {
			min = popped.level
			break
		}

		if popped.value.Left != nil {
			q = append(q, &Node{value: popped.value.Left, level: popped.level + 1})
		}
		if popped.value.Right != nil {
			q = append(q, &Node{value: popped.value.Right, level: popped.level + 1})
		}
	}

	return min
}
