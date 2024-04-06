package pathsum

import (
	"go-ds/src/data-structures/trees"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathSum(t *testing.T) {
	root := &trees.TreeNode{Val: 1}
	root.Left = &trees.TreeNode{Val: 2}
	root.Right = &trees.TreeNode{Val: 3}

	assert.Equal(t, hasPathSum(root, 4), true)
	assert.Equal(t, hasPathSum(root, 5), false)

	root = &trees.TreeNode{Val: 1}
	root.Left = &trees.TreeNode{Val: 2}
	root.Right = &trees.TreeNode{Val: 3}
	root.Left.Left = &trees.TreeNode{Val: 4}
	root.Left.Right = &trees.TreeNode{Val: 5}
	root.Right.Left = &trees.TreeNode{Val: 6}
	root.Right.Right = &trees.TreeNode{Val: 7}

	assert.Equal(t, hasPathSum(root, 10), true)
	assert.Equal(t, hasPathSum(root, 18), false)
}
