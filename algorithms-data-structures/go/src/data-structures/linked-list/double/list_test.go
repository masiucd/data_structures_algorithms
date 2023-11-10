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
