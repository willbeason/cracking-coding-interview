package linked_lists

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

func (n *Node[T]) Length() int {
	length := 0
	for ; n != nil; n = n.Next {
		length++
	}
	return length
}

func (n *Node[T]) At(i int) *Node[T] {
	for ; i > 0 && n != nil; i-- {
		n = n.Next
	}
	return n
}

func ToList[T comparable](values ...T) *Node[T] {
	if len(values) == 0 {
		return nil
	}

	nodes := make([]Node[T], len(values))
	for i := range nodes {
		nodes[i].Value = values[i]
		if i < len(nodes)-1 {
			nodes[i].Next = &nodes[i+1]
		}
	}
	return &nodes[0]
}

func FromList[T comparable](head *Node[T]) []T {
	values := make([]T, 0, head.Length())
	for ; head != nil; head = head.Next {
		values = append(values, head.Value)
	}
	return values
}
