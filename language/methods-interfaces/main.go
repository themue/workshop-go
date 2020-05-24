// Go Workshop - Language - Methods and Interfaces
package main

import (
	"fmt"
	"strconv"
)

// LikeAClass demonstrates how types may act like classes
// in object-oriented types.
type LikeAClass struct {
	a string
	b int
}

// NewLikeAClass shows th way how to constructors are done
// in Go.
func NewLikeAClass(a string, b int) *LikeAClass {
	// Return a reference for future changes.
	return &LikeAClass{
		a: a,
		b: b,
	}
}

// Increment uses the reference to modify b. Play with
// the pointer type.
func (l *LikeAClass) Increment() {
	l.b++
}

// String only reads the fields.
func (l LikeAClass) String() string {
	return fmt.Sprintf("LikeAClass{a: %s, b: %d}", l.a, l.b)
}

// Integer shows how a simple type can have methods too.
type Integer int

// Add does not motify but would be possible too.
func (i Integer) Add(j Integer) Integer {
	return Integer(int(i) + int(j))
}

// String returns the integer as string.
func (i Integer) String() string {
	return fmt.Sprintf("Integer{%d}", i)
}

// Embedding shows how types with methods can be
// embedded into structs.
type Embedding struct {
	Integer

	description string
}

// NewEmbedding creates the embedding instance.
func NewEmbedding(d string) *Embedding {
	return &Embedding{
		description: d,
	}
}

// Set changes the value of the embedded Integer.
func (e *Embedding) Set(i int) {
	e.Integer = Integer(i)
}

// String creates an own string representation of
// the Embedding.
func (e Embedding) String() string {
	return fmt.Sprintf("%s is %d", e.description, int(e.Integer))
}

// Methods shows the usage of types with methods.
func Methods() {
	l := NewLikeAClass("leet", 1336)

	l.Increment()

	ls := l.String()

	i := Integer(1300)
	j := Integer(37)
	k := i.Add(j)
	ks := k.String()

	e := NewEmbedding("embed")
	f := e.Add(k)
	es := e.String()
	fs := f.String()

	fmt.Println("----- Methods")
	fmt.Println("l as string:", ls, "/ k as string:", ks)
	fmt.Println("e as string:", es, "/ f as string:", fs)
}

// Adder defines the interface of types implementing
// an Add method. Integer and Embedded are matching
// types.
type Adder interface {
	Add(i Integer) Integer
}

// Stringer is implemented by all types with a String
// method. A similar interface is defined in fmt.
type Stringer interface {
	String() string
}

// AdderStringer shows the combination of interfaces,
// here only done by Embedding.
type AdderStringer interface {
	Adder
	Stringer
}

// NewAdder creates an Adder based on the value of s.
// In case of an int that will be an Integer, else an
// Embedding. Who cares.
func NewAdder(s string) Adder {
	i, err := strconv.Atoi(s)
	if err != nil {
		// It's no int, so take Embedding.
		return NewEmbedding(s)
	}
	return Integer(i)
}

// Println is using the Stringer interface.
func Println(prefix string, s Stringer) {
	fmt.Println(prefix, "is", s.String())
}

// Interfaces shows the usage of interfaces.
func Interfaces() {
	addA := NewAdder("1337")
	addB := NewAdder("leet")

	addAI := addA.Add(Integer(1))
	addBI := addB.Add(Integer(1))

	fmt.Println("----- Interfaces")
	Println("addAI", addAI)
	Println("addBI", addBI)
}

func main() {
	Methods()
	Interfaces()
}
