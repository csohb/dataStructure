package bst

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type LesserInt int

func (a LesserInt) Less(b Lesser) bool {
	value, ok := b.(LesserInt)
	if !ok {
		panic(fmt.Sprintf("b should be LesserInt type. type:%T", b))
	}

	return a < value
}

func TestBst(t *testing.T) {
	tree := &Tree{}
	tree.Add(LesserInt(5))
	tree.Add(LesserInt(3))
	tree.Add(LesserInt(2))
	tree.Add(LesserInt(4))
	tree.Add(LesserInt(7))
	tree.Add(LesserInt(6))
	tree.Add(LesserInt(8))
	//err := drawer.SaveTreeGraph(tree.Root, "./tree.png")
	//assert.Nil(t, err)

	assert.True(t, tree.Contains(LesserInt(7)))
	assert.True(t, tree.Contains(LesserInt(6)))
	assert.True(t, tree.Contains(LesserInt(8)))
	assert.False(t, tree.Contains(LesserInt(11)))
}
