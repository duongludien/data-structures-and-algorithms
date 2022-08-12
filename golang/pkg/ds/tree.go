package ds

type TreeNode struct {
	Value    int
	Parent   *TreeNode
	Children []*TreeNode
}

func NewTreeNode(value int, parent *TreeNode) *TreeNode {
	node := TreeNode{
		Value:    value,
		Parent:   parent,
		Children: nil,
	}

	if parent != nil {
		parent.Children = append(parent.Children, &node)
	}

	return &node
}

func (node *TreeNode) IsLeaf() bool {
	return node.Children == nil
}

func (node *TreeNode) DepthFirstSearch() {
	PrintInt(node.Value)

	if !node.IsLeaf() {
		for _, child := range node.Children {
			child.DepthFirstSearch()
		}
	}
}
