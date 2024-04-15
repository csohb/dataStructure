package binaryTree

import (
	"testing"
)

func TestDrawTree(t *testing.T) {
	root := &TreeNode{
		Value: "A",
	}

	b := &TreeNode{Value: "B"}
	root.Left = b
	c := &TreeNode{Value: "C"}
	root.Right = c

	d := &TreeNode{Value: "D"}
	b.Left = d

	e := &TreeNode{Value: "E"}
	c.Right = e
}
