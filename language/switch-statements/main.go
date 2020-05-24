// Go Workshop - Language - Switch Statements
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

// DefaultValueSwitch shows switching with a default.
func DefaultValueSwitch() {
	x := 5
	switch x {
	case 1:
		fmt.Println("default value switch: x is 1")
	case 2:
		fmt.Println("default value switch: x is 2")
	case 3:
		fmt.Println("default value switch: x is 3")
	default:
		fmt.Println("default value switch: x is something else (default at end)")
	}

	switch x {
	case 1:
		fmt.Println("default value switch: x is 1")
	default:
		fmt.Println("default value switch: x is something else (default not at end)")
	case 2:
		fmt.Println("default value switch: x is 2")
	case 3:
		fmt.Println("default value switch: x is 3")
	}
}

// FallthroughValueSwitch shows switching with fallthroughs.
func FallthroughValueSwitch() {
	x := 2
	switch x {
	case 1:
		fmt.Println("fallthrough value switch: x is 1")
		fallthrough
	case 2:
		fmt.Println("fallthrough value switch: x is 2")
		fallthrough
	case 3:
		fmt.Println("fallthrough value switch: x is 3?")
		fallthrough
	case 4:
		fmt.Println("fallthrough value switch: x is 4?")
		fallthrough
	case 5:
		fmt.Println("fallthrough value switch: x is 5?")
	}
}

// ExpressionSwitch shows switching based on expressions.
func ExpressionSwitch() {
	x := 5
	y := -5
	switch {
	case x < 0:
		fmt.Println("expression switch: x is lower than 0")
	case y > 0:
		fmt.Println("expression switch: y is greater than 0")
	case x == y:
		fmt.Println("expression switch: x and y have the same value")
	case x > 0 && y < 0:
		fmt.Println("expression switch: x is greater than 0 and y lower than 0")
	case x < 0 && y > 0:
		fmt.Println("expression switch: x is lower than 0 and y greater than 0")
	default:
		fmt.Println("expression switch: I don't care")
	}
}

// TypeSwitch shows switching based on the type of an interface.
func TypeSwitch() {
	var i interface{}
	i = "guess what I am"
	switch ti := i.(type) {
	case int:
		fmt.Println("type switch: i is an int:", ti)
	case string:
		fmt.Println("type switch: i is a string:", ti)
	case bool:
		fmt.Println("type switch: i is a bool:", ti)
	default:
		fmt.Println("type switch: I don't know")
	}
}

func main() {
	SingleValueSwitch()
	MultipleValuesSwitch()
	DefaultValueSwitch()
	FallthroughValueSwitch()
	ExpressionSwitch()
	TypeSwitch()
}
