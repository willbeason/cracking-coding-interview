package linked_lists

func RemoveDuplicatesHeapSort(list *Node[int]) *Node[int] {
	heapSort(list)

	head := list
	for ; head != nil; head = head.Next {
		for head.Next != nil && head.Value == head.Next.Value {
			head.Next = head.Next.Next
		}
	}

	return list
}

func heapSort(list *Node[int]) {
	if list == nil || list.Next == nil {
		// Lists of length <= 1 are sorted.
		return
	}

	end := list.Length()
	start := end / 2

	for end > 1 {
		if start > 0 {
			start--
		} else {
			end--

			// swap a[0], a[end]
			left := list.At(0)
			right := list.At(end)
			left.Value, right.Value = right.Value, left.Value
		}

		root := start
		for child := iLeftChild(root); child < end; child = iLeftChild(root) {
			biggestChild := list.At(child)
			if biggestChild.Next != nil && biggestChild.Value < biggestChild.Next.Value {
				biggestChild = biggestChild.Next
			}

			parent := list.At(root)
			if parent.Value < biggestChild.Value {
				parent.Value, biggestChild.Value = biggestChild.Value, parent.Value
				root = child
			} else {
				break
			}
		}
	}
}

func iLeftChild(i int) int {
	return 2*i + 1
}
