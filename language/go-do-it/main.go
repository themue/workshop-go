// Go Workshop - Language - Goroutines - Go do it

package main

import (
	"fmt"
	"time"
)

// DoIt is a simple function that will be used as goroutine.
// Those may have arguments and return values, but the
// latter will be ignored. The simpliest pattern is to start
// goroutine for some work and let it terminate when done.
func DoIt(id, count int, d time.Duration) {
	fmt.Println("----- Start Goroutine", id)
	for i := 0; i < count; i++ {
		fmt.Println("id:", id, "run", i)
		time.Sleep(d)
	}
	fmt.Println("----- End Goroutine", id)
}

func main() {
	// Start goroutines with go statement.
	go DoIt(1, 25, 2*time.Millisecond)
	go DoIt(2, 50, time.Millisecond)

	// Have to wait a bit, otherwise program would
	// terminate before goroutines are done. There's
	// no waiting until all are done.
	time.Sleep(time.Second)
}
