package linked_lists

// RemoveDuplicatesMergeSort sorts a linked list, removing duplicates as they are
// encountered. Returns the new first element of the list.
func RemoveDuplicatesMergeSort2(list *Node[int]) *Node[int] {
	return mergeSort2(list)
}

func mergeSort2(list *Node[int], length int) *Node[int] {
	// Already sorted.
	if length <= 1 {
		return list
	}

	leftLength := length / 2
	mergeSort2(list, leftLength)

	right := list.At(leftLength)
	mergeSort2(right, length-length/2)

	merge2(list, length)

	return list
}

func merge2(list *Node[int], length int) *Node[int] {
	left := list
	rightStart := list.At(length / 2)
	right := rightStart

	for left != rightStart {
		if left.Value < right.Value {
			left = left.Next
			continue
		}

		if left.Value > right.Value {
			// In-place swap
			left.Value, right.Value = right.Value, left.Value
			continue
		}
	}
}
