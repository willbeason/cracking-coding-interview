package chapter1

import (
	"gonum.org/v1/gonum/stat/distuv"
	"math/rand"
	"testing"
	"time"
)

type NameSolution1 struct {
	Name string
	F    func(string) bool
}

var solutions = []NameSolution1{
	//{Name: "Solution1A", F: Solution1A},
	//{Name: "Solution1A1", F: Solution1A1},
	{Name: "Solution1A2", F: Solution1A2},
	//{Name: "Solution1B", F: Solution1B},
	//{Name: "Solution1B1", F: Solution1B1},
	//{Name: "Solution1B2", F: Solution1B2},
	//{Name: "Solution1B3", F: Solution1B3},
	{Name: "Solution1B4", F: Solution1B4},
	{Name: "Solution1B5", F: Solution1B5},
	{Name: "Solution1C", F: Solution1C},
}

func TestSolution1(t *testing.T) {
	tt := []struct {
		s    string
		want bool
	}{
		{"", true},
		{"a", true},
		{"abc", true},
		{"abcde", true},
		{"abcabc", false},
		{"abcdefghijklmnopqrstuvwxyza", false},
		{"abcdefghijklmnopqrstuvwxyz", true},
	}

	for _, solution := range solutions {
		for _, tc := range tt {
			if got := solution.F(tc.s); got != tc.want {
				t.Errorf("%s(%q) = %v, want %v", solution.Name, tc.s, got, tc.want)
			}
		}
	}

}

const (
	n     = 1 << 10
	nMask = n - 1
)

func BenchmarkSolution1(b *testing.B) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	unif := distuv.Uniform{
		Min: 0.0,
		Max: 95.0,
		Src: rng,
	}

	randomStrings := make([]string, n)
	for i := 0; i < n; i++ {
		randomStrings[i] = RandomString(unif, StrLen, RandomASCII)
	}

	for _, solution := range solutions {
		b.Run(solution.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				solution.F(randomStrings[i&nMask])
			}
		})
	}
}

func BenchmarkSolution1P(b *testing.B) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	pareto := distuv.Pareto{
		Xm:    1.0,
		Alpha: 1.0,
		Src:   rng,
	}

	randomStrings := make([]string, n)
	for i := 0; i < n; i++ {
		randomStrings[i] = RandomString(pareto, StrLen, RandomASCII)
	}

	for _, solution := range solutions {
		b.Run(solution.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				solution.F(randomStrings[i&nMask])
			}
		})
	}
}
