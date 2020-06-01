// Go Workshop - Language - Recovers
package main

import (
	"fmt"
)

// RecoveredPanic executes the given function. In case of
// a panic it is recovered and returned as error.
func RecoveredPanic(f func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("error executing f(): %v", r)
		}
	}()
	f()
	return nil
}

// SafeExecutable is a demo interface for a type using
// the recovering.
type SafeExecutable interface {
	Execute() error
	Recover(r interface{}) error
}

// Executor shows ho the given executable is executed
// and in case of a panic it gets the chance try a
// recovering.
func Executor(se SafeExecutable) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = se.Recover(r)
		}
	}()
	return se.Execute()
}

func main() {
	fmt.Println("----- No Panic")
	err := RecoveredPanic(func() {
		fmt.Println("nothing happens")
	})
	fmt.Printf("Err: %v\n", err)
	fmt.Println("----- Recovered Panic")
	err = RecoveredPanic(func() {
		panic("shit happens")
	})
	fmt.Printf("Err: %v\n", err)
}
