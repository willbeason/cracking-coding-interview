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

func ToList[T comparable](values ...T) *Node[T] {
	var head *Node[T]
	for i := len(values) - 1; i >= 0; i-- {
		head = &Node[T]{Value: values[i], Next: head}
	}
	return head
}

func FromList[T comparable](head *Node[T]) []T {
	var values []T
	for ; head != nil; head = head.Next {
		values = append(values, head.Value)
	}
	return values
}
