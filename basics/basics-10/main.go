// Go Workshop - Basics 10
// If
package main

import (
	"errors"
	"fmt"
)

// SimpleIf shows an if without an else.
func SimpleIf() {
	if 5 < 10 {
		fmt.Println("simple if: inside if")
		return
	}
	fmt.Println("simple if: never reached")
}

// SimpleIfElse shows an if with an else.
func SimpleIfElse() {
	if 5 < 10 {
		fmt.Println("simple if/else: inside if branch")
	} else {
		fmt.Println("simple if/else: inside else branch")
	}
}

// MultipleIfs shows usage of multiple ifs in one
// statement.
func MultipleIfs() {
	x := 5
	if x < 5 {
		fmt.Println("multiple ifs: smaller than 5")
	} else if x > 5 {
		fmt.Println("multiple ifs: larger than 5")
	} else {
		fmt.Println("multiple ifs: equal to 5")
	}
}

// MultipleValueIf shows usage of if with multiple value
// returning functions, e.g. with an error.
func MultipleValueIf() {
	doIt := func() (int, error) {
		return 0, errors.New("ouch")
	}

	if v, err := doIt(); err != nil {
		fmt.Println("multiple value if: error is", err)
	} else {
		fmt.Println("multiple value if: value is:", v)
	}
}

func main() {
	SimpleIf()
	SimpleIfElse()
	MultipleIfs()
	MultipleValueIf()
}
