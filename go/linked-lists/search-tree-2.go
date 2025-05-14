package linked_lists

const size = 16

type SearchTree struct {
	Values      [size]int
	Index       int
	Left, Right *SearchTree
}

func (n *SearchTree) Insert(value int) bool {
	for i := 0; i < n.Index; i++ {
		if n.Values[i] == value {
			return true
		}
	}
	if n.Index < size {
		n.Values[n.Index] = value
		n.Index++
		return false
	}

	if value%2 == 0 {
		if n.Left == nil {
			n.Left = &SearchTree{
				Values: [size]int{value / 2},
				Index:  1,
			}
			return false
		}
		return n.Left.Insert(value / 2)
	}
	if n.Right == nil {
		n.Right = &SearchTree{
			Values: [size]int{value / 2},
			Index:  1,
		}
		return false
	}
	return n.Right.Insert(value / 2)
}
