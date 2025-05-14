package linked_lists

// BinaryTree stores a set of values in a binary tree.
// A node stores the value first encountered that applies to the node.
// If a node already contains a value, even values are added to the left
// and odd values to the right.
// At each layer of the tree, the rightmost bit is removed.
type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

// Insert adds value to the tree.
// Returns true if the value was already present.
func (n *BinaryTree) Insert(value int) bool {
	if n.Value == value {
		return true
	}

	if value%2 == 0 {
		if n.Left == nil {
			n.Left = &BinaryTree{Value: value / 2}
			return false
		}
		return n.Left.Insert(value / 2)
	}
	if n.Right == nil {
		n.Right = &BinaryTree{Value: value / 2}
		return false
	}
	return n.Right.Insert(value / 2)
}
