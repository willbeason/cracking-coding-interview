package linked_lists

const Values = 1 << 6
const Children = 1 << 2

type QuaternaryTree struct {
	Values   [Values]int
	Children [Children]*QuaternaryTree
}

func NewQuaternaryTree(value int) *QuaternaryTree {
	result := &QuaternaryTree{}
	result.Values[value%Values] = value
	return result
}

func (n *QuaternaryTree) Insert(value int) bool {
	for {
		valueIndex := value % Values
		if n.Values[valueIndex] == 0 {
			n.Values[valueIndex] = value
			return false
		}
		if n.Values[valueIndex] == value {
			return true
		}

		childIndex := value % Children
		value /= Children
		if n.Children[childIndex] == nil {
			n.Children[childIndex] = NewQuaternaryTree(value)
			return false
		}
		n = n.Children[childIndex]
	}
}
