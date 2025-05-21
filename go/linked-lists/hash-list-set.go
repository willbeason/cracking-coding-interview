package linked_lists

type HashListSet struct {
	mask    int
	lists   []*Node[int]
	nodeIdx int
	nodes   []Node[int]
}

func NewHashListSet(size, value int) *HashListSet {
	keySize := size - 1
	if keySize < 0 {
		keySize = 0
	}
	result := &HashListSet{
		mask:    1<<keySize - 1,
		lists:   make([]*Node[int], 1<<keySize),
		nodeIdx: 0,
		nodes:   make([]Node[int], 1<<size),
	}

	result.Insert(value)

	return result
}

func (m *HashListSet) Insert(value int) bool {
	idx := value & m.mask
	return m.insert(value, idx)
}

func (m *HashListSet) insert(value, idx int) bool {
	list := m.lists[idx]
	// First element in list.
	if list == nil {
		list = &m.nodes[m.nodeIdx]
		m.lists[idx] = list
		list.Value = value
		m.nodeIdx++
		return false
	} else if list.Value == value {
		return true
	} else if list.Next == nil {
		list.Next = &m.nodes[m.nodeIdx]
		m.nodeIdx++
		list.Next.Value = value
		return false
	}

	list = list.Next

	for list.Next != nil {
		if list.Value == value {
			return true
		}
		list = list.Next
	}

	if list.Value == value {
		return true
	}

	list.Next = &m.nodes[m.nodeIdx]
	m.nodeIdx++
	list.Next.Value = value
	return false
}

func RemoveDuplicatesHashListSet(head *Node[int]) {
	if head == nil {
		return
	}

	length := head.Length()
	mSize := 0
	for length > 0 {
		mSize++
		length /= 2
	}

	m := NewHashListSet(mSize, head.Value)
	for ; head != nil; head = head.Next {
		for head.Next != nil && m.Insert(head.Next.Value) {
			head.Next = head.Next.Next
		}
	}
}
