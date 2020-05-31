// Go Workshop - Language - Panics
package main

import (
	"flag"
	"math"
	"math/big"
)

// NewFloat is just an extension of big.NewFloat with a
// conditional precheck.
func NewFloat(precheck bool, f float64) *big.Float {
	if precheck && math.IsNaN(f) {
		panic("not a number")
	}
	return big.NewFloat(f)
}

func main() {
	var precheck bool

	flag.BoolVar(&precheck, "precheck", false, ".")
	flag.Parse()

	NewFloat(precheck, math.NaN())
}
