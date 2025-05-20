package linked_lists

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"math/rand"
	"sort"
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
	{Name: "ManualSet", F: RemoveDuplicatesManualMap},
}

var tt = []struct {
	name   string
	before []int
	want   []int
}{
	{
		name:   "empty",
		before: []int{},
		want:   []int{},
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
	{
		name:   "two distinct repeated 3",
		before: []int{1, 2, 2, 1},
		want:   []int{1, 2},
	},
	{
		name:   "five distinct",
		before: []int{1, 2, 3, 4, 5},
		want:   []int{1, 2, 3, 4, 5},
	},
	{
		name:   "three increase decrease",
		before: []int{1, 2, 3, 3, 2, 1},
		want:   []int{1, 2, 3},
	},
	{
		name:   "three increase decrease 2",
		before: []int{1, 2, 3, 2, 1},
		want:   []int{1, 2, 3},
	},
	{
		name:   "four increase decrease",
		before: []int{1, 2, 3, 4, 4, 3, 2, 1},
		want:   []int{1, 2, 3, 4},
	},
	{
		name:   "four increase decrease 2",
		before: []int{1, 2, 3, 4, 3, 2, 1},
		want:   []int{1, 2, 3, 4},
	},
	{
		name:   "four increase decrease 2",
		before: []int{1, 3, 5, 7, 2, 4, 6, 8},
		want:   []int{1, 3, 5, 7, 2, 4, 6, 8},
	},
	{
		name:   "five increase decrease",
		before: []int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1},
		want:   []int{1, 2, 3, 4, 5},
	},
	{
		name:   "ten distinct",
		before: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		want:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
	{
		name:   "ten repeated",
		before: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 3, 5, 7, 9, 2, 4, 6, 8, 10},
		want:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
}

func TestProblem1(t *testing.T) {
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

var sizes = []int{100, 200, 500, 1000, 2000, 5000, 10000}

func BenchmarkProblem1(b *testing.B) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

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

type NameSolution1B struct {
	Name string
	F    func(head *Node[int]) *Node[int]
}

var sortSolutions = []NameSolution1B{
	{Name: "MergeSort", F: RemoveDuplicatesMergeSort},
	{Name: "MergeSort2", F: RemoveDuplicatesMergeSort2},
	{Name: "HeapSort", F: RemoveDuplicatesHeapSort},
}

func TestProblem1Sorted(t *testing.T) {
	for _, solution := range sortSolutions {
		t.Run(solution.Name, func(t *testing.T) {
			for _, tc := range tt {
				t.Run(tc.name, func(t *testing.T) {
					before := ToList(tc.before...)
					afterList := solution.F(before)
					after := FromList(afterList)

					want := make([]int, len(tc.want))
					copy(want, tc.want)
					sort.Ints(want)

					if diff := cmp.Diff(want, after); diff != "" {
						t.Errorf("%s: %s", solution.Name, diff)
					}
				})
			}
		})
	}
}

func BenchmarkProblem1Sorted(b *testing.B) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomLists := make([][]int, len(sizes))
	for j, size := range sizes {
		randomLists[j] = make([]int, size)
		for i := 0; i < size; i++ {
			randomLists[j][i] = rng.Int()
		}
	}

	for _, solution := range sortSolutions {
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
