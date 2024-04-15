package doubleLinkedList

type Node[T any] struct {
	next *Node[T]
	prev *Node[T]

	Value T
}

type LinkedList[T any] struct {
	root  *Node[T]
	tail  *Node[T]
	count int
}

func (l *LinkedList[T]) PushBack(val T) {
	n := &Node[T]{
		Value: val,
	}

	if l.tail == nil {
		l.root = n
		l.tail = n
		l.count = 1
		return
	}

	l.tail.next = n
	n.prev = l.tail
	l.tail = n
	l.count++
}

func (l *LinkedList[T]) Front() *Node[T] {
	return l.root
}
func (l *LinkedList[T]) Back() *Node[T] {
	return l.tail
}

func (l *LinkedList[T]) Count() int {
	return l.count
}
func (l *LinkedList[T]) GetAt(index int) *Node[T] {
	if index >= l.Count() {
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

func (l *LinkedList[T]) PushFront(val T) {
	n := &Node[T]{Value: val}
	if l.root == nil {
		l.root = n
		l.tail = n
		l.count = 1
		return
	}

	l.root.prev = n
	n.next = l.root
	l.root = n
	l.count++
}

func (l *LinkedList[T]) IsInclude(node *Node[T]) bool {
	newNode := l.root
	for ; newNode != nil; newNode = newNode.next {
		if newNode == node {
			return true
		}
	}
	return false
}

func (l *LinkedList[T]) InsertAfter(node *Node[T], val T) {
	if !l.IsInclude(node) {
		return
	}

	n := &Node[T]{Value: val}

	if node == l.tail {
		l.PushBack(val)
		return
	}

	originNext := node.next
	node.next = n

	n.next = originNext
	n.prev = node

	originNext.prev = n

	l.count++
}

func (l *LinkedList[T]) InsertBefore(node *Node[T], val T) {
	if !l.IsInclude(node) {
		return
	}

	n := &Node[T]{Value: val}

	if node == l.root {
		l.PushFront(val)
	}

	originPrev := node.prev
	node.prev = n

	n.prev = originPrev
	n.next = node

	originPrev.next = n
	l.count++
}

func (l *LinkedList[T]) PopFront() *Node[T] {
	if l.root == nil {
		return nil
	}

	n := l.root

	l.root = n.next

	if l.root != nil {
		l.root.prev = nil
	} else {
		l.tail = nil
	}
	n.next = nil

	l.count--
	return n

}

func (l *LinkedList[T]) PopBack() *Node[T] {
	if l.tail == nil {
		return nil
	}
	n := l.tail
	l.tail = n.prev
	if l.tail != nil {
		l.tail.next = nil
	} else {
		l.root = nil
	}
	n.prev = nil
	l.count--
	return n
}

func (l *LinkedList[T]) Reverse() {
	if l.root == nil {
		return
	}
	node := l.root
	var next *Node[T]

	for node != nil {
		next = node.next
		node.prev, node.next = node.next, node.prev
		node = next
	}

	l.root, l.tail = l.tail, l.root
}
