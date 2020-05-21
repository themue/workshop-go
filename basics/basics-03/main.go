// Go Workshop - Basics 03
// Packages
package main

import (
	// Import of a standard package.
	"fmt"

	// Import of the own package.
	"github.com/themue/workshop-go/basics/basics-03/calculator"
)

func main() {
	// Usage of function Add() of the calculator package.
	var sum int = calculator.Add(1000, 300, 30, 7)

	fmt.Println("Sum is", sum)
}
