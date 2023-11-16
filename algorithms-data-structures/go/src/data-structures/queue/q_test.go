package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQ(t *testing.T) {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	assert.Equal(t, 3, q.Len())
	assert.Equal(t, "1,2,3", q.Print())

	node := q.Dequeue()
	assert.Equal(t, 1, node.Value)
	assert.Equal(t, 2, q.Len())

	node = q.Peek()
	assert.Equal(t, 2, node.Value)

}
