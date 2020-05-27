// Go Workshop - Language - Compound Types
package main

import (
	"fmt"
)

// Arrays shows creation and usage of arrays.
func Arrays() {
	var a [3]int
	var b [3]string = [3]string{"a", "-", "z"}

	// Index starts at 0 and unset elements get zero values.
	a[0] = 1
	a[2] = 99

	fmt.Println("----- Arrays")
	fmt.Println("ints:", a, "/ strings:", b)
}

// Slices shows creation and usage of slices.
func Slices() {
	var a []bool
	var b []byte

	// Creation with initial size.
	a = make([]bool, 3)
	a[0] = true
	a[1] = true

	// Direct creation during assignment.
	b = []byte{0, 1, 3, 3}

	// Append is a built-in statement to append
	// elements to slices.
	b = append(b, 1)

	// Or get sub-slices. Take care, they are just
	// references sharing the memory.
	c := b[1:]
	c[3] = 7

	fmt.Println("----- Slices")
	fmt.Println("bools:", a, "/ bytes b:", b, "/ bytes c:", c)
}

// Structs shows working with structures.
func Structs() {
	a := struct {
		b bool
		s string
		i int
	}{true, "go", 1}

	type B struct {
		x string
		y int
		z string
	}

	var b B

	// Unset fields get zero values.
	b = B{
		x: "X",
		z: "Z",
	}

	// Nesting is no problem.
	type C struct {
		c string
		d B
	}

	c := C{"C", B{"x", 1, "y"}}

	// And embedding.
	type D struct {
		C
		d string
	}

	d := D{C{"A", B{"B", 99, "C"}}, "D"}

	fmt.Println("----- Structs")
	fmt.Println("struct a:", a, "/ struct b:", b, "/ struct c:", c, "/ struct d:", d)
}

// Functions shows different slices.
func Functions() {
	var s func() string = func() string {
		return "s"
	}

	type I func() int

	var i I = func() int {
		return 1337
	}

	b := func() bool {
		return true
	}

	fmt.Println("----- Functions")
	fmt.Println("function s:", s(), "/ function i:", i(), "/ struct b:", b())
}

// Channels shows channels and how they are used for
// sending and receiving.
func Channels() {
	var a chan string
	var b chan string

	// Create without buffer.
	a = make(chan string)

	go func() {
		// Send in goroutine, otherwise blocking.
		a <- "Hello"
	}()

	// Create with buffer size 1, sending uses it.
	b = make(chan string, 1)

	b <- "World"

	fmt.Println("----- Channels")
	fmt.Println("receiving a:", <-a, "/ receiving b:", <-b)
}

func main() {
	Arrays()
	Slices()
	Structs()
	Functions()
	Channels()
}
