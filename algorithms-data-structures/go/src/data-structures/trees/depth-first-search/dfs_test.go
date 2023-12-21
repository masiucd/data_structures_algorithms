package depthfirstsearch

import (
	"go-ds/src/data-structures/trees/bst"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDfs(t *testing.T) {
	root := &bst.Node{Value: 10}
	root.Left = &bst.Node{Value: 5}
	root.Right = &bst.Node{Value: 15}
	root.Left.Left = &bst.Node{Value: 3}
	root.Left.Right = &bst.Node{Value: 7}
	root.Right.Left = &bst.Node{Value: 13}
	root.Right.Right = &bst.Node{Value: 18}
	root.Left.Left.Left = &bst.Node{Value: 1}
	root.Left.Left.Left.Left = &bst.Node{Value: 12}

	result := depthFirstSearchIter(root)
	expected := []int{10, 5, 3, 1, 12, 7, 15, 13, 18}
	assert.Equal(t, expected, result)

	result = depthFirstSearchRec(root)
	assert.Equal(t, expected, result)

}
