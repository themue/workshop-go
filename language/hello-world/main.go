// Go Workshop - Language - Hello, World

// Package main has a special role in the world of
// Go packages. It expects a function named main()
// will be compiled into the executable program.
package main

// import adds packages from standard library,
// external repositories, or own code.
import (
	"fmt"
)

// main is the main program without arguments
// and return codes.
func main() {
	// Prefix fmt tells to use the exported function
	// Println from package fmt.
	fmt.Println("Hello, World")
}
