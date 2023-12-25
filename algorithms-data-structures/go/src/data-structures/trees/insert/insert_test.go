package insert

import (
	"go-ds/src/data-structures/trees"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBstInsert(t *testing.T) {
	root := &trees.TreeNode{Val: 10}
	root.Left = &trees.TreeNode{Val: 2}
	root.Right = &trees.TreeNode{Val: 12}

	res := insertIntoBST(root, 5)
	assert.Equal(t, 5, res.Left.Right.Val)

	res = insertIntoBST(root, 100)
	assert.Equal(t, 100, res.Right.Right.Val)

}
