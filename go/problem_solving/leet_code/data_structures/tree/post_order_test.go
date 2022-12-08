package tree

import (
	"data_structures_algos_go/common/tree"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostOrder(t *testing.T) {
	got := postorderTraversal(&tree.TreeNode{Val: 1, Right: &tree.TreeNode{Val: 2, Left: &tree.TreeNode{Val: 3}}})
	want := []int{3, 2, 1}

	for i := range got {
		assert.Equal(t, want[i], got[i], fmt.Sprintf("Expected %d, got %d", want[i], got[i]))
	}

}
