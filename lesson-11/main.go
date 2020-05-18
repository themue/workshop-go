// Go Workshop - Lesson 11
// Switch
package main

import (
	"fmt"
)

// SingleValueSwitch shows switching of a single value and
// no default.
func SingleValueSwitch() {
	x := 5
	switch x {
	case 4:
		fmt.Println("single value switch: x is 4")
	case 5:
		fmt.Println("single value switch: x is 5")
	case 6:
		fmt.Println("single value switch: x is 6")
	}
	fmt.Println("single value switch: after switch")
}

// MultipleValuesSwitch shows switching of a multiple
// values and no default.
func MultipleValuesSwitch() {
	x := 5
	switch x {
	case 1, 2, 3:
		fmt.Println("multiple values switch: 1, 2, 3")
	case 4, 5, 6:
		fmt.Println("multiple values switch: 4, 5, 6 / it is", x)
	case 7, 8, 9:
		fmt.Println("multiple values switch: x is 6")
	}
}

func main() {
	SingleValueSwitch()
	MultipleValuesSwitch()
}
