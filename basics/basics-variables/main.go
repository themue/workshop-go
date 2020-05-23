// Go Workshop - Basics - Variables
package main

import (
	"fmt"
)

func main() {
	// Explicit declaration, later assignment.
	var a int

	a = 1

	// Explicit declaration with direct assignment.
	var b string = "Go"

	// Implicit declaration with direct assignment.
	var c = true

	// Implicit creation on direct assignment.
	d := 13.37

	fmt.Println("a:", a, "b:", b, "c:", c, "d:", d)
}
