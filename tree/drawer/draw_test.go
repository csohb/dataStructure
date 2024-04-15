package drawer

import (
	"dataStructure/tree"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDrawTree(t *testing.T) {
	root := &tree.TreeNode[string]{
		Value: "A",
	}

	b := root.Add("B")
	root.Add("C")
	d := root.Add("D")

	b.Add("E")
	b.Add("F")

	d.Add("G")

	err := SaveTreeGraph[string](root, "./tree.png")
	assert.Nil(t, err, true)
}
