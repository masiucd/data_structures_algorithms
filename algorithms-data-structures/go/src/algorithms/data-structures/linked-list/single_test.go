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
