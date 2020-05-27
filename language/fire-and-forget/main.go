// Go Workshop - Language - Goroutines - Fire and Forget

package main

import (
	"fmt"
	"time"
)

// DoIt is a simple function that will be used as goroutine.
// It may have arguments but the return value(s) will be
// ignored.
//
// The simpliest pattern is let a goroutine do something
// and end then.
func DoIt(id, count int, d time.Duration) {
	fmt.Println("----- Start Goroutine", id)
	for i := 0; i < count; i++ {
		fmt.Println("id:", id, "run", i)
		time.Sleep(d)
	}
	fmt.Println("----- End Goroutine", id)
}

func main() {
	// Start goroutines with go.
	go DoIt(1, 25, 2*time.Millisecond)
	go DoIt(2, 50, time.Millisecond)

	// Have to wait, otherwise program would stop
	// too fast.
	time.Sleep(time.Second)
}
