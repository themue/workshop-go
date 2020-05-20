// Go Workshop - Lesson 12
// For
package main

import (
	"fmt"
)

// SingleConditionFor shows a for based on a boolean condition.
func SingleConditionFor() {
	x := 0
	for x < 5 {
		x++

		fmt.Println("single condition for: x is", x)
	}
}

// ForClauseFor shows a for using a for clause woth an init statement,
// a condition, and a post statement.
func ForClauseFor() {
	for x := 0; x < 5; x++ {
		fmt.Println("for clause for: x is", x)
	}
}

func main() {
	SingleConditionFor()
	ForClauseFor()
}
