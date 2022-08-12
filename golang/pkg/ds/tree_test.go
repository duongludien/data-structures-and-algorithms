package ds

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewTreeNode_NewRootNode_HasCorrectValues(t *testing.T) {
	nodeValue := 1
	node := NewTreeNode(nodeValue, nil)

	assert.Equal(t, nodeValue, node.Value)
	assert.Nil(t, node.Parent)
	assert.Nil(t, node.Children)
}

func Test_NewTreeNode_WithChildren_HasCorrectStructure(t *testing.T) {
	rootValue := 0

	root := NewTreeNode(rootValue, nil)
	NewTreeNode(1, root)
	NewTreeNode(2, root)
	NewTreeNode(3, root)

	assert.Equal(t, rootValue, root.Value)
	assert.Nil(t, root.Parent)
	assert.Equal(t, 3, len(root.Children))

	for index, child := range root.Children {
		assert.Equal(t, index+1, child.Value)
		assert.True(t, child.IsLeaf())
		assert.Equal(t, root, child.Parent)
		assert.Nil(t, child.Children)
	}
}

func TestDfs(t *testing.T) {
	rootValue := 0

	root := NewTreeNode(rootValue, nil)
	child1 := NewTreeNode(1, root)
	child2 := NewTreeNode(2, root)
	child3 := NewTreeNode(3, root)
	NewTreeNode(4, child1)
	NewTreeNode(5, child1)
	NewTreeNode(6, child2)
	NewTreeNode(7, child2)
	NewTreeNode(8, child2)
	NewTreeNode(9, child3)

	root.DepthFirstSearch()
}
