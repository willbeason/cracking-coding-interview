package linked_lists

// RemoveDuplicatesMergeSort sorts a linked list, removing duplicates as they are
// encountered. Returns the new first element of the list.
func RemoveDuplicatesMergeSort(list *Node[int]) *Node[int] {
	return mergeSort(list)
}

func mergeSort(list *Node[int]) *Node[int] {
	switch {
	case list == nil || list.Next == nil:
		// Already sorted.
		return list
	case list.Next.Next == nil:
		// Trivial to sort in place.
		if list.Value > list.Next.Value {
			list.Value, list.Next.Value = list.Next.Value, list.Value
		} else if list.Value == list.Next.Value {
			list.Next = nil
		}
		return list
	case list.Next.Next.Next == nil:
		return sort3(list)
	}

	left, right := partitionList(list)
	left = mergeSort(left)
	right = mergeSort(right)

	return merge(left, right)
}

func sort3(list *Node[int]) *Node[int] {
	a := list.Value
	b := list.Next.Value
	c := list.Next.Next.Value
	if a < b {
		if b < c {
			// Nothing
		} else if b == c || a == c {
			list.Next.Next = nil
		} else {
			// b > c
			if a < c {
				// b, c = c, b
				list.Next.Value, list.Next.Next.Value = list.Next.Next.Value, list.Next.Value
			} else {
				// a > c
				// a, b, c = c, a, b
				list.Value, list.Next.Value, list.Next.Next.Value = list.Next.Next.Value, list.Value, list.Next.Value
			}
		}
	} else if a > b {
		if b > c {
			list.Value, list.Next.Next.Value = list.Next.Next.Value, list.Value
		} else if b == c {
			list.Next.Next = nil
			// a, b = b, c
			list.Value, list.Next.Value = list.Next.Value, list.Value
		} else {
			// b < c
			if a < c {
				list.Value, list.Next.Value = list.Next.Value, list.Value
			} else if a > c {
				list.Value, list.Next.Value, list.Next.Next.Value = list.Next.Value, list.Next.Next.Value, list.Value
			} else {
				list.Next.Next = nil
				list.Value, list.Next.Value = list.Next.Value, list.Value
			}
		}
	} else if a == b {
		list.Next = list.Next.Next
		if a > c {
			list.Value = list.Next.Value
		} else if a == c {
			list.Next = nil
		}
	}
	return list
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
