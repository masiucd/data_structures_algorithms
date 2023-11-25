package sumofthetree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeSum(t *testing.T) {
	var root *treeNode = &treeNode{value: 1,
		left:  &treeNode{value: 2},
		right: &treeNode{value: 3},
	}

	got := treeSum(root)
	want := 6
	assert.Equal(t, got, want)

	root = &treeNode{value: 100, left: &treeNode{value: 200}, right: &treeNode{value: 300,
		left:  &treeNode{value: 400, left: &treeNode{value: 500}},
		right: &treeNode{value: 600, left: &treeNode{value: 700, right: &treeNode{value: 800}}},
	}}
	got = treeSum(root)
	want = 3600
	assert.Equal(t, got, want)

}
