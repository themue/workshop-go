// Go Workshop - Language - Black Box Testing
package main

import (
	"fmt"

	"github.com/themue/workshop-go/language/black-testing/calculator"
)

func main() {
	c := calculator.New()

	c.Add("a", 1.0, 1.0)

	fmt.Printf("Value of register 'a' is %f", c.Register("a"))
}
