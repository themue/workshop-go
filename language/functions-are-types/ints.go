// Go Workshop - Language - Functions are also types

package main

import "fmt"

// IntWorker defines the signature of a function that takes
// an integer as argument and returns an integer.
type IntWorker func(in int) (out int)

// Adder creates an IntWorker for adding based on the
// given addend. Shows also how named return values
// work.
func Adder(addend int) (adder IntWorker) {
	// adder defines the IntWorker with a closure.
	adder = func(in int) (out int) {
		return addend + in
	}
	return
}

// Multiplier creates an IntWorker for multiplying based
// on the given multiplicand.
func Multiplier(multiplicand int) (multiplier IntWorker) {
	// multiplier defines the IntWorker with a closure.
	multiplier = func(in int) (out int) {
		return multiplicand * in
	}
	return
}

// SomeIntegerWorkers shows the usage of the functions.
func SomeIntegerWorkers() {
	fmt.Println("----- Some Integer Workers")

	oneAdder := Adder(1)
	fiveAdder := Adder(5)
	twoMultiplier := Multiplier(2)
	tenMultiplier := Multiplier(10)

	fmt.Printf("One-Adder with 1 = %d\n", oneAdder(1))
	fmt.Printf("One-Adder with 9 = %d\n", oneAdder(9))

	fmt.Printf("Five-Adder with 0 = %d\n", fiveAdder(0))
	fmt.Printf("Five-Adder with 10 = %d\n", fiveAdder(10))

	fmt.Printf("Two-Multiplier with 1 = %d\n", twoMultiplier(1))
	fmt.Printf("Two-Multiplier with 5 = %d\n", twoMultiplier(5))

	fmt.Printf("Ten-Multiplier with 1 = %d\n", twoMultiplier(1))
	fmt.Printf("Ten-Multiplier with 5 = %d\n", twoMultiplier(5))

	queued := tenMultiplier(fiveAdder(twoMultiplier(oneAdder(1))))

	fmt.Printf("Queued with 1 = %d\n", queued)
}
