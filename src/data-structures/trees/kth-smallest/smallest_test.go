package kthsmallest

import (
	"go-ds/src/data-structures/trees"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthSmallest(t *testing.T) {
	root := &trees.TreeNode{Val: 3}
	root.Left = &trees.TreeNode{Val: 1}
	root.Right = &trees.TreeNode{Val: 4}
	root.Left.Right = &trees.TreeNode{Val: 2}

	result := kthSmallest(root, 1) // 1
	assert.Equal(t, 1, result)

	root = &trees.TreeNode{Val: 5}
	root.Left = &trees.TreeNode{Val: 3}
	root.Right = &trees.TreeNode{Val: 6}
	root.Left.Left = &trees.TreeNode{Val: 2}
	root.Left.Right = &trees.TreeNode{Val: 4}
	root.Left.Left.Left = &trees.TreeNode{Val: 1}

	result = kthSmallest(root, 3) // 3
	assert.Equal(t, 3, result)
}
