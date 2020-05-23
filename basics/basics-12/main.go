// Go Workshop - Basics 12
// For
package main

import (
	"fmt"
)

// SingleConditionFor shows a for based on a boolean condition.
func SingleConditionFor() {
	fmt.Println("----- Single Condition For")
	x := 0
	for x < 5 {
		x++

		fmt.Println("single condition for: x is", x)
	}
}

// ForClauseFor shows a for using a for clause woth an init statement,
// a condition, and a post statement.
func ForClauseFor() {
	fmt.Println("----- For Clause For")
	for x := 0; x < 5; x++ {
		fmt.Println("for clause for: x is", x)
	}

	// Action also possible inside the loop.
	for y := 0; y < 5; {
		fmt.Println("for clause for: y is", y)
		y++
	}
}

// RangeClauseFor shows a for using a range clause to iterate
// over arrays, slices, maps, and strings. Another one about
// channels will be shown with concurrency.
func RangeClauseFor() {
	fmt.Println("----- Range Clause For")
	as := [5]int{1, 2, 3, 4, 5}
	bs := []bool{true, false, true, false}
	cs := map[string]int{"one": 1, "two": 2, "three": 3}
	d := "Hello, 世界"

	for i, a := range as {
		fmt.Println("int", i, "has value", a)
	}
	for i, b := range bs {
		fmt.Println("bool", i, "has value", b)
	}
	for ckey, cval := range cs {
		fmt.Println("key", ckey, "has value", cval)
	}
	// Take a deeper look at the index values here.
	for i, r := range d {
		fmt.Println("string", i, "has rune", r)
	}
}

// EndlessFor shows a for where the leaving of the loop
// is controlled inside the loop. A major usage for goroutines
// will be shown later.
func EndlessFor() {
	fmt.Println("----- Endless For")
	x := 0
	for {
		x++
		if x > 5 {
			// Leave the loop.
			break
		}
		fmt.Println("endless for:", x)
	}
	fmt.Println("left endless for")
}

// ContinuedFor shows how to continue a loop depending
// on a condition.
func ContinuedFor() {
	fmt.Println("----- Continued For")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("continued for: i is", i)
	}
}

func main() {
	SingleConditionFor()
	ForClauseFor()
	RangeClauseFor()
	EndlessFor()
	ContinuedFor()
}
