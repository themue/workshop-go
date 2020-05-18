// Go Workshop - Lesson 10
// If
package main

import (
	"errors"
	"fmt"
)

func something() (int, error) {
	return 0, errors.New("failure")
}

func main() {
	// Simple if.
	x := 1
	y := 2
	if x < y {
		fmt.Println("x is smaller than y")
	}

	// Simple if/else.
	if y < x {
		fmt.Println("y is smaller than x")
	} else {
		fmt.Println("y is equal to or larger than x")
	}

	// Multiple ifs.
	if y < x {
		fmt.Println("y is smaller than x")
	} else if x == y {
		fmt.Println("x and y are equal")
	} else {
		fmt.Println("x is smaller than y")
	}

	// Multi-value ifs.
	if v, err := something(); err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("v:", v)
	}
}
