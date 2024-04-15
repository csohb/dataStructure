package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPush(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	assert.Equal(t, 1, s.Pop())
	assert.Equal(t, 2, s.Pop())
	assert.Equal(t, 3, s.Pop())
}

func TestPushSlice(t *testing.T) {
	s := NewSliceQueue[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	assert.Equal(t, 1, s.Pop())
	assert.Equal(t, 2, s.Pop())
	assert.Equal(t, 3, s.Pop())
}
