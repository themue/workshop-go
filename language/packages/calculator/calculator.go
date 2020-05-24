// Go Workshop - Language - Packages

// Package calculator defines the own package calculator,
// which has to be import by others for usage.
package calculator

// Add is an exportet function. No need for a keyword, it's
// done by writing it in upper-case. The function here is to
// add a number of integers to the first one.
func Add(a int, bs ...int) int {
	var tmp int = a
	for _, b := range bs {
		tmp = add(tmp, b)
	}
	return tmp
}

// add is a package private function as it is written in
// lower-case. It's just a little helper here adding two
// integers for demonstration.
func add(a, b int) int {
	return a + b
}
