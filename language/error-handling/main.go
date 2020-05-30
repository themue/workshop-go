// Go Workshop - Language - Error Handling
package main

import (
	"errors"
	"fmt"
	"os"
)

// Divide handles the division of dividend by divisor. As
// division by zero is not possible an error will be returned.
func Divide(dividend, divisor float64) (float64, error) {
	if divisor == 0.0 {
		return 0.0, errors.New("division by zero is not possible")
	}
	return dividend / divisor, nil
}

func main() {
	// Go has no exceptions, but the internal interface type
	// error. See https://golang.org/ref/spec#Errors. When needed
	// functions and methods return it as only or last value.
	// Returning nil means no error. When assigned the variable
	// name typically is err.
	d, err := Divide(5.0, 2.0)

	fmt.Println("Division:", d, "/ Error:", err)

	d, err = Divide(5.0, 0.0)

	fmt.Println("Division:", d, "/ Error:", err)

	// Often used quick check with if.
	if f, err := os.Open("/this/file/definitely/does/not/exist"); err == nil {
		// Only here if file exists.
		defer f.Close()
		fmt.Println("*shrug*")
	}
}
