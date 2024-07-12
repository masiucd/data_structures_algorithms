package inordertraverse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInorderTraverse(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}

	got := inorderTraversal(root)
	assert.Equal(t, got, []int{1, 3, 2})

	root = &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}

	got = inorderTraversal(root)
	assert.Equal(t, got, []int{3, 2, 4, 1, 5, 6})

}
