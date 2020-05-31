// Go Workshop - Language - Compound Types
package main

import (
	"fmt"
)

// Arrays are sets of variables with fixed sizes.
func Arrays() {
	fmt.Println("----- Arrays")

	var a [3]int
	var b [3]string = [3]string{"a", "-", "z"}

	// Index starts at 0 and unset elements get zero values.
	a[0] = 1
	a[2] = 99

	fmt.Println("ints:", a, "/ strings:", b)
}

// Slices are sets of variables with variable lengths.
func Slices() {
	fmt.Println("----- Slices")

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

	fmt.Println("bools:", a, "/ bytes b:", b, "/ bytes c:", c)
}

// Maps are sets of key/value assignments. Keys have to be
// comparable. So are string, bool, all numeric types, channel,
// pointer, and interfaces. Also valid are structs and array
// that only contains those types.
func Maps() {
	fmt.Println("----- Maps")

	var msi map[string]int

	msi = make(map[string]int)

	msi["foo"] = 12
	msi["bar"] = 34
	msi["baz"] = 56

	// Access with non-assigned keys returns the default value.
	foo := msi["foo"]
	yadda := msi["yadda"]

	fmt.Println("foo:", foo, "/ yadda:", yadda)

	// Access with a second bool assignment returns if key
	// has been set.
	bar, barOK := msi["bar"]
	dunno, dunnoOK := msi["dunno"]

	fmt.Println("bar:", bar, "/ ok:", barOK)
	fmt.Println("dunno:", dunno, "/ ok:", dunnoOK)
}

// Structs are types containing a number of fields. These
// fields may also be other structs or compound types.
func Structs() {
	fmt.Println("----- Structs")

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

	fmt.Println("struct a:", a, "/ struct b:", b, "/ struct c:", c, "/ struct d:", d)
}

// Functions in Go are types too. They can be declared on top level,
// as methods, inside of functions, or dynamically defined functions.
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
	fmt.Println("anonymous function:", func() string {
		return "I'm anonymous"
	})
}

// Channels are used for sending data between goroutines. Types
// may be all, even functions or other channels.
func Channels() {
	fmt.Println("----- Channels")

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

	fmt.Println("receiving a:", <-a, "/ receiving b:", <-b)
}

func main() {
	Arrays()
	Slices()
	Maps()
	Structs()
	Functions()
	Channels()
}
