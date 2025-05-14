package linked_lists

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"math/rand"
	"testing"
	"time"
)

type NameSolution1 struct {
	Name string
	F    func(head *Node[int])
}

var solutions = []NameSolution1{
	{Name: "Map", F: RemoveDuplicatesMap},
	{Name: "MapUninitialized", F: RemoveDuplicatesMapUninitialized},
	{Name: "Set", F: RemoveDuplicatesSet},
	{Name: "NoBuffer", F: RemoveDuplicatesNoBuffer},
	{Name: "ArrayUnsorted", F: RemoveDuplicatesArrayUnsorted},
	{Name: "ArraySorted", F: RemoveDuplicatesArraySorted},
	{Name: "BinaryTree", F: RemoveDuplicatesBinaryTree},
	{Name: "SearchTree", F: RemoveDuplicatesSearchTree},
	{Name: "QuaternaryTree", F: RemoveDuplicatesQuaternaryTree},
}

func TestProblem1(t *testing.T) {
	tt := []struct {
		name   string
		before []int
		want   []int
	}{
		{
			name:   "empty",
			before: nil,
			want:   nil,
		},
		{
			name:   "single",
			before: []int{1},
			want:   []int{1},
		},
		{
			name:   "repeated 2",
			before: []int{1, 1},
			want:   []int{1},
		},
		{
			name:   "repeated 3",
			before: []int{1, 1, 1},
			want:   []int{1},
		},
		{
			name:   "two distinct",
			before: []int{1, 2},
			want:   []int{1, 2},
		},
		{
			name:   "three distinct",
			before: []int{1, 2, 3},
			want:   []int{1, 2, 3},
		},
		{
			name:   "three distinct decreasing",
			before: []int{3, 2, 1},
			want:   []int{3, 2, 1},
		},
		{
			name:   "two distinct repeated",
			before: []int{1, 2, 1, 2},
			want:   []int{1, 2},
		},
		{
			name:   "two distinct repeated 2",
			before: []int{2, 1, 2, 1},
			want:   []int{2, 1},
		},
	}

	for _, solution := range solutions {
		t.Run(solution.Name, func(t *testing.T) {
			for _, tc := range tt {
				t.Run(tc.name, func(t *testing.T) {
					before := ToList(tc.before...)
					solution.F(before)

					after := FromList(before)
					if diff := cmp.Diff(tc.want, after); diff != "" {
						t.Errorf("%s: %s", solution.Name, diff)
					}
				})
			}
		})
	}
}

func BenchmarkProblem1(b *testing.B) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	sizes := []int{100, 200, 500, 1000, 2000, 5000, 10000}

	randomLists := make([][]int, len(sizes))
	for j, size := range sizes {
		randomLists[j] = make([]int, size)
		for i := 0; i < size; i++ {
			randomLists[j][i] = rng.Int()
		}
	}

	for _, solution := range solutions {
		for j, size := range sizes {
			b.Run(fmt.Sprintf("%s@%d", solution.Name, size), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.StopTimer()
					before := ToList(randomLists[j]...)
					b.StartTimer()
					solution.F(before)
				}
			})
		}
	}
}
