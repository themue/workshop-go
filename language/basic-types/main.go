// Go Workshop - Language - Basic Types
package main

import (
	"fmt"
	"math"
)

// Integers shows the integers.
func Integers() {
	var a int8 = math.MaxInt8
	var b int16 = math.MaxInt16
	var c int32 = math.MaxInt32
	var d int64 = math.MaxInt64
	var e int = math.MaxInt64 // On 64 bit architectures.

	fmt.Println("----- Integer")
	fmt.Println("int8:", a, "/ int16:", b, "/ int32:", c, "/ int64:", d, "/ int:", e)
}

// UnsignedIntegers shows the unsigned integers.
func UnsignedIntegers() {
	var a uint8 = math.MaxUint8
	var b uint16 = math.MaxUint16
	var c uint32 = math.MaxUint32
	var d uint64 = math.MaxUint64
	var e uint = math.MaxUint64 // On 64 bit architectures.

	fmt.Println("----- Unsigned Integer")
	fmt.Println("uint8:", a, "/ uint16:", b, "/ uint32:", c, "/ uint64:", d, "/ uint:", e)
}

// Floats shows the floating points.
func Floats() {
	var a float32 = math.MaxFloat32
	var b float64 = math.MaxFloat64

	fmt.Println("----- Floats")
	fmt.Println("float32:", a, "/ float64:", b)
}

// Complexes shows the complex numbers.
func Complexes() {
	var a complex64 = complex(math.MaxFloat32, math.MaxFloat32)
	var b complex128 = complex(math.MaxFloat64, math.MaxFloat64)

	fmt.Println("----- Complexes")
	fmt.Println("complex64:", a, "/ complex128:", b)
}

// OtherBasicTypes shows the other simple types.
func OtherBasicTypes() {
	var a string = "Hello, World"
	var b bool = true
	var c byte = math.MaxUint8 // byte is an alias for uint8.
	var d rune = rune('µ')     // rune is an alias for int32.

	fmt.Println("----- Other Simple Tyoes")
	fmt.Println("string:", a, "/ bool:", b, "/ byte:", c, "/ rune:", d)
}

// ZeroValues shows how all types have a zero value
// when they are not initialized.
func ZeroValues() {
	var a int
	var b uint
	var c float64
	var d complex128
	var e string
	var f bool

	fmt.Println("----- Zero Values")
	fmt.Println("int:", a, "/ uint:", b, "/ float64:", c, "/ complex128:", d, "/ string:", e, "/ bool:", f)
}

func main() {
	Integers()
	UnsignedIntegers()
	Floats()
	Complexes()
	OtherBasicTypes()
	ZeroValues()
}
