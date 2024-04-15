package singleLinkedList

type Node[T any] struct {
	next  *Node[T]
	Value T
}

type LinkedList[T any] struct {
	root *Node[T]
	tail *Node[T]

	count int
}

// PushBack push new node to back of the list
func (l *LinkedList[T]) PushBack(value T) {
	node := &Node[T]{
		Value: value,
	}
	l.count++
	if l.root == nil {
		l.root = node
		l.tail = node
		return
	}

	l.tail.next = node
	l.tail = node
}

// PushFront push new node to the front of the list
func (l *LinkedList[T]) PushFront(value T) {
	node := &Node[T]{
		Value: value,
	}
	l.count++
	if l.root == nil {
		l.root = node
		l.tail = node
		return
	}

	node.next = l.root
	l.root = node
}

func (l *LinkedList[T]) Front() *Node[T] {
	return l.root
}
func (l *LinkedList[T]) Back() *Node[T] {
	return l.tail
}

func (l *LinkedList[T]) Count() int {
	node := l.root
	cnt := 0

	for ; node != nil; node = node.next {
		cnt++
	}
	return cnt
}

func (l *LinkedList[T]) Count2() int {
	return l.count
}

func (l *LinkedList[T]) GetAt(index int) *Node[T] {
	if index >= l.Count2() {
		return nil
	}
	i := 0
	for node := l.root; node != nil; node = node.next {
		if i == index {
			return node
		}
		i++
	}
	return nil
}

func (l *LinkedList[T]) IsInclude(node *Node[T]) bool {
	nowNode := l.root
	for ; nowNode != nil; nowNode = nowNode.next {
		if nowNode == node {
			return true
		}
	}
	return false
}

func (l *LinkedList[T]) InsertAfter(node *Node[T], value T) {
	if !l.IsInclude(node) {
		return
	}

	newNode := &Node[T]{Value: value}

	originNext := node.next

	node.next = newNode
	newNode.next = originNext
	l.count++
}

func (l *LinkedList[T]) FindPreviousNode(node *Node[T]) *Node[T] {
	// find previous node
	rootNode := l.root
	for ; rootNode != nil; rootNode = rootNode.next {
		if rootNode.next == node {
			return rootNode
		}
	}
	return nil
}

func (l *LinkedList[T]) InsertBefore(node *Node[T], value T) {
	if node == l.root {
		l.PushFront(value)
		return
	}

	prevNode := l.FindPreviousNode(node)
	if prevNode == nil {
		return
	}

	newNode := &Node[T]{Value: value}

	prevNode.next = newNode
	newNode.next = node
	l.count++
}

func (l *LinkedList[T]) PopFront() *Node[T] {
	if l.root == nil {
		return nil
	}
	n := l.root
	l.root.next, l.root = nil, l.root.next
	if l.root == nil {
		l.tail = nil
	}
	l.count--
	return n
}

func (l *LinkedList[T]) Remove(node *Node[T]) {
	if node == l.root {
		l.PopFront()
		return
	}

	prev := l.FindPreviousNode(node)
	if prev == nil {
		return
	}
	prev.next = node.next
	node.next = nil
	if node == l.tail {
		l.tail = prev
	}
	l.count--
}
func (l *LinkedList[T]) Reverse() {
	newL := &LinkedList[T]{}
	for l.root != nil {
		n := l.PopFront()
		newL.PushFront(n.Value)
	}
	l.count = newL.count
	l.root = newL.root
	l.tail = newL.tail
}
func (l *LinkedList[T]) Reverse2() {
	if l.root == nil {
		return
	}
	node := l.root
	next := node.next
	l.root.next = nil

	for next != nil {
		nextNext := next.next
		next.next = node

		node = next
		next = nextNext
	}

	l.root, l.tail = l.tail, l.root
}
