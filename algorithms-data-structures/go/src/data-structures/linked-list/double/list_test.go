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
