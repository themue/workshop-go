// Go Workshop - Part 03
// Packages
package main

import (
	"fmt"

	// Import of the own package.
	"github.com/themue/workshop-go/lesson-03/calculator"
)

func main() {
	// Usage of function Add() of the calculator package.
	sum := calculator.Add(1000, 300, 30, 7)

	fmt.Println("Sum is", sum)
}
