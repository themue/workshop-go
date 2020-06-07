// Go Workshop - Language - Functions are also types

package main

import (
	"fmt"
)

// Greeter defines the signature of a function greeting
// somebody.
type Greeter func(whom string) string

// Dog is a type with a method matching to the
// Greeter definition.
type Dog struct{}

// Bark has a different name but the signature matches
// Greeter.
func (d Dog) Bark(whom string) string {
	return fmt.Sprintf("Bark, %s!", whom)
}

// Duck is another type with a method matching to the
// Greeter definition.
type Duck struct{}

// Quack has a different name too but the signature
// also matches Greeter.
func (d Duck) Quack(whom string) string {
	return fmt.Sprintf("Quack quack, %s!", whom)
}

// WorkshopGreet is a simple function matching to Greeter.
func WorkshopGreet(whom string) string {
	return fmt.Sprintf("Welcome to my Go workshop, %s.", whom)
}

// SomeGreetings shows the usage of the functions.
func SomeGreetings() {
	fmt.Println("----- Some Greetings")

	var greet Greeter
	var dog = Dog{}
	var duck = Duck{}

	greet = dog.Bark
	fmt.Printf("Dog greets: %q\n", greet("John Doe"))

	greet = duck.Quack
	fmt.Printf("Duck greets: %q\n", greet("John Doe"))

	greet = WorkshopGreet
	fmt.Printf("Trainer greets: %q\n", greet("John Doe"))
}
