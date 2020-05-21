// Go Workshop - Basics 04
// Functions
package main

import (
	"errors"
	"fmt"
)

// how has no no arguments and one return value.
func how() string {
	return "Hello"
}

// who has no no arguments and one return value.
func who() string {
	return "World"
}

// greet has multiple arguments and return values.
func greet(how, who string) (string, error) {
	if how == "" {
		return "", errors.New("need to know how to greet")
	}
	if who == "" {
		return "", errors.New("need to know who to greet")
	}
	return how + ", " + who, nil
}

// output has one argument and no return values.
func output(in string) {
	fmt.Println(in)
}

// doIt has no arguments and no return values.
func doIt() {
	var greeting string
	var err error
	greeting, err = greet(how(), who())
	if err != nil {
		fmt.Printf("cannot greet: %v", err)
		return
	}
	output(greeting)
}

func main() {
	doIt()
}
