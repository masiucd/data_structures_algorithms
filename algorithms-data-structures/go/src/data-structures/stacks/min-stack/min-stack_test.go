package minstack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinStack(t *testing.T) {
	m := Constructor()
	m.Push(-2)
	m.Push(0)
	m.Push(-3)
	min := m.GetMin()
	assert.Equal(t, -3, min)
	m.Pop()
	top := m.Top()
	assert.Equal(t, 0, top)
	min = m.GetMin()
	assert.Equal(t, -2, min)
}
