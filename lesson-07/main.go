// Go Workshop - Lesson 07
// Types
package main

import (
	"fmt"
)

// Integers shows the integers.
func Integers() {
	var a int8 = -128
	var b int16 = -32768
	var c int32 = -2147483648
	var d int64 = -9223372036854775808
	var e int = -9223372036854775808

	fmt.Println("----- Integer")
	fmt.Println("a:", a, "b:", b, "c:", c, "d:", d, "e:", e)
}

// UnsignedIntegers shows the unsigned integers.
func UnsignedIntegers() {
	var a uint8 = 255
	var b uint16 = 65535
	var c uint32 = 4294967295
	var d uint64 = 18446744073709551615
	var e uint = 18446744073709551615

	fmt.Println("----- Unsigned Integer")
	fmt.Println("a:", a, "b:", b, "c:", c, "d:", d, "e:", e)
}

// Floats shows the floating points.
func Floats() {
	var a float32 = -13.37
	var b float64 = 13.37

	fmt.Println("----- Floats")
	fmt.Println("a:", a, "b:", b)
}

// Complexes shows the complex numbers.
func Complexes() {
	var a complex64 = complex(-13.37, 73.31)
	var b complex128 = complex(13.37, -73.31)

	fmt.Println("----- Complexes")
	fmt.Println("a:", a, "b:", b)
}

// OtherNumbers shows the numerical types.
func OtherNumbers() {
	var a byte = 255 // byte is an alias for uint8.
	var b rune = 181 // rune is an alias for int32.

	fmt.Println("----- Byte and Rune")
	fmt.Println("a:", a, "b:", b)
}

func main() {
	Integers()
	UnsignedIntegers()
	Floats()
	Complexes()
	OtherNumbers()
}
