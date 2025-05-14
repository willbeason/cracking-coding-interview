package linked_lists

type QuaternaryTree struct {
	Value0, Value1, Value2, Value3, Value4 int
	Children                               [4]*QuaternaryTree
}

func NewQuaternaryTree(value int) *QuaternaryTree {
	return &QuaternaryTree{
		Value0: value,
		Value1: -1,
		Value2: -1,
		Value3: -1,
		Value4: -1,
	}
}

func (n *QuaternaryTree) Insert(value int) bool {
	switch {
	case n.Value0 == value:
		return true
	case n.Value1 == -1:
		n.Value1 = value
		return false
	case n.Value1 == value:
		return true
	case n.Value2 == -1:
		n.Value2 = value
		return false
	case n.Value2 == value:
		return true
	case n.Value3 == -1:
		n.Value3 = value
		return false
	case n.Value3 == value:
		return true
	case n.Value4 == -1:
		n.Value4 = value
		return false
	case n.Value4 == value:
		return true
	}

	childIndex := value % 4
	if n.Children[childIndex] == nil {
		n.Children[childIndex] = NewQuaternaryTree(value / 4)
		return false
	}
	return n.Children[childIndex].Insert(value / 4)
}
