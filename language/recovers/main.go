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
