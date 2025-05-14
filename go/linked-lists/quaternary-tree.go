package linked_lists

type QuaternaryTree struct {
	Value0, Value1, Value2                     int
	LeftLeft, LeftRight, RightLeft, RightRight *QuaternaryTree
}

func NewQuaternaryTree(value int) *QuaternaryTree {
	return &QuaternaryTree{Value0: value, Value1: -1, Value2: -1}
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
	}

	switch value % 4 {
	case 0:
		if n.LeftLeft == nil {
			n.LeftLeft = NewQuaternaryTree(value / 4)
			return false
		}
		return n.LeftLeft.Insert(value / 4)
	case 1:
		if n.LeftRight == nil {
			n.LeftRight = NewQuaternaryTree(value / 4)
			return false
		}
		return n.LeftRight.Insert(value / 4)
	case 2:
		if n.RightLeft == nil {
			n.RightLeft = NewQuaternaryTree(value / 4)
			return false
		}
		return n.RightLeft.Insert(value / 4)
	default:
		if n.RightRight == nil {
			n.RightRight = NewQuaternaryTree(value / 4)
			return false
		}
		return n.RightRight.Insert(value / 4)
	}
}
