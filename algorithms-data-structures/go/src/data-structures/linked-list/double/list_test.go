package double

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	list := NewDoubleList()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	assert.Equal(t, list.Size, 3)
	assert.Equal(t, list.Tail.Value, 3)
	assert.Equal(t, list.Head.Value, 1)
	assert.Equal(t, list.Head.Next.Value, 2)

}

func TestPrepend(t *testing.T) {
	list := NewDoubleList()
	list.Prepend(1)
	list.Prepend(2)
	list.Prepend(3)

	assert.Equal(t, list.Size, 3)
	assert.Equal(t, list.Head.Value, 3)
	assert.Equal(t, list.Tail.Value, 1)
	assert.Equal(t, list.Head.Next.Value, 2)

}

func TestGet(t *testing.T) {
	list := NewDoubleList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)

	assert.Equal(t, list.Size, 4)
	node := list.Get(0)
	assert.Equal(t, node.Value, 1)

	node = list.Get(list.Size - 1)
	assert.Equal(t, node.Value, 4)

	node = list.Get(1)
	assert.Equal(t, node.Value, 2)

	node = list.Get(2)
	assert.Equal(t, node.Value, 3)

	node = list.Get(3)
	assert.Equal(t, node.Value, 4)
}

func TestInsertAt(t *testing.T) {
	list := NewDoubleList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)

	assert.Equal(t, list.Size, 4)

	list.InsertAt(100, 0)
	assert.Equal(t, list.Size, 5)
	assert.Equal(t, list.Head.Value, 100)

	list.InsertAt(10, list.Size-1)
	assert.Equal(t, list.Size, 6)
	assert.Equal(t, list.Tail.Value, 10)

	list.InsertAt(40, 3)
	assert.Equal(t, list.Size, 7)
	assert.Equal(t, list.Get(3).Value, 40)
}
