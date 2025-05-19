package linked_lists

// ManualSet is a statically sized hash set of integers which uses
// the ending bits as a hash.
type ManualSet struct {
	Mask   int
	Values []int
}

// NewManualSet constructs a ManualSet which can contain up to 2^size entries
// and initializes it with value.
func NewManualSet(size, value int) *ManualSet {
	result := &ManualSet{
		Mask:   1<<size - 1,
		Values: make([]int, 1<<size),
	}
	result.Values[value&result.Mask] = value

	return result
}

// Insert adds the value to the set.
// Returns true if the value is already present.
// Does not handle the case where the set is full and a new value is added.
func (m *ManualSet) Insert(value int) bool {
	idx := value & m.Mask
	for m.Values[idx] != 0 {
		if m.Values[idx] == value {
			return true
		}
		idx = (idx + 1) & m.Mask
	}
	m.Values[idx] = value

	return false
}
