package bst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func buildTree() Tree {
	bst := NewBst()
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)
	bst.Insert(20)
	bst.Insert(2)
	bst.Insert(1)
	return bst

}

func TestFind(t *testing.T) {
	bst := buildTree()
	node := bst.Find(15)
	assert.Equal(t, 15, node.Value)

	node = bst.Find(100)
	assert.Nil(t, node)
}

func TestContains(t *testing.T) {
	bst := buildTree()
	assert.True(t, bst.Contains(15))
	assert.False(t, bst.Contains(100))
}

func TestRemove(t *testing.T) {
	bst := buildTree()
	assert.True(t, bst.Contains(15))
	bst.Remove(15)
	assert.False(t, bst.Contains(15))
}

func TestMin(t *testing.T) {
	bst := buildTree()
	assert.Equal(t, 1, bst.Min().Value)
}

func TestMax(t *testing.T) {
	bst := buildTree()
	assert.Equal(t, 20, bst.Max().Value)
}

func TestBFS(t *testing.T) {
	bst := buildTree()
	assert.Equal(t, []int{10, 5, 15, 2, 20, 1}, bst.BFS())
}

func TestDFSInOrder(t *testing.T) {
	bst := buildTree()
	assert.Equal(t, []int{1, 2, 5, 10, 15, 20}, bst.DFS("IN"))
}

func TestDFSPreOrder(t *testing.T) {
	bst := buildTree()
	assert.Equal(t, []int{10, 5, 2, 1, 15, 20}, bst.DFS("PRE"))
}

func TestDFSPostOrder(t *testing.T) {
	bst := buildTree()
	assert.Equal(t, []int{1, 2, 5, 20, 15, 10}, bst.DFS("POST"))
}
