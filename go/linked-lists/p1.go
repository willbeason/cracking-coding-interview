package linked_lists

import (
	"slices"
)

// RemoveDuplicatesMap removes entries from a linked list with the same value.
// Uses a map to record seen values.
func RemoveDuplicatesMap(head *Node[int]) {
	if head == nil {
		return
	}

	seen := make(map[int]bool, head.Length())

	for ; head != nil; head = head.Next {
		seen[head.Value] = true

		for head.Next != nil && seen[head.Next.Value] {
			head.Next = head.Next.Next
		}
	}
}

// RemoveDuplicatesArrayUnsorted removes entries from a linked list with the same value.
// Uses an unsorted array to record seen values.
func RemoveDuplicatesArrayUnsorted(head *Node[int]) {
	seen := make([]int, 0, head.Length())
	for ; head != nil; head = head.Next {
		seen = append(seen, head.Value)

		for head.Next != nil {
			isSeen := false
			for _, v := range seen {
				if v == head.Next.Value {
					isSeen = true
					break
				}
			}
			if isSeen {
				head.Next = head.Next.Next
			} else {
				break
			}
		}
	}
}

// RemoveDuplicatesArraySorted removes entries from a linked list with the same value.
// Uses a sorted array to record seen values.
func RemoveDuplicatesArraySorted(head *Node[int]) {
	if head == nil {
		return
	}

	seen := make([]int, 1, head.Length())
	seen[0] = head.Value
	for ; head != nil; head = head.Next {
		for head.Next != nil {
			idx, isSeen := slices.BinarySearch(seen, head.Next.Value)
			if !isSeen {
				seen = slices.Insert(seen, idx, head.Next.Value)
				break
			}

			head.Next = head.Next.Next
		}
	}
}
