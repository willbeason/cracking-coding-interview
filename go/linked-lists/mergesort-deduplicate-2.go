package linked_lists

// RemoveDuplicatesMergeSort2 sorts a linked list, removing duplicates as they are
// encountered. Returns the new first element of the list.
func RemoveDuplicatesMergeSort2(list *Node[int]) *Node[int] {
	return mergeSort2(list, list.Length())
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
	beforeRight := list.At(length/2 - 1)
	right := list.At(length / 2)

	for left != right && right != nil {
		if left.Value > right.Value {
			// In-place swap.
			left.Value, right.Value = right.Value, left.Value

			// Check if the right list is still sorted. We only need to check the right list.
			// As the element of the left list has only decreased, it cannot have a sorting violation.
			if right.Next != nil && right.Value > right.Next.Value {
				// The right list is no longer sorted, so insert the node after the current node in the left list.

				// Remove right.
				beforeRight.Next = right.Next

				// Insert right.
				right.Next = left.Next
				left.Next = right

				// Advance right.
				if left == beforeRight {
					// Since we inserted at the boundary, shift right by one.
					beforeRight = beforeRight.Next
				}
				right = beforeRight.Next
			}
			continue
		} else if left.Value == right.Value {
			// Remove right since it is equivalent to left.
			beforeRight.Next = right.Next
			right = beforeRight.Next
			continue
		}

		// Advance left.
		for left.Next != nil && left != beforeRight && left.Value == left.Next.Value {
			// Remove identical values when advancing left.
			left.Next = left.Next.Next
		}
		left = left.Next
		continue
	}

	return list
}
