package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat/distuv"
	"math/rand"
	"strings"
	"time"
)

const letters = "etaoinshrdlcumwfgypbvkjxqz"

func main() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	pareto := distuv.Pareto{
		Xm:    3.0,
		Alpha: 0.4,
		Src:   rng,
	}

	n := 1 << 20
	counts := make([]int, 1<<10)

	for i := 0; i < n; i++ {
		r := uint(pareto.Rand()) - 3
		r %= 95
		counts[r]++
	}

	for i := range 26 {
		fmt.Printf("%s,%.02f%%\n", strings.ToUpper(string(letters[i])), float64(counts[i]*100)/float64(n))
	}

	//fmt.Printf("%.02f%%\n", 100*float64(c)/float64(n))
}
