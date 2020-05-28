// Go Workshop - Language - Goroutines - Go do it with result

package main

import (
	"fmt"
)

// DoItWithResult also takes a channel as argument. It's for
// returning the result to the caller.
func DoItWithResult(id, count int, resultc chan<- int) {
	fmt.Println("----- Start Goroutine", id)
	sum := 0
	for i := 0; i < count; i++ {
		sum += i
	}
	// Send the result.
	resultc <- sum
	fmt.Println("----- End Goroutine", id)
}

func main() {
	// Create the channels.
	resultc1 := make(chan int)
	resultc2 := make(chan int)

	// Start the goroutines with the channels.
	go DoItWithResult(1, 25, resultc1)
	go DoItWithResult(2, 50, resultc2)

	// Wait for the results.
	sum1 := <-resultc1
	sum2 := <-resultc2

	fmt.Println("Result 1 is", sum1)
	fmt.Println("Result 2 is", sum2)
}
