package maxdepth

import (
	"go-ds/src/data-structures/trees"

	list "github.com/bahlo/generic-list-go"
)

// MaxDepth returns the maximum depth of a binary tree
func maxDepth(root *trees.TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)

	return max(l, r) + 1
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepthV2(root *trees.TreeNode) int {
	if root == nil {
		return 0
	}
	q := list.New[*NodeDepth]()
	q.PushBack(&NodeDepth{node: root, depth: 1})
	max := 0
	for q.Len() > 0 {
		el := q.Front()
		q.Remove(el)
		node := el.Value.node
		depth := el.Value.depth
		if depth > max {
			max = depth
		}
		if node.Left != nil {
			q.PushBack(&NodeDepth{node: node.Left, depth: el.Value.depth + 1})
		}
		if node.Right != nil {
			q.PushBack(&NodeDepth{node: node.Right, depth: el.Value.depth + 1})
		}
	}

	return max
}

type NodeDepth struct {
	node  *trees.TreeNode
	depth int
}
