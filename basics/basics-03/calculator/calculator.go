// Go Workshop - Basics 03
// Packages
package calculator

// Add is an exportet function to add multiple integers.
func Add(a int, bs ...int) int {
	var tmp int = a
	for _, b := range bs {
		tmp = add(tmp, b)
	}
	return tmp
}

// add is a package private to add two integers.
func add(a, b int) int {
	return a + b
}
