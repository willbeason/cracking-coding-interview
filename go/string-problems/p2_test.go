package string_problems

import (
	"fmt"
	"gonum.org/v1/gonum/stat/distuv"
	"math/rand"
	"testing"
	"time"
)

type NameSolution2 struct {
	Name string
	F    func(string) string
}

var solution2s = []NameSolution2{
	{Name: "Solution2A", F: Solution2A},
	//{Name: "Solution2A1", F: Solution2A1},
}

func TestSolution2(t *testing.T) {
	tt := []struct {
		s    string
		want string
	}{
		{"", ""},
		{"a", "a"},
		{"ab", "ba"},
		{"abc", "cba"},
		{"abcde", "edcba"},
		{"abcabc", "cbacba"},
		{"abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba"},
	}

	for _, solution := range solution2s {
		for _, tc := range tt {
			if got := solution.F(tc.s); got != tc.want {
				t.Errorf("%s(%q) = %q, want %q", solution.Name, tc.s, got, tc.want)
			}
		}
	}
}

func BenchmarkSolution2(b *testing.B) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	charDistribution := distuv.Uniform{
		Min: 0.0,
		Max: 95.0,
		Src: rng,
	}

	strLengths := []int{2, 5, 10, 20, 50, 100}
	nStrings := 1 << 10

	randomStrings := make([][]string, len(strLengths))
	for j, strLen := range strLengths {
		randomStrings[j] = make([]string, nStrings)
		for i := 0; i < nStrings; i++ {
			randomStrings[j][i] = RandomString(charDistribution, strLen, RandomASCII)
		}
	}

	for _, solution := range solution2s {
		for j, strLength := range strLengths {
			b.Run(fmt.Sprintf("%s@%d", solution.Name, strLength), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					solution.F(randomStrings[j][i&nMask])
				}
			})
		}
	}
}
