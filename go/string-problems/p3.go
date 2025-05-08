package string_problems

// Problem 3: Given two strings, write a method to decide if one is a permutation of the other.

// Solution3A uses a straightforward solution using character counts.
// Autocompleted solution.
func Solution3A(left, right string) bool {
	leftCounts := make(map[rune]int)
	for _, r := range left {
		leftCounts[r]++
	}

	rightCounts := make(map[rune]int)
	for _, r := range right {
		rightCounts[r]++
	}

	// Autocomplete begins here.
	for r, lc := range leftCounts {
		rc, ok := rightCounts[r]
		if !ok || lc != rc {
			return false
		}
	}

	return true
}

// Solution3A2 fixes the problems with the above.
func Solution3A2(left, right string) bool {
	// Strings of different lengths cannot have the same counts of characters.
	if len(left) != len(right) {
		return false
	}

	leftCounts := make(map[rune]int)
	for _, r := range left {
		leftCounts[r]++
	}

	rightCounts := make(map[rune]int)
	for _, r := range right {
		rightCounts[r]++
	}

	for r, lc := range leftCounts {
		rc, ok := rightCounts[r]
		if !ok || lc != rc {
			return false
		}
	}

	return true
}

func Solution3B(left, right string) bool {
	if len(left) != len(right) {
		return false
	}

	counts := make([]int, 128)
	for _, r := range left {
		// Keep track of character counts.
		counts[r]++
	}
	for _, r := range right {
		// If we see a character and there isn't at least one unaccounted for, we can exit immediately.
		if counts[r] < 1 {
			return false
		}
		// One instance of this character in left is now accounted for in right.
		counts[r]--
	}

	return true
}
