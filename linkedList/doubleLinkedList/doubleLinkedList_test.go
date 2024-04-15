package doubleLinkedList

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPushBack(t *testing.T) {
	var l LinkedList[int]

	assert.Nil(t, l.root)
	assert.Nil(t, l.tail)

	l.PushBack(1)

	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.Back().Value)
	assert.Equal(t, 1, l.Front().Value)

	l.PushBack(2)

	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 2, l.Back().Value)

	l.PushBack(3)

	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)

	assert.Equal(t, 3, l.Count())

	assert.Nil(t, l.GetAt(3))
}

func TestPushFront(t *testing.T) {
	var l LinkedList[int]

	assert.Nil(t, l.root)
	assert.Nil(t, l.tail)

	l.PushFront(1)

	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.Back().Value)
	assert.Equal(t, 1, l.Front().Value)

	l.PushFront(2)

	assert.NotNil(t, l.root)
	assert.Equal(t, 2, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	l.PushFront(3)

	assert.NotNil(t, l.root)
	assert.Equal(t, 3, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	assert.Equal(t, 3, l.Count())

	assert.Nil(t, l.GetAt(3))
}

func TestInsertAfter(t *testing.T) {
	var l LinkedList[int]

	assert.Nil(t, l.root)
	assert.Nil(t, l.tail)

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)

	node := l.GetAt(1)
	l.InsertAfter(node, 4)

	assert.Equal(t, 4, l.GetAt(2).Value)
	assert.Equal(t, 3, l.Back().Value)
	assert.Equal(t, 4, node.next.Value)
	assert.Equal(t, 3, node.next.next.Value)
	assert.Equal(t, 4, node.next.next.prev.Value)
}
func TestInsertBefore(t *testing.T) {
	var l LinkedList[int]

	assert.Nil(t, l.root)
	assert.Nil(t, l.tail)

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)

	node := l.GetAt(1)
	l.InsertBefore(node, 4)

	assert.Equal(t, 4, l.GetAt(1).Value)
	assert.Equal(t, 3, l.Back().Value)
	assert.Equal(t, 4, node.prev.Value)
	assert.Equal(t, 1, node.prev.prev.Value)
	assert.Equal(t, 4, node.prev.prev.next.Value)
}

func TestPopFront(t *testing.T) {
	var l LinkedList[int]

	assert.Nil(t, l.root)
	assert.Nil(t, l.tail)

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	n := l.PopFront()
	assert.Equal(t, 1, n.Value)
	assert.Equal(t, 2, l.root.Value)

}

func TestReverse(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.Reverse()

	assert.Equal(t, 4, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)
	assert.Equal(t, 3, l.GetAt(1).Value)
}
