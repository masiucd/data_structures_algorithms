package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleList_Append(t *testing.T) {
	ll := New()
	ll.Append(10)
	ll.Append(20)
	ll.Append(30)
	assert.Equal(t, 3, ll.Size)
	assert.Equal(t, ll.Tail.Value, 30)
}

func TestSingleList_Prepend(t *testing.T) {
	ll := New()
	ll.Prepend(10)
	ll.Prepend(20)
	ll.Prepend(30)
	assert.Equal(t, 3, ll.Size)
	assert.Equal(t, ll.Head.Value, 30)
}

func TestSingleList_Get(t *testing.T) {
	ll := New()
	ll.Append(10)
	ll.Append(20)
	ll.Append(30)
	assert.Equal(t, 3, ll.Size)
	assert.Equal(t, ll.Get(0).Value, 10)
	assert.Equal(t, ll.Get(1).Value, 20)
	assert.Equal(t, ll.Get(2).Value, 30)
	assert.Nil(t, ll.Get(100))
	assert.Nil(t, ll.Get(-1))
}

func TestSingleList_InsertAt(t *testing.T) {
	ll := New()
	ll.Append(10)
	ll.Append(20)
	ll.Append(30)
	ll.InsertAt(40, 1)
	assert.Equal(t, 4, ll.Size)
	assert.Equal(t, ll.Get(0).Value, 10)
	assert.Equal(t, ll.Get(1).Value, 40)
	assert.Equal(t, ll.Get(2).Value, 20)
	assert.Equal(t, ll.Get(3).Value, 30)
}

func TestSingleList_Search(t *testing.T) {
	ll := New()
	ll.Append(10)
	ll.Append(20)
	ll.Append(30)
	assert.Equal(t, true, ll.Search(10))
	assert.Equal(t, true, ll.Search(20))
	assert.Equal(t, true, ll.Search(30))
	assert.Equal(t, false, ll.Search(100))
}
