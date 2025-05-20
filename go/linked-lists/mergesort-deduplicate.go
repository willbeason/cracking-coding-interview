package linked_lists

// RemoveDuplicatesMergeSort sorts a linked list, removing duplicates as they are
// encountered. Returns the new first element of the list.
func RemoveDuplicatesMergeSort(list *Node[int]) *Node[int] {
	return mergeSort(list)
}

func mergeSort(list *Node[int]) *Node[int] {
	// Already sorted.
	if list == nil || list.Next == nil {
		return list
	}

	left, right := partitionList(list)
	left = mergeSort(left)
	right = mergeSort(right)

	return merge(left, right)
}

func partitionList(list *Node[int]) (*Node[int], *Node[int]) {
	if list == nil || list.Next == nil {
		return list, nil
	}

	// Partition list.
	leftStart := list
	rightStart := list.Next

	left := leftStart
	right := rightStart
	head := list.Next.Next
	for head != nil {
		left.Next = head
		left = left.Next
		head = head.Next
		if head != nil {
			right.Next = head
			right = right.Next
			head = head.Next
		}
	}
	left.Next = nil
	if right != nil {
		right.Next = nil
	}

	return leftStart, rightStart
}

func merge(left, right *Node[int]) *Node[int] {
	if left == nil {
		if right == nil {
			return nil
		}
		return right
	}
	if right == nil {
		return left
	}

	var head *Node[int]
	var tail *Node[int]
	if left.Value < right.Value {
		head = left
		left = left.Next
	} else {
		head = right
		right = right.Next
	}
	tail = head

	for left != nil && right != nil {
		var next *Node[int]
		if left.Value < right.Value {
			next = left
			left = left.Next
		} else {
			next = right
			right = right.Next
		}

		if next.Value == tail.Value {
			// Ignore duplicates.
			continue
		}
		tail.Next = next
		tail = tail.Next
	}
	for left != nil {
		if tail.Value != left.Value {
			tail.Next = left
			tail = tail.Next
		}
		left = left.Next
	}
	for right != nil {
		if tail.Value != right.Value {
			tail.Next = right
			tail = tail.Next
		}
		right = right.Next
	}
	tail.Next = nil

	return head
}
