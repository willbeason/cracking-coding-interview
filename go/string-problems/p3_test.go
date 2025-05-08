package string_problems

import (
	"fmt"
	"gonum.org/v1/gonum/stat/distuv"
	"math/rand"
	"testing"
	"time"
)

type NameSolution3 struct {
	Name string
	F    func(string, string) bool
}

var solution3s = []NameSolution3{
	//{Name: "Solution3A", F: Solution3A},
	{Name: "Solution3A2", F: Solution3A2},
	{Name: "Solution3B", F: Solution3B},
}

type Problem3TestCase struct {
	name        string
	left, right string
	want        bool
}

func TestSolution3(t *testing.T) {
	tt := []Problem3TestCase{
		{name: "empty", left: "", right: "", want: true},
		{name: "left empty", left: "", right: "a", want: false},
		{name: "right empty", left: "a", right: "", want: false},
		{name: "single character", left: "a", right: "a", want: true},
		{name: "repeated character left", left: "aa", right: "a", want: false},
		{name: "repeated character right", left: "a", right: "aa", want: false},
		{name: "repeated character both", left: "aa", right: "aa", want: true},
		{name: "single character", left: "abcdefghij", right: "jihgfedcba", want: true},
		{name: "different repeated character", left: "aa", right: "bb", want: false},
		{name: "different repeated character, same set", left: "aab", right: "abb", want: false},
	}

	for _, solution := range solution3s {
		for _, tc := range tt {
			if got := solution.F(tc.left, tc.right); got != tc.want {
				t.Errorf("%s(%q, %q) = %v, want %v", solution.Name, tc.left, tc.right, got, tc.want)
			}
		}
	}
}

type LeftRightString struct {
	left  string
	right string
}

type LengthData struct {
	length int
	data   []LeftRightString
}

func BenchmarkS3Different(b *testing.B) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	unif := distuv.Uniform{
		Min: 0.0,
		Max: 95.0,
		Src: rng,
	}

	lengths := []int{1, 2, 5, 10, 20, 50, 100}
	data := make([]LengthData, len(lengths))
	for j, length := range lengths {
		data[j] = LengthData{length: length, data: make([]LeftRightString, n)}
		for i := 0; i < n; i++ {
			data[j].data[i].left = RandomString(unif, length, RandomASCII)
			data[j].data[i].right = RandomString(unif, length, RandomASCII)
		}
	}

	for _, solution := range solution3s {
		b.Run(solution.Name, func(b *testing.B) {
			for _, cases := range data {
				b.Run(fmt.Sprintf("%d", cases.length), func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						idx := i & nMask
						solution.F(cases.data[idx].left, cases.data[idx].right)
					}
				})
			}
		})
	}
}

func BenchmarkS3Permutation(b *testing.B) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	unif := distuv.Uniform{
		Min: 0.0,
		Max: 95.0,
		Src: rng,
	}

	lengths := []int{1, 2, 5, 10, 20, 50, 100}
	data := make([]LengthData, len(lengths))
	for j, length := range lengths {
		data[j] = LengthData{length: length, data: make([]LeftRightString, n)}
		for i := 0; i < n; i++ {
			data[j].data[i].left = RandomString(unif, length, RandomASCII)
			right := []rune(data[j].data[i].left)
			rand.Shuffle(len(right), func(i, j int) {
				right[i], right[j] = right[j], right[i]
			})
			data[j].data[i].right = string(right)
		}
	}

	for _, solution := range solution3s {
		b.Run(solution.Name, func(b *testing.B) {
			for _, cases := range data {
				b.Run(fmt.Sprintf("%d", cases.length), func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						idx := i & nMask
						solution.F(cases.data[idx].left, cases.data[idx].right)
					}
				})
			}
		})
	}
}
