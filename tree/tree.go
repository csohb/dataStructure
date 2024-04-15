package tree

import "dataStructure/tree/nodeinterface"

type TreeNode[T any] struct {
	Value    T
	Children []*TreeNode[T]
}

func (t *TreeNode[T]) Add(val T) *TreeNode[T] {
	n := &TreeNode[T]{
		Value: val,
	}

	t.Children = append(t.Children, n)
	return n
}

// 요소마다 작업을 실행 할 수 있도록 function type value
func (t *TreeNode[T]) Preorder(fn func(val T)) {
	if t == nil {
		return
	}
	fn(t.Value)

	for _, n := range t.Children {
		n.Preorder(fn)
	}
}

func (t *TreeNode[T]) Postorder(fn func(val T)) {
	if t == nil {
		return
	}

	for _, n := range t.Children {
		n.Postorder(fn)
	}

	fn(t.Value)
}

func (t *TreeNode[T]) BFS(fu func(val T)) {
	queue := make([]*TreeNode[T], 0)
	queue = append(queue, t)

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]

		fu(front.Value)

		for _, n := range front.Children {
			queue = append(queue, n)
		}
	}
}

func (t *TreeNode[T]) DFS(fn func(val T)) {
	var stack []*TreeNode[T]

	stack = append(stack, t)

	for len(stack) > 0 {
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fn(last.Value)

		stack = append(stack, last.Children...)
	}
}

func (t *TreeNode[T]) GetChildren() []nodeinterface.Node {
	var children []nodeinterface.Node
	for _, c := range t.Children {
		children = append(children, c)
	}
	return children
}

func (t *TreeNode[T]) GetValue() any {
	return t.Value
}
