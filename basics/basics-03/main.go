// Go Workshop - Basics 03
// Packages
package main

import (
	// Import of a standard package.
	"fmt"

	// Import of the own calculator package with full reference.
	"github.com/themue/workshop-go/basics/basics-03/calculator"
)

func main() {
	// Usage of function Add() exported by calculator package.
	var sum int = calculator.Add(1000, 300, 30, 7)

	// Usage of function Println() exported by fmt package.
	fmt.Println("Sum is", sum)
}
