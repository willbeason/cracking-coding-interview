package string_problems

import (
	"gonum.org/v1/gonum/stat/distuv"
	"math"
	"math/rand"
)

const StrLen = 12

type RandFloat interface {
	Rand() float64
}

func RandomEnglishWordLength(rng *rand.Rand) int {
	// See https://www.sciencedirect.com/science/article/abs/pii/0378375886901692
	// We want a mean of about 7 and a variance of about 6, so we just add 1 to a
	// Poisson distribution of lambda=6.
	p := distuv.Poisson{
		Lambda: 6,
		Src:    rng,
	}

	return 1 + int(p.Rand())
}

func RandomUppercase(rng RandFloat) rune {
	return rune(int(rng.Rand()*26) + 65)
}

func RandomAlpha(rng RandFloat) rune {
	r := int(rng.Rand() * 52)
	if r < 26 {
		return rune(65 + r)
	} else {
		return rune(97 + (r - 26))
	}
}

func RandomASCII(rng RandFloat) rune {
	r := int(rng.Rand())
	if r >= 95 {
		r %= 95
	}
	return rune(32 + r)
}

func RandomRune(rng RandFloat) rune {
	return rune(rng.Rand() * math.MaxInt32)
}

func RandomString(rng RandFloat, n int, generateRune func(rng RandFloat) rune) string {
	chars := make([]rune, n)

	for i := range chars {
		chars[i] = generateRune(rng)
	}

	return string(chars)
}

// Problem 1: Implement an algorithm to determine if a string has all unique characters.
// What if you cannot use additional data structures?

// Solution1A is a straightforward attempt to solve the problem with a hashmap.
func Solution1A(s string) bool {
	seen := make(map[rune]bool)
	for _, r := range s {
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}

// Solution1A1 is a slight modification to Solution1A where runes are reduced to
// bytes as all are assumed to be ASCII.
func Solution1A1(s string) bool {
	seen := make(map[byte]bool)
	for _, r := range s {
		if seen[byte(r)] {
			return false
		}
		seen[byte(r)] = true
	}
	return true
}

// Solution1A2 is a slight modification to Solution1A where maps are initialized
// to size 12 to avoid needing to reallocate the underlying array.
func Solution1A2(s string) bool {
	seen := make(map[rune]bool, 12)
	for _, r := range s {
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}

// Solution1B uses a bit array.
func Solution1B(s string) bool {
	seen := make([]bool, 1<<7)
	for _, r := range s {
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}

// Solution1B1 is like the above but uses an array directly instead of a slice.
func Solution1B1(s string) bool {
	var seen [1 << 7]bool
	for _, r := range s {
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}

var (
	seenC  = make([]bool, 1<<7)
	seenC2 = make([]bool, 1<<7)
)

// Solution1B2 is like 1B but avoids allocating the slice with each call, instead copying over it.
func Solution1B2(s string) bool {
	copy(seenC, seenC2)
	for _, r := range s {
		if seenC[r] {
			return false
		}
		seenC[r] = true
	}
	return true
}

// Solution1B3 is like 1B but iterates a static number of times.
func Solution1B3(s string) bool {
	seen := make([]bool, 1<<7)
	for i := 0; i < 12; i++ {
		r := s[i]
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}

// Solution1B4 is like 1B but uses the string length to set the number of iterations.
func Solution1B4(s string) bool {
	seen := make([]bool, 1<<7)
	for i := 0; i < len(s); i++ {
		r := s[i]
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}

// Solution1B5 is like 1B4 except it manually creates bit masks.
func Solution1B5(s string) bool {
	seenLow := 0
	seenHigh := 0

	for i := 0; i < len(s); i++ {
		r := s[i]
		if r < 64 {
			rb := 1 << r
			if seenLow&rb != 0 {
				return false
			}
			seenLow |= rb
		} else {
			rb := 1 << (r - 64)
			if seenHigh&rb != 0 {
				return false
			}
			seenHigh |= rb
		}
	}

	return true
}

// Solution1C avoids using other data structures by directly comarping pairs of characters.
func Solution1C(s string) bool {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}

	return true
}
