package linked_lists

// HashSet is a statically sized hash set of integers which uses
// the ending bits as a hash.
type HashSet struct {
	mask   int
	values []int
}

// NewHashSet constructs a HashSet which can contain up to 2^size entries
// and initializes it with value.
func NewHashSet(size, value int) *HashSet {
	result := &HashSet{
		mask:   1<<size - 1,
		values: make([]int, 1<<size),
	}
	result.values[value&result.mask] = value

	return result
}

// Insert adds the value to the set.
// Returns true if the value is already present.
// Does not handle the case where the set is full and a new value is added.
func (m *HashSet) Insert(value int) bool {
	idx := value & m.mask
	for m.values[idx] != 0 {
		if m.values[idx] == value {
			return true
		}
		idx = (idx + 1) & m.mask
	}
	m.values[idx] = value

	return false
}
