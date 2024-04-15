package binaryTree

import "dataStructure/tree/nodeinterface"

type Tree struct {
	Root *TreeNode
}

type TreeNode struct {
	Value any

	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) GetChildren() []nodeinterface.Node {
	var children []nodeinterface.Node
	if t.Left == nil {
		children = append(children, nil)
	} else {
		children = append(children, t.Left)
	}
	if t.Right == nil {
		children = append(children, nil)
	} else {
		children = append(children, t.Right)
	}
	return children
}

func (t *TreeNode) GetValue() any {
	return t.Value
}
