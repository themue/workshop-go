// Go Workshop - Basics - Functions are also types
package main

import (
	"fmt"
)

// IntFunc describes a function that takes an integer
// as argument and returns an integer.
type IntFunc func(in int) (out int)

// Adder creates an IntFunc for adding based on the
// passed addend. Shows also how named return values
// work.
func Adder(addend int) (adder IntFunc) {
	// adder defines the IntFunc with a closure.
	adder = func(in int) (out int) {
		return addend + in
	}
	return
}

// Apply applies a value to a passed IntFunc.
func Apply(v int, i IntFunc) {
	fmt.Printf("applied %d to %T returns %d\n", v, i, i(v))
}

func main() {
	var oneAdder IntFunc = Adder(1)
	var fiveAdder IntFunc = Adder(5)

	Apply(1, oneAdder)
	Apply(2, oneAdder)

	Apply(1, fiveAdder)
	Apply(5, fiveAdder)
}
